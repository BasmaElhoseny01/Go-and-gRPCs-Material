package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// Resolve UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err.Error())
		return
	}

	// Listen for incoming UDP packets
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("Server started. Listening on port 8080...")

	// Buffer for receiving data
	buffer := make([]byte, 1024)

	for {
		// Read data from client
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		fmt.Println("Received data from:", addr)

		// Convert received data to string and capitalize it
		receivedText := strings.TrimSpace(string(buffer[:n]))
		capitalizedText := strings.ToUpper(receivedText)

		// Send capitalized text back to client
		_, err = conn.WriteToUDP([]byte(capitalizedText), addr)
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
	}
}
