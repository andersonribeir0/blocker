package main

import (
	"net"

	"github.com/andersonribeir0/blocker/proto"
	"github.com/andersonribeir0/blocker/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logConfig := zap.NewProductionConfig()
	logConfig.DisableCaller = true
	logConfig.Level.SetLevel(zap.DebugLevel)
	logger, _ := logConfig.Build()

	node := server.NewNode(logger)

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	proto.RegisterNodeServer(grpcServer, node)
	grpcServer.Serve(ln)
}
