package client

import (
	ebarimtProto "github.com/lambda-platform/ebarimtgrpc/proto"
	"google.golang.org/grpc"
)

type PosApiClient struct {
	client ebarimtProto.PosApiClient
}

func NewPosApiClient(conn *grpc.ClientConn) *PosApiClient {
	return &PosApiClient{client: ebarimtProto.NewPosApiClient(conn)}
}

// Implement your gRPC client methods here, e.g., GetInformation, CheckApi, CallFunction, Put, PutBatch, ReturnBill, SendData.

func ConnectToServer(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
}
