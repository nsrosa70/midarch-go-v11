package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	protobuf "google.golang.org/grpc/examples/fibonacci/fibonacci"
	"google.golang.org/grpc/reflection"
	"shared/parameters"
	"strconv"
	"fmt"
	"shared/net"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Fibo(ctx context.Context, in *protobuf.FiboRequest) (*protobuf.FiboReply, error) {
	n := F(in.N)
	return &protobuf.FiboReply{N: n}, nil
}

func F(n int32) int32 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return F(n-1) + F(n-2)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(parameters.FIBONACCI_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterGreeterServer(s, &server{})

	// Register reflection service on gRPC server.
	fmt.Println("Fibonacci Server ready at: "+netshared.ResolveHostIp()+" Port: "+strconv.Itoa(parameters.FIBONACCI_PORT))
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
