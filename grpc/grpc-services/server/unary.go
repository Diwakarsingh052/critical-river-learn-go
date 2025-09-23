package main

import (
	"context"
	// renaming the import to pb
	pb "server/gen/proto"
	// proto is a package that exists in the protobuf repo
	//"github.com/golang/protobuf/proto"
)

func Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	//proto.Marshal()
	return nil, nil
}
