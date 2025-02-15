package mq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/config"
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

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%v:%v@localhost:5672/", conf.RabbitMQ.User, conf.RabbitMQ.Password))
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
		shared.TransactionBalanceQueue,
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
		shared.TransactionBalanceRoutingKey,
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

func (mq *MQ) ConsumeAccountBalance(ctx context.Context, res chan<- entity.Account) error {

	msg, err := mq.channel.Consume(
		shared.TransactionBalanceQueue,
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
			r := entity.Account{}
			err = json.Unmarshal(d.Body, &r)
			if err != nil {
				return err
			}
			res <- r
		}
	}

}
