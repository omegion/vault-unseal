package cmd

import (
	"github.com/spf13/cobra"

	"github.com/omegion/vault-unseal/internal/controller"
	"github.com/omegion/vault-unseal/internal/vault"
)

// setupAddCommand sets default flags.
func setupGetCommand(cmd *cobra.Command) {
	cmd.Flags().String("address", "", "Vault Address")

	if err := cmd.MarkFlagRequired("address"); err != nil {
		cobra.CheckErr(err)
	}

	cmd.Flags().StringSliceP("shard", "s", []string{}, "Shards to unseal")

	if err := cmd.MarkFlagRequired("shard"); err != nil {
		cobra.CheckErr(err)
	}

	cmd.Flags().Bool("tls-skip-verify", true, "Skip TLS Verification for Vault")
}

// Unseal unseals Vault.
func Unseal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unseal",
		Short: "Unseal Vault.",
		RunE: func(cmd *cobra.Command, args []string) error {
			address, _ := cmd.Flags().GetString("address")
			shards, _ := cmd.Flags().GetStringSlice("shard")
			TLSSkipVerify, _ := cmd.Flags().GetBool("tsl-skip-verify")

			api, err := vault.NewAPI(vault.APIOptions{
				Address:       address,
				TLSSkipVerify: TLSSkipVerify,
			})
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
