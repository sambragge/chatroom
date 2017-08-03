package controller

import "time"

//ChatMessage is the format of chat messages
type ChatMessage struct {
	Content   string
	TimeStamp time.Time
}

//NewChatMessage is used when entering a new message into the clients ChatMessages slice
func NewChatMessage(content string) ChatMessage {
	return ChatMessage{
		Content:   content,
		TimeStamp: time.Now(),
	}
}
