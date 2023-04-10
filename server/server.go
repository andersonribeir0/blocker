package server

import (
	"context"
	"net"

	"github.com/andersonribeir0/blocker/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
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

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, err := peer.FromContext(ctx)

	n.logger.Info("Received from: %s\t%v", peer, tx.Version)
	return &proto.Ack{ok: true}, err
}
