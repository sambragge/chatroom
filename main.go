package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sambragge/chatroom/controller"
)

func main() {

	port := ":8000"
	r := mux.NewRouter()

	r.HandleFunc("/", controller.ServeView)
	r.HandleFunc("/bundle", controller.ServeBundle)
	r.HandleFunc("/ws", controller.ServeWS)

	http.Handle("/", r)
	fmt.Println("Server listining on port ", port)
	http.ListenAndServe(port, nil)

}
