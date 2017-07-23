package lib

import (
	"github.com/gorilla/websocket"
	"log"
)

type Connection struct {
	ws     *websocket.Conn
	pool   *Pool
	action string
}

func (c *Connection) SendMessage(msg string) {
	go func() {
		log.Println("send message: ", msg)
		err := c.ws.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			c.pool.Leave <- c
			c.ws.Close()
		}
	}()
}

func (c *Connection) receive() {
	defer func() {
		log.Println("start Player leave: ", c, c.pool.Leave)
		c.pool.Leave <- c
		c.ws.Close()
	}()

	for {
		log.Println("run receive")
		_, action, err := c.ws.ReadMessage()
		if err != nil {
			log.Println("Read message: ", err)
			break
		}
		c.action = string(action)
		c.pool.AddAction <- c
	}
}

func NewConnection(ws *websocket.Conn, pool *Pool) *Connection {
	c := &Connection{ws: ws, pool: pool}
	log.Println("NewConnection created")
	go c.receive()
	log.Println("NewConnection receive start")

	c.pool.Join <- c

	return c
}
