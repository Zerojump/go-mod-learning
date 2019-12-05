package rabbit_demo

import (
	"fmt"
	"github.com/streadway/amqp"
	"go13-learning/src/commons"
	"log"
	"testing"
	"time"
)

func TestTopicEmit(t *testing.T) {
	ch, err := conn.Channel()
	commons.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	commons.FailOnError(err, "Failed to declare an exchange")

	body := fmt.Sprintf("now is %s", time.Now())
	err = ch.Publish(
		"logs_topic",         // exchange
		"to.from.fish.panda", // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	commons.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

func TestTopicReceive(t *testing.T) {
	ch, err := conn.Channel()
	commons.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	commons.FailOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	commons.FailOnError(err, "Failed to declare a queue")

	routingKey := "#.fish"
	log.Printf("Binding queue %s to exchange %s with routing key %s",
		q.Name, "logs_topic", routingKey)
	err = ch.QueueBind(
		q.Name,       // queue name
		routingKey,   // routing key
		"logs_topic", // exchange
		false,
		nil)
	commons.FailOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	commons.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
