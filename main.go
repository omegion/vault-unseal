package main

import (
	"github.com/omegion/vault-unseal/cmd"
	"os"
)

func main() {
	commander := cmd.Commander{}

	if err := commander.Execute(); err != nil {
		os.Exit(1)
	}
}
