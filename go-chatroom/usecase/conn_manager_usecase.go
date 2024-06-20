package usecase

import (
	"go-chatroom/setup"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type connManager struct {
	connPool *sync.Map
	env *setup.Env
}

func NewConnManager(connPool *sync.Map, env *setup.Env) *connManager {
	return &connManager{connPool: connPool, env: env}
}

func (cm *connManager) AddConn(uid string, conn *websocket.Conn) error {
	cm.connPool.Store(uid, conn)
	return nil
}

func (cm *connManager) RemoveConn(uid string) error {
	cm.connPool.Delete(uid)
	return nil
}

func (cm *connManager) GetOnlineUsers() []string {
	var keys []string
	cm.connPool.Range(func(key any, value any) bool{
		str, _ := key.(string)
		keys = append(keys, str)
		return true
	})
	return keys
}

func (cm *connManager) Broadcast(msg string) {
	cm.connPool.Range(func(key any, value any) bool {
		conn, _ := value.(*websocket.Conn)
		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			conn.Close()
			cm.connPool.Delete(key)
			return true
		}
		return true
	})
}