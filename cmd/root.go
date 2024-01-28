package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "protoc-gen-rsocket-go-example [flags]",
		Short: "protoc-gen-rsocket-go-example",
		RunE:  help,
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func help(cmd *cobra.Command, args []string) error {
	err := cmd.Help()
	if err != nil {
		return err
	}
	return nil
}
