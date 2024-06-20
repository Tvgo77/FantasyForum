package controller

import (
	"go-backend/domain"
	"go-backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

type chatController struct {
	chatUsecase domain.ChatUsecase
	env *setup.Env
}

func NewChatController(cu domain.ChatUsecase, env *setup.Env) *chatController {
	return &chatController{chatUsecase: cu, env: env}
}

func (cc *chatController) Chat(c *gin.Context) {
	// Parse request
	var req domain.ChatRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &domain.ErrorResponse{Message: "Bad Request"})
	}
	uid := c.GetString("userID")

	// Send message
	err = cc.chatUsecase.SendMessage(uid, req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &domain.ErrorResponse{Message: err.Error()})
	}

	// Response
	c.JSON(http.StatusOK, nil)
}