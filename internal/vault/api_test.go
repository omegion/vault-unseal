package vault

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	testCases := []struct {
		apiOptions            APIOptions
		expectedTLSSkipVerify bool
	}{
		{
			APIOptions{
				TLSSkipVerify: false,
			},
			false,
		}, {
			APIOptions{
				TLSSkipVerify: true,
			},
			true,
		},
	}

	for _, test := range testCases {
		api, err := NewAPI(test.apiOptions)
		assert.NoError(t, err)

		transport, ok := api.Config.HttpClient.Transport.(*http.Transport)
		assert.Equal(t, true, ok)

		assert.Equal(t, test.expectedTLSSkipVerify, transport.TLSClientConfig.InsecureSkipVerify)
	}
}
