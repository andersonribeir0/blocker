package server

import (
	"context"
	"net"

	"github.com/andersonribeir0/blocker/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Node struct {
	proto.UnimplementedNodeServer

	logger *zap.Logger
	peers  map[net.Addr]*grpc.ClientConn
}

func NewNode(logger *zap.Logger) *Node {
	return &Node{
		logger: logger,
	}
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.None, error) {
	return nil, nil
}
