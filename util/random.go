package util

import (
	cryptoRand "crypto/rand"
	"io"
	"math/rand"
	"time"

	"github.com/andersonribeir0/blocker/proto"
)

func RandomHash() []byte {
	hash := make([]byte, 32)
	io.ReadFull(cryptoRand.Reader, hash)

	return hash
}

func RandomBlock() *proto.Block {
	header := &proto.Header{
		Version:   1,
		Height:    int32(rand.Intn(1000)),
		PrevHash:  RandomHash(),
		RootHash:  RandomHash(),
		Timestamp: time.Now().UnixNano(),
	}

	return &proto.Block{
		Header: header,
	}
}
