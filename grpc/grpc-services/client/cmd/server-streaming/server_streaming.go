package main

import (
	pb "client/gen/proto"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

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
	// creating the variable of the client interface
	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	stream, err := client.GetPosts(ctx, &pb.GetPostsRequest{UserId: 123})
	if err != nil {
		log.Println(err)
		return
	}

	// we will receive multiple responses from the server
	// we will loop over the stream and receive the responses
	for {
		// receive the first batch of posts
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			// if stream is closed, break the loop
			break
		}
		if err != nil {
			log.Println(err)
			return
		}
		select {
		case <-stream.Context().Done():
			fmt.Println("server cancelled the request or stream is closed")
			return
		default:
			/// server is still alive
		}

		fmt.Println("received response", resp)
		fmt.Println()

	}
}
