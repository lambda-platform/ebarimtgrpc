package server

import (
	"github.com/lambda-platform/ebarimt/posapi"
	ebarimtProto "github.com/lambda-platform/ebarimtgrpc/proto"
	"google.golang.org/grpc"
	"net"
)

type PosApiServer struct {
	ebarimtProto.UnimplementedPosApiServer
	api *posapi.PosAPI
}

func NewPosApiServer(api *posapi.PosAPI) *PosApiServer {
	return &PosApiServer{api: api}
}

// Implement your gRPC methods here, e.g., GetInformation, CheckApi, CallFunction, Put, PutBatch, ReturnBill, SendData.

func StartServer(address string, api *posapi.PosAPI) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	ebarimtProto.RegisterPosApiServer(grpcServer, NewPosApiServer(api))

	return grpcServer.Serve(lis)
}
