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
	port := ":8080"

	envPort := os.Getenv("PORT")
	if envPort != "" {
		port = envPort
	}

	fmt.Println("Using port from environment variable:", port)

	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		fmt.Println("Waiting for connection...")
		conn, err := ln.Accept()

		fmt.Println("Connection accepted")

		if err != nil {
			log.Println(err)
			return
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

}
