package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Resolve UDP server address
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err.Error())
		return
	}

	// Establish UDP connection
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	// Read input from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to capitalize: ")
	text, _ := reader.ReadString('\n')

	// Send input text to server
	_, err = conn.Write([]byte(strings.TrimSpace(text)))
	if err != nil {
		fmt.Println("Error sending data:", err.Error())
		return
	}

	// Buffer for receiving response
	buffer := make([]byte, 1024)

	// Read response from server
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error receiving data:", err.Error())
		return
	}

	// Print capitalized text received from server
	fmt.Println("Capitalized text:", string(buffer[:n]))
}
