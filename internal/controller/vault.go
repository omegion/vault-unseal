package controller

import (
	"github.com/omegion/vault-unseal/internal/vault"

	log "github.com/sirupsen/logrus"
)

// VaultController is the entry controller.
type VaultController struct {
	API *vault.APIInterface
}

// NewVaultController is a factory for VaultController.
func NewVaultController(api vault.APIInterface) *VaultController {
	return &VaultController{API: &api}
}

// UnsealOptions is options for unsealing.
type UnsealOptions struct {
	Shards []string
}

// Unseal unseals the Vault.
func (c VaultController) Unseal(options UnsealOptions) error {
	api := *c.API

	for _, shard := range options.Shards {
		status, err := api.SealStatus()
		if err != nil {
			return err
		}

		if status.Sealed {
			status, err = api.Unseal(shard)
			if err != nil {
				return err
			}
		} else {
			log.Infoln("It is unsealed.")
			break
		}
	}

	return nil
}
