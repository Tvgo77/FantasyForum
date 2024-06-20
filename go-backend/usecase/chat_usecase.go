package usecase

import (
	"encoding/json"
	"errors"
	"go-backend/domain"
	"go-backend/setup"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type chatUsecase struct {
	producer *kafka.Producer
	env *setup.Env
}

func NewChatUsecase(p *kafka.Producer, env *setup.Env) domain.ChatUsecase {
	return &chatUsecase{producer: p, env: env}
}

type chatMsg struct {
	Uid string `json:"uid"`
	Msg string `json:"message"`
}

func (cu *chatUsecase) SendMessage(uid string, msg string) error {
	// Marshal message to json
	jsonData, err := json.Marshal(&chatMsg{Uid: uid, Msg: msg})
	if err != nil {
		return err
	}

	topic := "chatMessage"
	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: jsonData,
	}
	
	err = cu.producer.Produce(kafkaMessage, nil)
	if err != nil {
		return err
	}
	n := cu.producer.Flush(10 * 1000)
	if n > 0 {
		return errors.New("message sending timeout")
	}
	return nil
}