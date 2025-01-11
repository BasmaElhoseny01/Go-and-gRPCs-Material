package server

import (
	"context"
	"log"

	pb "grpc-module/server/addressbook"
)

type addressbookServer struct {
	pb.UnimplementedAddressBookServiceServer
}

// Create constructor for addressbookServer
func NewAddressBookServer() *addressbookServer {
	return &addressbookServer{}
}

func (s *addressbookServer) ListPersons(ctx context.Context, req *pb.ListPersonsRequest) (*pb.ListPersonsResponse ,error){
	log.Println("Listing persons ...")
	return &pb.ListPersonsResponse{},nil
}

func (s *addressbookServer) AddPerson(ctx context.Context, person *pb.Person) (*pb.AddPersonResponse,error){
	log.Println("Adding person ...")
	return &pb.AddPersonResponse{},nil
}