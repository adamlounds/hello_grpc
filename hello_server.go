//go:generate protoc -I protobufs --go_out=plugins=grpc:protobufs protobufs/grpc_hello.proto

package main

import (
	"fmt"
	pb "github.com/adamlounds/grpc_hello/protobufs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("cannot serve: %v", err)
	}
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("helloworld\n")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
