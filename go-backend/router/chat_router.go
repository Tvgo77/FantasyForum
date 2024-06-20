package router

import (
	"github.com/gin-gonic/gin"

	"go-backend/controller"
	"go-backend/setup"
	"go-backend/usecase"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ChatRouterSetup(env *setup.Env, group *gin.RouterGroup) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	// Empty producer event handler
	go func() {
		for e := range producer.Events() {
			if e != nil {
				continue
			}
			// switch ev := e.(type) {
			// case *kafka.Message:
			// 	if ev.TopicPartition.Error != nil {
			// 		fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
			// 	} else {
			// 		fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
			// 			*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
			// 	}
			// }
		}
	}()


	cu := usecase.NewChatUsecase(producer, env)
	cc := controller.NewChatController(cu, env)

	group.Handle("POST", "/chat", cc.Chat)
}
