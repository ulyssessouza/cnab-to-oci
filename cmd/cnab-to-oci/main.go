package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "cnab-to-oci <subcommand> [options]",
		SilenceUsage: true,
	}
	cmd.AddCommand(fixupCmd(), pushCmd(), pullCmd(), versionCmd())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO(ulyssessouza) Implement versioning")
			return nil
		},
	}
	return cmd
}
