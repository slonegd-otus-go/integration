package server

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "server",
	Short: "Run grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		// there will be server
	},
}
