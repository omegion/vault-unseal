package main

import (
	"os"

	"github.com/omegion/vault-unseal/cmd"
)

func main() {
	commander := cmd.Commander{}

	if err := commander.Execute(); err != nil {
		os.Exit(1)
	}
}
