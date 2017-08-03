package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}

//Handler is the format of the handler
type Handler func(*Client, interface{})

//ServeView serves the main html page
func ServeView(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./view/index.html")
	return

}

//ServeBundle serves the bundled js
func ServeBundle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./view/bundle.js")
	return
}

//ServeWS opens up the websocket connection and starts the clients read and write jobs.
func ServeWS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in the ServeWS handler")
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	client := NewClient(socket)
	go client.Write()
	go client.Read()
}

//Handlers is a convinent list of all socket responses. To automate the handling of socket messages.
var Handlers = map[string]Handler{
	"connected": func(client *Client, data interface{}) {
		fmt.Println("a user joined the room")
		client.UserCount++
		client.Send <- Message{
			Name: "server handshake",
		}
		return
	},

	"disconnect": func(client *Client, data interface{}) {
		fmt.Println("a user left the room")
		client.UserCount--
		return
	},
}
