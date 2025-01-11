package main

import (
	"fmt"
	"log"
	"net"

	"grpc-module/server"
	pb "grpc-module/server/addressbook"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main(){
	fmt.Println("Hello world")

	lis,err := net.Listen("tcp","localhost:50051")
	if err != nil{
		log.Fatal("Could not listen on port")
	}

	grpcServer :=grpc.NewServer()
	// To use CLI so that ls show he protos for the grpcServer
	// Which is AddressBookService in the .proto
	reflection.Register(grpcServer)
	pb.RegisterAddressBookServiceServer(grpcServer,server.NewAddressBookServer())
	log.Println("Starting server ...")
	grpcServer.Serve(lis)
}

// go mod init <module_name>   in your project directory. Replace <module_name> with the desired module name