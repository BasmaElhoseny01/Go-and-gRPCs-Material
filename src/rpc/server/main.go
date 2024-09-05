package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strings"
)

// Define a struct to hold the methods that will be exposed via RPC
type TextHandler struct{}

// Method to capitalize the text
func (t *TextHandler) Capitalize(text string, capitalizedText *string) error {
	*capitalizedText = strings.ToUpper(text)
	return nil
}

func main() {
	// Apply RPC on TCP connection :D
	// Register the TextHandler struct with the RPC server
	rpc.Register(new(TextHandler))

	// Listen for incoming TCP connections on port 8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on port 8080...")

	// Accept incoming connections and handle them in a separate goroutine
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		// Serve the Connection 
		go rpc.ServeConn(conn)
	}
}
