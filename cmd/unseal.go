package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/omegion/vault-unseal/internal/controller"
	"github.com/omegion/vault-unseal/internal/vault"
)

const splitExpectedParts = 2

// Unseal unseals Vault.
func Unseal() *cobra.Command {
	type flags struct {
		Address       string
		Shard         []string
		TLSSkipVerify bool
		CustomHeaders []string
	}

	cmdFlags := flags{}

	cmd := &cobra.Command{
		Use:   "unseal",
		Short: "Unseal Vault.",
		RunE: func(cmd *cobra.Command, args []string) error {
			address, _ := cmd.Flags().GetString("address")
			shards, _ := cmd.Flags().GetStringSlice("shard")
			customHeaders, _ := cmd.Flags().GetStringSlice("custom-header")
			TLSSkipVerify, _ := cmd.Flags().GetBool("tsl-skip-verify")

			api, err := vault.NewAPI(vault.APIOptions{
				Address:       address,
				TLSSkipVerify: TLSSkipVerify,
				CustomHeaders: getCustomHeadersMap(customHeaders),
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

	cmd.Flags().StringVarP(&cmdFlags.Address,
		"address", "a", "", "Vault server address")
	cmd.Flags().StringSliceVarP(&cmdFlags.Shard,
		"shard", "s", []string{}, "Shard to unseal")
	cmd.Flags().BoolVarP(&cmdFlags.TLSSkipVerify,
		"tls-skip-verify", "t", false, "Skip TLS Verification for Vault server")
	cmd.Flags().StringSliceVarP(&cmdFlags.CustomHeaders,
		"custom-header", "c", []string{}, "Custom header key value; \"key=value\"")

	for _, flag := range []string{"address", "shard"} {
		err := cmd.MarkFlagRequired(flag)
		if err != nil {
			cobra.CheckErr(err)
		}
	}

	return cmd
}

func getCustomHeadersMap(headers []string) map[string]string {
	headersMap := map[string]string{}

	for _, header := range headers {
		parts := strings.SplitN(header, "=", splitExpectedParts)

		key, val := parts[0], parts[1]
		headersMap[key] = val
	}

	return headersMap
}
