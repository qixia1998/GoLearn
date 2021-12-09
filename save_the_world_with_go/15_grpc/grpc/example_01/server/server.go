package main

import (
	"context"
	"fmt"
	pb "GoLearn/save_the_world_with_go/15_grpc/grpc/example_01/user"
	"google.golang.org/grpc"
	"net"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (u *UserServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	fmt.Println("Server received:", req.String())
	return &pb.User{UserId: "John", Email: "john@gmail.com"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserServer{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
