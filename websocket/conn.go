package websocketConn

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func NewByGin(c *gin.Context) (*Conn, error) {
	wsConnect := &websocket.Upgrader{
		ReadBufferSize:   4096,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Minute,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ctx, cancel := context.WithCancel(context.Background())

	conn := &Conn{
		inChan:  make(chan []byte, 1000),
		outChan: make(chan []byte, 1000),
		ctx:     ctx,
		cancel:  cancel,
	}
	wsConn, err := wsConnect.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return nil, err
	}
	if wsConn == nil {
		return nil, errors.New("无长连接信息")
	}
	conn.wsConnect = wsConn

	go conn.readLoop()
	go conn.writeLoop()
	return conn, nil
}

type Conn struct {
	wsConnect *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	ctx       context.Context
	cancel    context.CancelFunc
}

func (c *Conn) Close() {
	// 线程安全，可多次调用
	_ = c.wsConnect.Close()

	c.cancel()
}
func (c *Conn) Read() (data []byte, err error) {
	select {
	case data = <-c.inChan:
	case <-c.ctx.Done():
		err = errors.New("connection is closed")
	}
	return
}
func (c *Conn) Write(data []byte) (err error) {
	select {
	case c.outChan <- data:
	case <-c.ctx.Done():
		err = errors.New("connection is closed")
	}
	return
}
func (c *Conn) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = c.wsConnect.ReadMessage(); err != nil {
			goto ERR
		}

		//阻塞在这里，等待inChan有空闲位置
		select {
		case c.inChan <- data:
		case <-c.ctx.Done(): // closeChan 感知 conn断开
			goto ERR
		}
	}
ERR:
	c.Close()
}
func (c *Conn) writeLoop() {
	var (
		data []byte
		err  error
	)
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-c.ctx.Done():
			goto ERR
		case data = <-c.outChan:
		case <-ticker.C:
			data = []byte("HeartBeat")
		}
		if err = c.wsConnect.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	c.Close()
}
