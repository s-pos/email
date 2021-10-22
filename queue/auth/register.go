package auth

import (
	"spos/email/queue"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

const registerQueueName = "email_registration_otp"

type registerQueue struct {
	channel *amqp.Channel
}

func NewRegisterQueue(ch *amqp.Channel) queue.Queue {
	return &registerQueue{
		channel: ch,
	}
}

func (evq *registerQueue) GetQueue() string {
	return registerQueueName
}

func (evq *registerQueue) Bind() string {
	return strings.ReplaceAll(registerQueueName, "_", ".")
}

func (evq *registerQueue) Consume(name string) (<-chan amqp.Delivery, error) {
	deliveries, err := evq.channel.Consume(
		name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	return deliveries, err
}

func (evq *registerQueue) Setup() error {
	q, err := evq.channel.QueueDeclare(
		evq.GetQueue(),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = evq.channel.QueueBind(
		q.Name,
		evq.Bind(),
		queue.Exchange,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (evq *registerQueue) HandleConsume(fn queue.ConsumeFunc) error {
	msg, err := evq.Consume(evq.GetQueue())
	if err != nil {
		return err
	}

	var worker = func(workerNumber int, jobs <-chan amqp.Delivery) {
		for job := range jobs {
			logrus.Infof("worker %d", workerNumber)
			fn(job)
		}
	}

	for i := 1; i <= 5; i++ {
		go worker(i, msg)
	}

	return nil
}
