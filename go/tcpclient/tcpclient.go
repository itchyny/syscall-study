package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	port := "8080"
	conn, err := net.Dial("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if tcpconn, ok := conn.(*net.TCPConn); ok {
		tcpconn.CloseWrite()
	}
	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
