package websocketConn

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"sync"
)

// 使用该方法时，请注意 onlineMap 的清理，防止 onlineMap 无限变大导致内存泄漏
/*
	第一种方式，建议断开连接时清理
		conn, id, err := websocketConn.Init(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		defer websocketConn.Delete(id)

	第二种方式 补漏， 清理发送消息未反馈的websocket
		定时清理 websocketConn.Clean()
*/
var onlineMap sync.Map

func Init(c *gin.Context) (*Conn, string, error) {
	conn, err := NewByGin(c)
	if err != nil {
		return nil, "", err
	}

	id := uuid.New().String()
	for i := 0; i < 100; i++ {
		if _, ok := onlineMap.Load(id); !ok {
			break
		}
		id = uuid.New().String()
	}
	if id == "" {
		return nil, "", errors.New("生成唯一id失败")
	}

	onlineMap.Store(id, conn)
	return conn, id, nil
}
func GetConn(id string) (*Conn, error) {
	conn, ok := onlineMap.Load(id)
	if !ok {
		return nil, errors.New("连接不存在")
	}
	c, ok := conn.(*Conn)
	if !ok {
		onlineMap.Delete(id)
		return nil, errors.New("连接不存在")
	}
	return c, nil
}
func Message(id string, message []byte) error {
	conn, err := GetConn(id)
	if err != nil {
		return err
	}

	if err := conn.Write(message); err != nil {
		Close(id)
	}
	return nil
}
func Clean() {
	onlineMap.Range(func(key, value any) bool {
		conn, ok := value.(Conn)
		if !ok {
			onlineMap.Delete(key)
			return true
		}
		if err := conn.Write([]byte("HeartBeat")); err != nil {
			conn.Close()
			onlineMap.Delete(key)
		}
		return true
	})
}
func Close(id string) {
	value, ok := onlineMap.LoadAndDelete(id)
	if !ok {
		return
	}
	if conn, ok := value.(Conn); ok {
		conn.Close()
	}
}
