package main

import (
	"context"
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

	ln, err := net.Listen("tcp", ":3005")
	if err != nil {
		panic(err)
	}

	proto.RegisterNodeServer(grpcServer, node)
	grpcServer.Serve(ln)
}

func makeRequest(server *proto.NodeServer) {
	conn, err := grpc.Dial(":3005", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewNodeClient(conn)

	client.Handshake(context.TODO(), &proto.Version{Version: "blocker-0.0.1", Height: 1})

	client.HandleTransaction(context.TODO(), &proto.Transaction{Version: "blocker-0.1.1"})

	conn.Close()
}
