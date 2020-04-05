package tests

import (
	"context"
	"testing"

	"github.com/slonegd-otus-go/integration/internal/client"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{
			name: "2 + 3",
			a:    2,
			b:    3,
			want: 5,
		},
	}
	for _, tt := range tests {
		var got int
		client.Run(context.Background(), "127.0.0.1:8080", tt.a, tt.b, &got)

		assert.Equal(t, tt.want, got)
	}
}
