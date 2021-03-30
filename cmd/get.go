package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// Unseal unseals Vault.
func Unseal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unseal",
		Short: "Unseal Vault.",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")

			fmt.Printf("Item %s not found\n", name)

			return nil
		},
	}

	cmd.Flags().String("name", "", "Name")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("now", "", "Unseals immediately")

	return cmd
}
