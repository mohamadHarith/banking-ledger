package mq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mohamadHarith/banking-ledger/services/transaction-logger-service/internal/config"
	"github.com/mohamadHarith/banking-ledger/shared/entity"
	shared "github.com/mohamadHarith/banking-ledger/shared/mq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func New() *MQ {
	conf := config.GetConf()

	mqHost := conf.RabbitMQ.ServiceName
	if conf.IsLocalEnvironment() {
		mqHost = "localhost"
	}

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:5672/", conf.RabbitMQ.User, conf.RabbitMQ.Password, mqHost))
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	err = channel.ExchangeDeclare(
		shared.MQExchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	q, err := channel.QueueDeclare(
		shared.TransactionLogQueue,
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	err = channel.QueueBind(
		q.Name,
		shared.TransactionLogRoutingKey,
		shared.MQExchange,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	return &MQ{
		conn:    conn,
		channel: channel,
	}
}

func (mq *MQ) Close() {
	mq.channel.Close()
	mq.conn.Close()
}

func (mq *MQ) ConsumeTransactionLog(ctx context.Context, res chan<- entity.TransactionLog) error {

	msg, err := mq.channel.Consume(
		shared.TransactionLogQueue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case d := <-msg:
			r := entity.TransactionLog{}
			err = json.Unmarshal(d.Body, &r)
			if err != nil {
				return err
			}
			res <- r
		}
	}

}
