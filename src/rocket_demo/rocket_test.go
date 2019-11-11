package rocket_demo

import (
	rocketmq "github.com/didapinchegit/go_rocket_mq"
	"log"
	"testing"
)

func TestConsume(t *testing.T) {
	conf := &rocketmq.Config{Nameserver: "192.168.204.128:9876", ClientIp: "10.224.16.75", InstanceName: "DEFAULT",}
	consumer, err := rocketmq.NewDefaultConsumer("C_TEST", conf)
	if err != nil {
		log.Panic(err)
	}
	consumer.Subscribe("test2", "")
	consumer.Subscribe("test3", "")
	consumer.RegisterMessageListener(func(msgs []*rocketmq.MessageExt) error {
		for i, msg := range msgs {
			log.Print(i, string(msg.Body))
		}
		return nil
	})
	consumer.Start()
}

func TestProduce(t *testing.T) {
	conf := &rocketmq.Config{Nameserver: "192.168.204.128:9876", ClientIp: "10.224.16.75", InstanceName: "DEFAULT",}
	mqClient := rocketmq.NewMqClient()
	mqClient.

}
