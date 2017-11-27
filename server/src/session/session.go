package session

import (
	"net"
	"sync"
	"log"
	"sender"
)

type Session struct{
	name string
	socket *net.Conn
	sender *sender.SenderImpl
	receiver *receiver.Receiver
	wg sync.WaitGroup
}

func NewSession(socket *net.Conn, name string, messageHandler *messageHandler.MessageHandler)(*Session){
	session := new(Session)

	session.name = name
	session.socket = socket
	session.sender = sender.NewSenderImpl(socket)
	session.receiver = receiver.NewReceiver(socket, messageHandler)
	session.wg.Add(2)

	return session
}

func (session *Session)Start(){
	log.Print("Starting session " + session.name)

	go session.sender.Loop(session.wg)
	go session.receiver.Loop(session.wg)

	session.wg.Wait()

	log.Print("Session ended " + session.name)
}

func (session *Session)DeleteSession(){
	session.socket.Close()
}

func (session *Session)Send(bytes []byte){
	session.sender.send(bytes)
}

func (session *Session)UnlockSending(){
	session.sender.UnlockSending()
}
