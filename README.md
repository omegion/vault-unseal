<h1 align="center">
Vault Auto Unseal
</h1>

<p align="center">
  <a href="https://omegion.github.io/vault-unseal-docs/" target="_blank">
    <img width="180" src="https://omegion.github.io/vault-unseal-docs/img/logo.svg" alt="logo">
  </a>
</p>

<p align="center">
    <img src="https://img.shields.io/github/workflow/status/omegion/vault-unseal/Code%20Check" alt="Check"></a>
    <img src="https://coveralls.io/repos/github/omegion/vault-unseal/badge.svg?branch=master" alt="Coverall"></a>
    <img src="https://goreportcard.com/badge/github.com/omegion/vault-unseal" alt="Report"></a>
    <a href="http://pkg.go.dev/github.com/omegion/vault-unseal"><img src="https://img.shields.io/badge/pkg.go.dev-doc-blue" alt="Doc"></a>
    <a href="https://github.com/omegion/vault-unseal/blob/master/LICENSE"><img src="https://img.shields.io/github/license/omegion/vault-unseal" alt="License"></a>
</p>

```shell
CLI command to manage SSH keys stored on Bitwarden

Usage:
  vault-unseal [command]

Available Commands:
  add         Add SSH key to Bitwarden.
  get         Get SSH key from Bitwarden.
  help        Help about any command
  version     Print the version/build number

Flags:
  -h, --help   help for vault-unseal

Use "vault-unseal [command] --help" for more information about a command.

```

## Requirements

* Vault Server

## What does it do?

Injects SSL keys to `ssh-agent` stored in Bitwarden.

## How to use it

1. Login to Bitwarden with `bw`.
1. Create a folder named `SSHKeys` folder in your Bitwarden.
1. Add your key pairs to Bitwarden

```shell
vault-unseal add --name my-server-1 --private-key $PK_PATH --public-key $PUB_KEY
```

## Improvements to be made

* 100% test coverage.
* Better covering for other features.

