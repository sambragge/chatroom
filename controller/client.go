package controller

import (
	"fmt"

	"github.com/gorilla/websocket"
)

//Client handles the reading and writing of socket messages
type Client struct {
	Send         chan Message
	Socket       *websocket.Conn
	UserCount    int
	ChatMessages []ChatMessage
}

//NewClient returns a pointer to a new client
func NewClient(s *websocket.Conn) *Client {
	return &Client{
		Send:         make(chan Message),
		Socket:       s,
		UserCount:    0,
		ChatMessages: make([]ChatMessage, 0),
	}
}

func (c *Client) findHandler(name string) (Handler, bool) {
	for i, v := range Handlers {
		if i == name {
			fmt.Println("found the handler")
			return v, true
		}
	}
	fmt.Println("could not find the handler")
	return nil, false
}

func (c *Client) Read() {
	var message Message
	for {
		if err := c.Socket.ReadJSON(&message); err != nil {
			break
		}
		handler, found := c.findHandler(message.Name)
		if found {
			handler(c, message.Data)
		}
	}
	c.Socket.Close()
}

func (c *Client) Write() {
	for msg := range c.Send {
		if err := c.Socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.Socket.Close()
}
