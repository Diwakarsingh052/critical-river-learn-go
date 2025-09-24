package main

import (
	"log"
	"net"
	pb "server/gen/proto"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserService struct {
	// add dependencies for grpc methods,
	v *validator.Validate
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{v: validator.New()}
}
func main() {

	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Println(err)
		return
	}

	// Create a gRPC server variable of type *grpc.Server
	s := grpc.NewServer()
	us := NewUserService()
	pb.RegisterUserServiceServer(s, us)

	//exposing gRPC service methods to be tested by postman
	reflection.Register(s)
	// start the gRPC server on the tcp port 5001
	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}

}
