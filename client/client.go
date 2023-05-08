package client

import (
	"context"
	ebarimtProto "github.com/lambda-platform/ebarimtgrpc/proto"
	"google.golang.org/grpc"
	"time"
)

type PosApiClient struct {
	client ebarimtProto.PosApiClient
}

func NewPosApiClient(conn *grpc.ClientConn) *PosApiClient {
	return &PosApiClient{client: ebarimtProto.NewPosApiClient(conn)}
}

// Implement your gRPC client methods here, e.g., GetInformation, CheckApi, CallFunction, Put, PutBatch, ReturnBill, SendData.

func ConnectToServer(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
}

// Example
func Info() (*ebarimtProto.InformationOutput, error) {
	conn, err := ConnectToServer("192.168.67.2:8111")

	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := NewPosApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	in := ebarimtProto.GetInformationRequest{
		PosID: "tatatonga",
	}
	r, err := c.client.GetInformation(ctx, &in)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func Check() (*ebarimtProto.InformationOutput, error) {
	conn, err := ConnectToServer("192.168.67.2:8111")

	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := NewPosApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	in := ebarimtProto.GetInformationRequest{
		PosID: "tatatonga",
	}
	r, err := c.client.GetInformation(ctx, &in)

	if err != nil {
		return nil, err
	}
	return r, nil
}
