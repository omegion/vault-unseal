package main

import (
	"github.com/omegion/vault-unseal/cmd"
	"os"

	"github.com/spf13/cobra"
)

// RootCommand is the main entry point of this application.
func RootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:          "vault-unseal",
		Short:        "Vault Auto Unseal",
		Long:         "CLI command to automatically unseal Vault",
		SilenceUsage: true,
	}

	root.AddCommand(cmd.Version())
	root.AddCommand(cmd.Unseal())

	return root
}

func main() {
	if err := RootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
