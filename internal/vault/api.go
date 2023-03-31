package vault

import (
	"fmt"
	"time"

	"github.com/hashicorp/vault/api"
	log "github.com/sirupsen/logrus"
)

const (
	maxRetries = 3
	timeout    = 15 * time.Second
)

// APIInterface is an interface for API.
//

//go:generate mockgen -destination=mocks/api_mock.go -package=mocks github.com/omegion/vault-unseal/internal/vault APIInterface
type APIInterface interface {
	SealStatus() (api.SealStatusResponse, error)
	Unseal(shard string) (api.SealStatusResponse, error)
}

// API is main struct of Vault.
type API struct {
	Config *api.Config
	Client *api.Client
}

// APIOptions is options for API.
type APIOptions struct {
	Address       string
	TLSSkipVerify bool
	CustomHeaders map[string]string
}

// NewAPI creates AI struct for Vault.
func NewAPI(options APIOptions) (API, error) {
	config := api.DefaultConfig()
	config.Address = options.Address
	config.MaxRetries = maxRetries
	config.Timeout = timeout

	err := config.ConfigureTLS(&api.TLSConfig{Insecure: options.TLSSkipVerify})
	if err != nil {
		return API{}, err
	}

	client, err := api.NewClient(config)
	if err != nil {
		return API{}, err
	}

	for key, val := range options.CustomHeaders {
		client.AddHeader(key, val)
	}

	return API{
		Config: config,
		Client: client,
	}, nil
}

// SealStatus returns status of seal.
func (a API) SealStatus() (api.SealStatusResponse, error) {
	log.Debugln("Getting Vault seal status.")

	status, err := a.Client.Sys().SealStatus()
	if err != nil {
		return api.SealStatusResponse{}, err
	}

	return *status, nil
}

// Unseal starts to unseal with given shard.
func (a API) Unseal(shard string) (api.SealStatusResponse, error) {
	log.Debugln(fmt.Sprintf("Unsealing with shard: %s", shard))

	status, err := a.Client.Sys().Unseal(shard)
	if err != nil {
		return api.SealStatusResponse{}, err
	}

	log.Infoln(fmt.Sprintf("Unsealed with shard: %s", shard))

	return *status, nil
}
