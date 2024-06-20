package domain

import "github.com/gorilla/websocket"

// type NonBlockWriter struct {
// 	Conn *websocket.Conn
// 	WriteChan chan string
// 	HasFailed bool
// }

// func (wsc *NonBlockWriter) Writer() {
// 	for msg := range wsc.WriteChan {
// 		if err := wsc.Conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
// 			wsc.HasFailed = true
// 			break
// 		}
// 	}
// }

// func (wsc *NonBlockWriter) SendMessage(msg string) {
// 	select {
// 	case wsc.WriteChan <- msg:
// 	default:
// 		wsc.HasFailed = true
// 	}
// }

type ConnManagerUsecase interface {
	AddConn(uid string, conn *websocket.Conn) error
	RemoveConn(uid string) error
	Broadcast(msg string)
	GetOnlineUsers() []string
}