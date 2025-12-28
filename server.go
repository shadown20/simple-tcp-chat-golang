package main

import (
	"fmt"
	"net"
	"sync"
)

var (
	clients   = make(map[net.Conn]bool)
	clientsMU sync.Mutex
)

func main() {
	start, err := net.Listen("tcp", ":4444")
	if err != nil {
		fmt.Printf("Error Set up server: %v", err)
		return
	}

	defer start.Close()
	fmt.Println("---SERVER-IS-UP---")

	for {
		conn, err := start.Accept()
		if err != nil {
			fmt.Printf("Error connection: %v", err)
			continue
		}

		username := make([]byte, 1024)
		conn.Read(username)

		fmt.Printf("%v joined to chat", string(username))

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer handleConnection(conn)
	clientsMU.Lock()
	delete(clients, conn)
	clientsMU.Unlock()

	buffer := make([]byte, 1024)
	for {
		reading, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Error sending data: %v", err)
			continue
		}
		message := string(buffer[:reading])

		broadcast(message, conn)
	}
}

func broadcast(message string, conn net.Conn) {
	clientsMU.Lock()
	defer clientsMU.Unlock()
	for client := range clients {
		if client != conn {
			_, err := client.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Printf("Error broadcasting: %v", err)
				client.Close()
				delete(clients, client)
			}

		}
	}
}
