package msg_consumer

import (
	"go-chatroom/domain"
	"go-chatroom/setup"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type messageConsumer struct {
	connManager domain.ConnManagerUsecase
	env *setup.Env
}

func NewMessageConsumer(cm domain.ConnManagerUsecase, env *setup.Env) *messageConsumer {
	return &messageConsumer{connManager: cm, env: env}
}

func (mc *messageConsumer) Run() {
	// Connect to Kafka 
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "msgGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	err = consumer.Subscribe("chatMessage", nil)
	if err != nil {
		panic(err)
	}

	// Consume message
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}
		//fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		mc.connManager.Broadcast(string(msg.Value))
	}
}
