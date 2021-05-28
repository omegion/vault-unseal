package controller

import (
	"testing"

	"github.com/hashicorp/vault/api"
	"github.com/omegion/vault-unseal/internal/vault/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAPI_SealStatus(t *testing.T) {
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
