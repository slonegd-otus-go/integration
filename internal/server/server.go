package server

import (
	"context"
	"log"
	"net"

	"github.com/slonegd-otus-go/integration/internal/grpc_api"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	server := grpc.NewServer()
	grpc_api.RegisterCalcServer(server, Calc{})
	log.Printf("start grpc server at %s", address)
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}

type Calc struct{}

func (calc Calc) Sum(ctx context.Context, args *grpc_api.SumArgs) (*grpc_api.SumResult, error) {
	return &grpc_api.SumResult{C: args.A + args.B}, nil
}
