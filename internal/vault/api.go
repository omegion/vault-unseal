package vault

import (
	"fmt"
	"time"

	"github.com/hashicorp/vault/api"

	log "github.com/sirupsen/logrus"
)

const (
	tlsSkipVerify = true
	maxRetries    = 3
	timeout       = 15 * time.Second
)

//nolint:lll // go generate is ugly.
//go:generate mockgen -destination=mocks/api_mock.go -package=mocks github.com/omegion/vault-unseal/internal/vault APIInterface
// APIInterface is an interface for API.
type APIInterface interface {
	SealStatus() (api.SealStatusResponse, error)
	Unseal(shard string) (api.SealStatusResponse, error)
}

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

// Unseal starts to unseal with given shard.
func (a API) Unseal(shard string) (api.SealStatusResponse, error) {
	status, err := a.Client.Sys().Unseal(shard)
	if err != nil {
		return api.SealStatusResponse{}, err
	}

	log.Infoln(fmt.Sprintf("Unsealed with shard: %s\n", shard))

	return *status, nil
}
