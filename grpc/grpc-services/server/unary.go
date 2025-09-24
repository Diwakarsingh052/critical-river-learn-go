package main

import (
	"context"
	"fmt"
	"server/models"

	// renaming the import to pb
	pb "server/gen/proto"
	// proto is a package that exists in the protobuf repo
	//"github.com/golang/protobuf/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (us *UserService) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	var u models.User
	user := req.GetUser()
	u.Name = user.GetName()
	u.Email = user.GetEmail()
	u.Password = user.GetPassword()
	u.Roles = user.GetRoles()

	err := us.v.Struct(u)
	if err != nil {
		// while returning error from grpc methods, use status.Errorf
		return nil, status.Errorf(codes.InvalidArgument, "provide values in correct format %v", err)
	}

	// call the business logic after validating the request
	//user.CreateUser(u)
	fmt.Println("received request and processed", u)

	return &pb.SignupResponse{Result: "User created Successfully"}, nil
}
