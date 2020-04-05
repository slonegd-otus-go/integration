package cmd

import (
	"github.com/spf13/cobra"

	"github.com/slonegd-otus-go/integration/cmd/server"
	"github.com/slonegd-otus-go/integration/cmd/client"
)

var Command = &cobra.Command{
	Use:   "integration",
	Short: "this is grpc demo with integration tests",
}

func init() {
	Command.AddCommand(server.Command)
	Command.AddCommand(client.Command)
}
