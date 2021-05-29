package cmd

import (
	"github.com/omegion/vault-unseal/internal/controller"
	"github.com/omegion/vault-unseal/internal/vault"

	"github.com/spf13/cobra"
)

// setupAddCommand sets default flags.
func setupGetCommand(cmd *cobra.Command) {
	cmd.Flags().String("address", "", "Vault Address")

	if err := cmd.MarkFlagRequired("address"); err != nil {
		cobra.CheckErr(err)
	}

	cmd.Flags().StringArray("shard", []string{}, "Shards to unseal")

	if err := cmd.MarkFlagRequired("shard"); err != nil {
		cobra.CheckErr(err)
	}
}

// Unseal unseals Vault.
func Unseal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unseal",
		Short: "Unseal Vault.",
		RunE: func(cmd *cobra.Command, args []string) error {
			address, _ := cmd.Flags().GetString("address")
			shards, _ := cmd.Flags().GetStringArray("shard")

			api, err := vault.NewAPI(address)
			if err != nil {
				return err
			}

			vaultController := controller.NewVaultController(&api)

			return vaultController.Unseal(
				controller.UnsealOptions{
					Shards: shards,
				},
			)
		},
	}

	setupGetCommand(cmd)

	return cmd
}
