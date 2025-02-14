package mq

import (
	"encoding/json"
	"fmt"

	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/config"
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

	return &MQ{
		conn:    conn,
		channel: channel,
	}
}

func (mq *MQ) PublishAccountBalance(balance uint32) error {

	r := shared.AccountBalanceMessage{Balance: balance}
	j, _ := json.Marshal(r)

	err := mq.channel.Publish(
		shared.MQExchange,
		shared.TransactionBalanceRoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        j,
		},
	)

	return err
}
