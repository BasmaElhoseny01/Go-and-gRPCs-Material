package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	pb "wireless_lab_1/grpc/capitalize" // Import the generated package

	"google.golang.org/grpc"
)

type textServer struct {
	pb.UnimplementedTextServiceServer
}

func (s *textServer) Capitalize(ctx context.Context, req *pb.TextRequest) (*pb.TextResponse, error) {
	fmt.Println("Serving a client")
	text := req.GetText()
	capitalizedText := strings.ToUpper(text)
	// Sleep for 10 secs
	fmt.Println("Sleeping...")
	time.Sleep(10 * time.Second)

	return &pb.TextResponse{CapitalizedText: capitalizedText}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterTextServiceServer(s, &textServer{})
	fmt.Println("Server started. Listening on port 8080...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
}
