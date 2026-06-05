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

	buf := make([]byte, 4096)

	for {
		n, err := conn.Read(buf)

		if n > 0 {
			written := 0
			for written < n {
				m, werr := conn.Write(buf[written:n])
				if werr != nil {
					return
				}
				written += m
			}
		}

		if err == io.EOF {
			return
		}

		if err != nil {
			return
		}
	}
}
