package main

import (
	"log"
	"net"
	pb "server/gen/proto"

	"google.golang.org/grpc"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Println(err)
		return
	}

	// Create a gRPC server variable of type *grpc.Server
	s := grpc.NewServer()

	// start the gRPC server on the tcp port 5001
	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}

}
