<h1 align="center">
Vault Auto Unseal
</h1>

<p align="center">
    <img src="https://img.shields.io/github/workflow/status/omegion/vault-unseal/Code%20Check" alt="Check"></a>
    <img src="https://coveralls.io/repos/github/omegion/vault-unseal/badge.svg?branch=master" alt="Coverall"></a>
    <img src="https://goreportcard.com/badge/github.com/omegion/vault-unseal" alt="Report"></a>
    <a href="http://pkg.go.dev/github.com/omegion/vault-unseal"><img src="https://img.shields.io/badge/pkg.go.dev-doc-blue" alt="Doc"></a>
    <a href="https://github.com/omegion/vault-unseal/blob/master/LICENSE"><img src="https://img.shields.io/github/license/omegion/vault-unseal" alt="License"></a>
</p>

```shell
CLI command to automatically unseal Vault

Usage:
  vault-unseal [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  unseal      Unseal Vault.
  version     Print the version/build number

Flags:
      --config string      config file (default is ~/.<YOUR_CONFIG_NAME>/<CONFIG_FILE_NAME>.<CONFIG_FILE_TYPE>)
  -h, --help               help for vault-unseal
      --logFormat string   Set the logging format. One of: text|json (default "text") (default "text")
      --logLevel string    Set the logging level. One of: debug|info|warn|error (default "info")

Use "vault-unseal [command] --help" for more information about a command.

```

## Requirements

* Vault Server

## What does it do?

Its a tool to unseal your Vault Server with given shards.

## How to use it

1. Run `unseal` command with your Vault address and shards.

```shell
vault-unseal unseal --address https://my.vault.server \
  --shard=<SHARD_1>
  --shard=<SHARD_2>
  --shard=<SHARD_3>
```

## Vault Unseal GitHub Action

You can use [Vault Unseal GitHub Action](https://github.com/omegion/vault-unseal-action) to create periodical action to be sure that your Vault is always unsealed.

## Improvements to be made

* 100% test coverage.
* Better covering for other features.

