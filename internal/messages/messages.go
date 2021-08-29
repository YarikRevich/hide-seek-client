package messages

import (
	// "github.com/hajimehoshi/ebiten/text"
)

var (
	messageListener *MessageListener 
)

type MessageListener struct {
	msgChan chan string
}

func (ms *MessageListener) Send(msg string){
	ms.msgChan <- msg
}

func (ms *MessageListener) Listen(){
	go func(){
		for range ms.msgChan{

			//Print message
		}
	}()
}

func GetInstance()*MessageListener{
	if messageListener == nil{
		return &MessageListener{
			msgChan: make(chan string, 1),
		}
	}
	return messageListener
}