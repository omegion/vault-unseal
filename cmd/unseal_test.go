package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Unseal(t *testing.T) {
	_, err := executeCommand(Unseal())

	assert.EqualError(t, err, "required flag(s) \"address\", \"shard\" not set")
}
