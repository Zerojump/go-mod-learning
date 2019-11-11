package kafka_demo

import (
	"github.com/Shopify/sarama"
	"go13-learning/src/commons"
	"testing"
)

func TestSend(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner  = sarama.NewRandomPartitioner
	config.Producer.Return.Successes  = true

		client, err := sarama.NewSyncProducer([]string{"192.168.204.128:9092"}, config)
	commons.FailOnError(err,"connect kafka err")
	defer client.Close()

	msg := &sarama.ProducerMessage{Topic: "test_log", Value: sarama.StringEncoder("this is a test log")}
	pid, offset, err := client.SendMessage(msg)
	commons.FailOnError(err, "send msg failed")
	t.Logf("pid:%v, offset:%v",pid,offset)
}