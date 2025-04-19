package cond

import (
	"log"

	"github.com/streadway/amqp"
	"github.com/tmazitov/auth_service.git/pkg/conductor/messages"
)

func (c *Conductor) worker(messageChan chan *messages.MessageInfo) {

	var (
		messageInfo *messages.MessageInfo
		messageBody []byte
		err         error
	)

	conn, err := amqp.Dial(c.config.AMQPConfig.GetDSN())
	if err != nil {
		log.Fatalf("Ошибка подключения к RabbitMQ: %s", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Ошибка создания канала: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"emails", // имя очереди
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("Ошибка объявления очереди: %s", err)
	}

	defer close(messageChan)

	for messageInfo = range messageChan {
		if messageBody, err = messageInfo.ToJson(); err != nil {
			log.Fatalf("Ошибка сериализации сообщения: %s", err)
		}
		if err = c.send(ch, q.Name, messageBody); err != nil {
			log.Fatalf("Ошибка отпр	авки сообщения: %s", err)
		}
	}
}

func (c *Conductor) send(ch *amqp.Channel, queueName string, messageBody []byte) error {
	return ch.Publish(
		"",        // exchange
		queueName, // routing key (имя очереди)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageBody,
		})
}
