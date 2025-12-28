package main

import (
	"fmt"
	"net"
)

var username string
var message string

func main() {

	var addr string

	fmt.Println("--CLIENT-STARTED--")

	fmt.Print("Enter the username: \n")
	fmt.Scan(&username)
	fmt.Print("Enter the address like ip:port\n")
	fmt.Scan(&addr)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("Error connecting: %v", err)
	}

	conn.Write([]byte(username))

	for {

		fmt.Print("Enter message to send: \n")
		fmt.Scan(&message)

		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error sending data: %v", err)
		}

		fmt.Printf("%v: %v", username, message)

		go handleconnection(conn)

	}

}

func handleconnection(conn net.Conn) {
	defer handleconnection(conn)
	buffer := make([]byte, 1024)
	for {
		read, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Error getting message: %v", err)
		}
		fmt.Printf("%v: %v", username, string(buffer[:read]))
	}
}
