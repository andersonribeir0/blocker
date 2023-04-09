package main

import (
	"net"

	"github.com/andersonribeir0/blocker/proto"
	"github.com/andersonribeir0/blocker/server"
	"google.golang.org/grpc"
)

func main() {
	node := server.Node{}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	proto.RegisterNodeServer(grpcServer, node)
	grpcServer.Serve(ln)
}
