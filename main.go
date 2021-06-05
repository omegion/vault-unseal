package main

import (
	"os"

	commander "github.com/omegion/cobra-commander"
	"github.com/omegion/vault-unseal/cmd"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:          "vault-unseal",
		Short:        "Vault Auto Unseal",
		Long:         "CLI command to automatically unseal Vault",
		SilenceUsage: true,
	}

	c := commander.NewCommander(root).
		SetCommand(
			cmd.Version(),
			cmd.Unseal(),
		).
		Init()

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
