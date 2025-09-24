package main

import (
	pb "client/gen/proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//Connect to the server
	dialOptions := []grpc.DialOption{
		// use insecure credentials, connection is not encrypted
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	// using ...(ellipsis), we can pass a slice to variadic parameter
	// create a connection, connect to the server
	conn, err := grpc.NewClient("localhost:5001", dialOptions...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// create the request

	// create the client
	// this client would be used to call the remote grpc methods
	client := pb.NewUserServiceClient(conn)
	req := pb.SignupRequest{
		User: &pb.User{
			Name:     "test",
			Email:    "test@email.com",
			Password: "hello123",
			Roles:    []string{"admin"}},
	}

	// send the request
	res, err := client.Signup(context.Background(), &req)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("received resp ", res)
	// receive the response
}
