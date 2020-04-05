package client

import (
	"context"
	"log"
	"time"

	"github.com/slonegd-otus-go/integration/internal/grpc_api"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, address string, a, b int) {
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	log.Printf("connect to grpc server: %s", address)

	client := grpc_api.NewCalcClient(connection)

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	result, err := client.Sum(ctx, &grpc_api.SumArgs{
		A: int32(a),
		B: int32(b),
	})
	if err != nil {
		log.Fatalf("sum failed: %s", err)
	}

	log.Printf("got result: %d", result.C)
}
