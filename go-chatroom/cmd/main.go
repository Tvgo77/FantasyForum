package main

import (
	"go-chatroom/controller"
	"go-chatroom/msg_consumer"
	"go-chatroom/setup"
	"go-chatroom/usecase"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	var connPool sync.Map
	env := setup.NewEnv()
	cm := usecase.NewConnManager(&connPool, env)
	msgConsumer := msg_consumer.NewMessageConsumer(cm, env)
	cc := controller.NewChatController(cm, env)

	/* Setup router */
	ginEngine := gin.Default()
	ginEngine.GET("/ws", cc.ConnHandler)

	go msgConsumer.Run()
	ginEngine.Run(":8888")
}