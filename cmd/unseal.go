package cmd

import (
	"log"

	"github.com/omegion/vault-unseal/pkg/vault"

	"github.com/spf13/cobra"
)

// setupAddCommand sets default flags.
func setupGetCommand(cmd *cobra.Command) {
	cmd.Flags().String("address", "", "Vault Address")

	if err := cmd.MarkFlagRequired("address"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().StringArray("shard", []string{}, "Shards to unseal")

	if err := cmd.MarkFlagRequired("shard"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
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

			err = api.UnsealWithShards(shards)
			if err != nil {
				return err
			}

			return nil
		},
	}

	setupGetCommand(cmd)

	return cmd
}
