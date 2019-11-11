package rabbit_demo

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"testing"
	"time"
	"go13-learning/src/commons"
)

func TestDirectEmit(t *testing.T) {
	ch, err := conn.Channel()
	commons.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	commons.FailOnError(err, "Failed to declare an exchange")

	body := fmt.Sprintf("now is %s", time.Now())
	err = ch.Publish(
		"logs_direct", // exchange
		"info",        // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	commons.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

func TestDirectReceive(t *testing.T) {
	ch, err := conn.Channel()
	commons.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
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

	routing_key := "info"

	log.Printf("Binding queue %s to exchange %s with routing key %s",
		q.Name, "logs_direct", routing_key)
	err = ch.QueueBind(
		q.Name,        // queue name
		routing_key,   // routing key
		"logs_direct", // exchange
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
