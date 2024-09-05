package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Buffer to read incoming data
	buffer := make([]byte, 1024)

	// Read data from client [Blocking]
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	// Convert received data to string and capitalize it
	receivedText := strings.TrimSpace(string(buffer[:n]))
	capitalizedText := strings.ToUpper(receivedText)

	// Send capitalized text back to client
	_, err = conn.Write([]byte(capitalizedText))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
		return
	}
}

func main() {
	// Listen for incoming connections on port 8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	// Close Connection at teh end of the main function
	defer listener.Close()
	fmt.Println("Server started. Listening on port 8080...")

	// Accept incoming connections
	for {
		// [Blocking]
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		fmt.Println("Client connected:", conn.RemoteAddr())

		// Handle incoming connection in a separate goroutine
		//threads=routine  (Here we write go then the function the  thread will run <3)
		go handleConnection(conn)
	}
}
