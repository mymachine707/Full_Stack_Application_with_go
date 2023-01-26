package main

import (
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func main() {

	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT) // : This creates a TCP

	if err != nil {
		log.Fatal("Error starting tcp server : ", err)
	}

	defer listener.Close() // This defer statement closes a TCP socket listener when the application closes.

	log.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		log.Println(conn)
	}
}
