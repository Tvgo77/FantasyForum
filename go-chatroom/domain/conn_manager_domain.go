package domain

import "github.com/gorilla/websocket"

type ConnManagerUsecase interface {
	AddConn(uid string, conn *websocket.Conn) error
	RemoveConn(uid string) error
	Broadcast(msg string)
	GetOnlineUsers() []string
}