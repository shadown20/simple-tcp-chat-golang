package client

import (
	"fmt"
	"log"
	"net"
)

func client() {
	var username string
	var addr string
	var message string

	fmt.Printf("Enter the username: \n")
	fmt.Scan(&username)
	fmt.Printf("Enter address and port of server: \n")
	fmt.Scan(&addr)

	connecte, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Error connecting: %s", err)
	}

	connecte.Write([]byte(username))

	for {

		fmt.Printf("Write your message: \n")
		fmt.Scan(&message)

		_, err = connecte.Write([]byte(message))
		if err != nil {
			log.Fatalf("Error sending data: %s", err)
		}

		fmt.Printf("%s: %s", username, message)

		recive := make([]byte, 1024)
		_, err = connecte.Read(recive)
		if err != nil {
			log.Fatalf("Error getting data: %s", err)
		}
	}
}
