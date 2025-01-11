package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
)

func main() {
	// Dial the RPC server at localhost:8080
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer client.Close()

	// Read input from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to capitalize: ")
	text, _ := reader.ReadString('\n')

	// Call the remote method "TextHandler.Capitalize" on the server
	var capitalizedText string
	err = client.Call("TextHandler.Capitalize", text, &capitalizedText)
	if err != nil {
		fmt.Println("Error calling remote method:", err)
		return
	}

	// Print the capitalized text received from the server
	fmt.Println("Capitalized text:", capitalizedText)
}
