package main

import (
	"context"
	"fmt"

	pb "wireless_lab_1/grpc/capitalize" // Import the generated package

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect:", err)
		return
	}
	defer conn.Close()
	c := pb.NewTextServiceClient(conn)

	// Read input from user
	fmt.Print("Enter text to capitalize: ")
	var text string
	fmt.Scanln(&text)

	// Call the RPC method
	resp, err := c.Capitalize(context.Background(), &pb.TextRequest{Text: text})
	if err != nil {
		fmt.Println("Error calling Capitalize:", err)
		return
	}

	// Print the result
	fmt.Println("Capitalized text:", resp.GetCapitalizedText())
}
