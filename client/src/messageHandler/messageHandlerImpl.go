package messageHandler

import (
	"sessionSender"
	_"log"
	"time"
	"cryptographer"
	"message"
)

type MessageHandlerImpl struct{
	sessSender sessionSender.SessionSender
	convosHandler conversationsHandler
	decrypter cryptographer.Cryptographer
}

func NewMessageHandlerImpl(sessSender sessionSender.SessionSender, decrypter cryptographer.Cryptographer)(*MessageHandlerImpl){
	mhi := new(MessageHandlerImpl)
	mhi.sessSender = sessSender
	mhi.decrypter = decrypter
	return mhi
}

func (handler *MessageHandlerImpl)HandleBytes(from string, bytes []byte){
	//log.Print("Handling bytes " + string(bytes))

	msg := message.FromBytes(from, bytes)

	//TODO remove debug delay
	time.Sleep(time.Second)
	handler.handle(msg)
}

func (handler *MessageHandlerImpl)handle(msg *message.Message){
	msg.SetMessageContent(handler.decrypter.Decrypt(msg.GetEncType(), msg.GetMessageContent()))

	switch msg.GetMessageType(){
	case message.MSG:
		handler.handleMSG(msg)
		break
	case message.MSG_OK:
		handler.handleMSG_OK(msg)
		break
	case message.PING:
		handler.handlePING(msg)
		break
	}
}