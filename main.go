package main

import (
	"os"

	commander "github.com/omegion/cobra-commander"
	"github.com/omegion/vault-unseal/cmd"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	// Config file name where a config file will be created.
	// For example, $HOME/.vault-unseal/config.yaml.
	configFileName = "vault-unseal"

	// The environment variable prefix of all environment variables bound to our command line flags.
	// For example, --address is bound to VUNSEAL_ADDRESS.
	configEnvPrefix = "VUNSEAL"
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
		)
	c.SetConfig(getConfig(c.Root.Flags())).Init()

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}

func getConfig(flags *pflag.FlagSet) *commander.Config {
	configName := configFileName
	environmentPrefix := configEnvPrefix

	return &commander.Config{
		Name:              &configName,
		EnvironmentPrefix: &environmentPrefix,
		Flags:             flags,
	}
}
