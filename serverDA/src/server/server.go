package server

import (
	_ "sync"
	"log"
	"net"
	"serverlist"
	_ "errors"
	"encoding/json"

)

type Server struct{
	myName string
	serverList *serverlist.ServerList
	//clientListener *listener.Listener
	//serverListener *listener.Listener
	//wg sync.WaitGroup
}


type msgMessage struct{
	Type    string
	Content string
	Content2 string
}

func NewServer(name string)(*Server){
	srv := new(Server)

	srv.myName = name
	srv.serverList = serverlist.NewList()
	return srv
}



func (srv *Server)Start(port string){
	log.Println("Port: ", port)
	ln, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Server up and listening on port ", port)

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {

    log.Printf("Client %v connected.", c.RemoteAddr())

	//to read json from stream
	d := json.NewDecoder(c)

    var msg msgMessage

    err := d.Decode(&msg)
    log.Println("Type: ",msg.Type)
    log.Println("Ip: ", msg.Content)
	log.Println("Key: ", msg.Content2)
	log.Println("Errors: ",err)

    c.Close()

    //messageBuffer := make([]byte, 2048)
    //for {
    //  n, err := c.Read(messageBuffer)
    //  log.Println(err)
    //  if err != nil {
    //      c.Close()
    //      break
    //  }
    //  msg := string(messageBuffer)
      //sendTo(ip, msg)

      //log.Printf("%v", )
      //log.Printf("Data: %v , %v", n, messageBuffer[0:n])

    log.Printf("Connection from %v closed.", c.RemoteAddr())
 }


// func (srv *Server)Start(clientPort, serverPort string)error{
// 	var err error
// 	srv.clientListener, err = listener.NewConnectionListener(clientPort)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	srv.serverListener, err = listener.NewConnectionListener(serverPort)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err == nil{
// 		go srv.clientListener.Loop(srv.wg)
// 		go srv.serverListener.Loop(srv.wg)
// 		log.Print("Server started")
// 		srv.wg.Wait()
// 	}
//
// 	return err
// }
//




func (srv *Server)GetName()string{
	return srv.myName
}
