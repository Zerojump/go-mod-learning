package rabbit_demo

import (
	"github.com/streadway/amqp"
)

var conn, _ = amqp.DialConfig("amqp://guest:guest@localhost:5672/", amqp.Config{SASL: []amqp.Authentication{&amqp.PlainAuth{Username: "admin", Password: "admin"}}})
