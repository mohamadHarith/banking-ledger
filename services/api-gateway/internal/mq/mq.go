package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/config"
	"github.com/mohamadHarith/banking-ledger/shared/entity"
	shared "github.com/mohamadHarith/banking-ledger/shared/mq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

var mq *MQ
var once sync.Once

func New() *MQ {
	once.Do(func() {
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

		mq = &MQ{
			conn:    conn,
			channel: channel,
		}
	})

	return mq
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

func (mq *MQ) Close() {
	mq.conn.Close()
	mq.channel.Close()
}
