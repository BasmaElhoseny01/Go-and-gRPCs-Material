package main // bec it is executable he added package main
import (
	"context"
	"log"

	pb "grpc-module/server/addressbook"

	"google.golang.org/grpc"
)

func main() {
	// grpc.WithBlock()  --> Blocking requests (Sync) 
	conn,err := grpc.Dial("localhost:50051",grpc.WithInsecure(), grpc.WithBlock())
	
	if err != nil{
		log.Fatal("Could not listen on port")
	}

	defer conn.Close()

	client :=pb.NewAddressBookServiceClient(conn)

	_,err = client.AddPerson(context.Background(),&pb.Person{
		Name:"Basma",
		Id:5,
	});
	if (err != nil){
		log.Fatal("Could not add person")
	}
}