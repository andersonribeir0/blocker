package main

import (
	"net"

	"google.golang.org/grpc"
)

func main() {
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	grpcServer.Serve(ln)
}
