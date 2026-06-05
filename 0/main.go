package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Println(err)
			return
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

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
