package controller

import (
	"go-chatroom/domain"
	"go-chatroom/setup"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type chatController struct {
	connManager domain.ConnManagerUsecase
	env *setup.Env
}

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(req *http.Request) bool {return true},
}

func NewChatController(cm domain.ConnManagerUsecase, env *setup.Env) *chatController {
	return &chatController{connManager: cm, env: env}
}

func (cc *chatController) ConnHandler(c *gin.Context) {
	// Open websocket connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("fail to upgrade to websocket")
		c.JSON(http.StatusInternalServerError, &domain.ErrorResponse{Message: "Error in websocket"})
	}

	// Push connection to connection pool
	userID := c.GetString("userID")
	cc.connManager.AddConn(userID, conn)
}