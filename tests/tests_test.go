package tests

import (
	"context"
	"log"
	"testing"

	docker_client "github.com/docker/docker/client"

	"github.com/slonegd-otus-go/integration/internal/client"
	"github.com/slonegd-otus-go/integration/tests/docker"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	cli, err := docker_client.NewEnvClient()
	if err != nil {
		log.Fatalf("new docker client failed: %s", err)
	}
	id, err := docker.StartNewContainer(cli, "server")
	if err != nil {
		log.Fatalf("start container failed: %s", err)
	}
	log.Printf("start container")

	tests := []struct {
		name string
		a, b int
		want int
	}{
		{
			name: "2 + 3 = 5",
			a:    2,
			b:    3,
			want: 5,
		},
		{
			name: "1 + 2 = 3",
			a:    1,
			b:    2,
			want: 3,
		},
	}
	for _, tt := range tests {
		var got int
		client.Run(context.Background(), "127.0.0.1:8080", tt.a, tt.b, &got)

		assert.Equal(t, tt.want, got)
	}

	err = docker.StopContainer(cli, id)
	if err != nil {
		log.Fatalf("stop container failed: %s", err)
	}
	log.Printf("stop container")
}
