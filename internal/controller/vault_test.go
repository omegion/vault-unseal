package controller

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hashicorp/vault/api"
	"github.com/omegion/vault-unseal/internal/vault/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAPI_Unseal_Successful(t *testing.T) {
	ctrl := gomock.NewController(t)
	vaultAPI := mocks.NewMockAPIInterface(ctrl)

	vaultAPI.EXPECT().SealStatus().Return(api.SealStatusResponse{Sealed: true}, nil)
	vaultAPI.EXPECT().Unseal(gomock.Any()).Return(api.SealStatusResponse{Sealed: false}, nil)

	c := NewVaultController(vaultAPI)

	err := c.Unseal(UnsealOptions{
		Shards: []string{
			"test-1",
		},
	})

	assert.Equal(t, nil, err)
}

func TestAPI_Unseal_SealStatus_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	vaultAPI := mocks.NewMockAPIInterface(ctrl)

	vaultAPI.EXPECT().SealStatus().Return(
		api.SealStatusResponse{Sealed: false},
		errors.New("error"),
	).MaxTimes(1)

	c := NewVaultController(vaultAPI)

	err := c.Unseal(UnsealOptions{
		Shards: []string{
			"test-1",
		},
	})

	assert.Error(t, err)
}

func TestAPI_Unseal_Unseal_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	vaultAPI := mocks.NewMockAPIInterface(ctrl)

	vaultAPI.EXPECT().SealStatus().Return(api.SealStatusResponse{Sealed: true}, nil)
	vaultAPI.EXPECT().Unseal(gomock.Any()).Return(
		api.SealStatusResponse{Sealed: false},
		errors.New("error"),
	)

	c := NewVaultController(vaultAPI)

	err := c.Unseal(UnsealOptions{
		Shards: []string{
			"test-1",
		},
	})

	assert.Error(t, err)
}

func TestAPI_Unseal_Unseal_Break(t *testing.T) {
	ctrl := gomock.NewController(t)
	vaultAPI := mocks.NewMockAPIInterface(ctrl)

	vaultAPI.EXPECT().SealStatus().Return(api.SealStatusResponse{Sealed: false}, nil)
	vaultAPI.EXPECT().Unseal(gomock.Any()).Return(
		api.SealStatusResponse{Sealed: false},
		errors.New("error"),
	)

	c := NewVaultController(vaultAPI)

	err := c.Unseal(UnsealOptions{
		Shards: []string{
			"test-1",
		},
	})

	assert.NoError(t, err)
}
