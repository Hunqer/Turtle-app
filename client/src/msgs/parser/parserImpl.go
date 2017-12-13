package parser

import (
	_"log"
	"time"
	"crypt"
	"message"
	"conversationsHandler"
	"log"
	"sessions/sender"
)

type ParserImpl struct{
	ss            sender.SessionSender
	convoshandler conversationsHandler.ConversationsHandler
	crypt         crypt.Cryptographer
}

func New(sessSender sender.SessionSender, convohandler conversationsHandler.ConversationsHandler, crypt crypt.Cryptographer)(*ParserImpl){
	mhi := new(ParserImpl)
	mhi.ss = sessSender
	mhi.convoshandler = convohandler
	mhi.crypt = crypt
	return mhi
}

func (handler *ParserImpl)HandleBytes(from string, bytes []byte){
	//log.Print("Handling bytes " + string(bytes))

	msg, err := message.FromBytes(bytes)

	if err != nil{
		log.Print(err)
		return
	}
	//TODO remove debug delay
	time.Sleep(time.Second)
	handler.handle(from, msg)
}

func (handler *ParserImpl)handle(from string, msg *message.Message){
	decrypted, err := handler.crypt.Decrypt(msg.GetEncType(), msg.GetMessageContent())
	if err != nil{
		log.Print(err.Error())
		return
	}
	msg.SetMessageContent(decrypted)

	switch msg.GetMessageType(){
	case message.DEFAULT:
		handler.handleDEFAULT(from, msg)
		break
	case message.OK:
		handler.handleOK(from, msg)
		break
	case message.PING:
		handler.handlePING(from, msg)
		break
	}
}