package server

import (
	"connectionListener"
	"sync"
	"serverCrypto"
	"log"
	"fmt"
	"net"
	"strings"
	"serverEntry"
)

type Server struct{
	sessions map[string][256]byte
	clientListener *connectionListener.ConnectionListener
	serverListener *connectionListener.ConnectionListener
	serverList map[string]serverEntry.ServerEntry
	serverCrypto *serverCrypto.ServerCrypto
	wg sync.WaitGroup
}

func NewServer()(*Server){
	srv := new(Server)

	srv.wg.Add(2)
	srv.serverCrypto = serverCrypto.NewServerCrypto();
	return srv
}

func (srv *Server)sendTo(ip string,  message []byte) {

	fmt.Println(ip)



	connection, err := net.Dial("tcp", ip )
	if err != nil {
		log.Fatal(err)
	}
	connection.Write(message)
}

func getPort(ip string)string{
	return strings.Split(ip, ":")[1]
}

func (srv *Server)Start()error{
	var err error
	srv.clientListener, err = connectionListener.NewConnectionListener("4000", srv)
	if err != nil {
		log.Fatal(err)
	}
	srv.serverListener, err = connectionListener.NewConnectionListener("4001", srv)
	if err != nil {
		log.Fatal(err)
	}
	if err == nil{
		go srv.clientListener.Loop(srv.wg)
		go srv.serverListener.Loop(srv.wg)
		srv.wg.Wait()
	}

	return err
}

func (srv *Server)CreateSession(name string, socket net.Conn){
	//TODO
}

func (srv *Server)RemoveSession(name string){
	//TODO
}