package main

import (
	"log"
	"net"

	"github.com/maeldonn/tcpchat/pkg/server"
)

func main() {
	server := server.NewServer()
	go server.Run()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("Started server on port :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %s", err.Error())
			continue
		}

		client := server.NewClient(conn)
		go client.ReadInput()
	}
}
