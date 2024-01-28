package cmd

import (
	"context"
	"github.com/ngyewch/protoc-gen-rsocket-go-example/server"
	"github.com/spf13/cobra"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server [flags]",
		Short: "Server",
		RunE:  doServer,
	}
)

func doServer(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	s := server.New()
	return s.Start(ctx)
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
