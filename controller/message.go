package controller

//Message is the standard communication format for the app
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}
