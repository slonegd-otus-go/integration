package client

import (
	"fmt"

	"github.com/slonegd-otus-go/integration/internal/client"
	"github.com/spf13/cobra"
)

var host string
var port int
var a int
var b int

func init() {
	Command.Flags().StringVar(&host, "host", "127.0.0.1", "host to listen")
	Command.Flags().IntVar(&port, "port", 8080, "port to listen")
	Command.Flags().IntVar(&a, "a", 0, "a argument")
	Command.Flags().IntVar(&b, "b", 0, "b argument")
}

var Command = &cobra.Command{
	Use:   "client",
	Short: "Run grpc client",
	Run: func(cmd *cobra.Command, args []string) {

		address := fmt.Sprintf("%s:%d", host, port)
		var res int
		client.Run(cmd.Context(), address, a, b, &res)
	},
}
