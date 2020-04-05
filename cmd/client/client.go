package client

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "client",
	Short: "Run grpc client",
	Run: func(cmd *cobra.Command, args []string) {
		// there will be tests
	},
}
