package main

import (
	"io"
	"log"
	"net"
)

func main() {
	port := "8080"
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
	err = ln.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func handleConnection(conn net.Conn) {
	_, err := io.Copy(conn, conn)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
