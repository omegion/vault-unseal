package vault

import (
	"fmt"
	"time"

	"github.com/hashicorp/vault/api"
)

const (
	tlsSkipVerify = true
	maxRetries    = 3
	timeout       = 15 * time.Second
)

// API is main struct of Vault.
type API struct {
	Config *api.Config
	Client *api.Client
}

// NewAPI creates AI struct for Vault.
func NewAPI(address string) (API, error) {
	config := api.DefaultConfig()
	config.Address = address
	config.MaxRetries = maxRetries
	config.Timeout = timeout

	err := config.ConfigureTLS(&api.TLSConfig{Insecure: tlsSkipVerify})
	if err != nil {
		return API{}, err
	}

	client, err := api.NewClient(config)
	if err != nil {
		return API{}, err
	}

	return API{
		Config: config,
		Client: client,
	}, nil
}

// SealStatus returns status of seal.
func (a API) SealStatus() (api.SealStatusResponse, error) {
	status, err := a.Client.Sys().SealStatus()
	if err != nil {
		return api.SealStatusResponse{}, err
	}

	return *status, nil
}

// UnsealWithShards unseals the Vault with given shards.
func (a API) UnsealWithShards(shards []string) error {
	for _, shard := range shards {
		status, err := a.SealStatus()
		if err != nil {
			return err
		}

		if status.Sealed {
			status, err = a.Unseal(shard)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("It is unsealed.")
			break
		}
	}

	return nil
}

// Unseal starts to unseal with given shard.
func (a API) Unseal(shard string) (api.SealStatusResponse, error) {
	status, err := a.Client.Sys().Unseal(shard)
	if err != nil {
		return api.SealStatusResponse{}, err
	}

	fmt.Printf("Unsealed with shard: %s\n", shard)

	return *status, nil
}
