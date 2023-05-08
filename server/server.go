package server

import (
	"github.com/lambda-platform/ebarimtgrpc/config"
	"google.golang.org/grpc"

	ebarimtProto "github.com/lambda-platform/ebarimtgrpc/proto"
	"log"
	"net"
)

func NewServer() *Server {
	return &Server{}
}

type Server struct {
	ebarimtProto.UnimplementedPosApiServer
}

func (server *Server) Run() error {

	lis, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	ebarimtProto.RegisterPosApiServer(s, server)

	log.Printf("GRPC server listening at %v", lis.Addr())

	return s.Serve(lis)
}

func StartGRPC() {
	go func() {
		var grpcServer *Server = NewServer()
		if err := grpcServer.Run(); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
