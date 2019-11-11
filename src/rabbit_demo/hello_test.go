package rabbit_demo

import (
	"github.com/streadway/amqp"
	"log"
	"testing"
)

func TestSend(t *testing.T) {

	ch, err := conn.Channel()
	commons.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	commons.FailOnError(err, "Failed to declare a queue")

	body := "Hello world!"
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})

	log.Printf(" [x] Sent %s", body)
	commons.FailOnError(err, "Failed to publish a message")
}

func TestReceive(t *testing.T) {

	ch, err := conn.Channel()
	commons.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	commons.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d:=range msgs  {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}