package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on port:", port)

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Handling connection...")

	buffer := make([]byte, 0, 4096)
	tmp := make([]byte, 256)

	for {

		n, err := conn.Read(tmp)
		if err == io.EOF {
			fmt.Println("connection closed cause of EOF")
			break
		}

		if err != nil {
			fmt.Println("error reading the from the connection", err.Error())
			return
		}

		buffer = append(buffer, tmp[:n]...)
	}

	for len(buffer) > 0 {
		n, err := conn.Write(buffer)
		if err != nil {
			log.Println("write error:", err)
			return
		}

		buffer = buffer[n:]
	}
	fmt.Println("Connection closed")

}
