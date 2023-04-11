package server

import (
	"context"
	"fmt"
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
	peer, _ := peer.FromContext(ctx)

	n.logger.Info(fmt.Sprintf("Received from: %v\t%v", peer, tx.Version))

	return &proto.Ack{Ok: true}, nil
}

func (n *Node) Handshake(ctx context.Context, version *proto.Version) (*proto.Version, error) {
	peer, _ := peer.FromContext(ctx)

	n.logger.Info(fmt.Sprintf("Received from: %v\t%v", peer, version.Version))

	return &proto.Version{Version: "1.1.1"}, nil
}
