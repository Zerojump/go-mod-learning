package rabbit_demo

import (
	"github.com/streadway/amqp"
)

var conn, _ = amqp.DialConfig("amqp://guest:guest@192.168.204.128:5672/",amqp.Config{SASL: []amqp.Authentication{&amqp.PlainAuth{Username:"admin",Password:"admin"}}})

