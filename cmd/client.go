package cmd

import (
	"context"
	"github.com/ngyewch/protoc-gen-rsocket-go-example/client"
	"github.com/spf13/cobra"
)

var (
	clientCmd = &cobra.Command{
		Use:   "client [flags]",
		Short: "Client",
		RunE:  doClient,
	}
)

func doClient(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	c := client.New()
	return c.Start(ctx)
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
