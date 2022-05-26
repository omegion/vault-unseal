## Installation

You can use `go` to build SSH Manager locally with:

```shell
go install github.com/omegion/vault-unseal@latest
```

Or, you can use the usual commands to install or upgrade:

On OS X

```shell
sudo curl -fL https://github.com/omegion/vault-unseal/releases/download/{{.Env.VERSION}}/vault-unseal-darwin-amd64 -o /usr/local/bin/vault-unseal \
&& sudo chmod +x /usr/local/bin/vault-unseal
```

On Linux

```shell
sudo curl -fL https://github.com/omegion/vault-unseal/releases/download/{{.Env.VERSION}}/vault-unseal-linux-amd64 -o /usr/local/bin/vault-unseal \
&& sudo chmod +x /usr/local/bin/vault-unseal
```

On Windows (Powershell)

```powershell
Invoke-WebRequest -Uri https://github.com/omegion/vault-unseal/releases/download/{{.Env.VERSION}}
/vault-unseal-windows-amd64 
-OutFile $home\AppData\Local\Microsoft\WindowsApps\vault-unseal.exe
```

Otherwise, download one of the releases from the [release page](https://github.com/omegion/vault-unseal/releases/)
directly.

See the install [docs](https://vault-unseal.omegion.dev) for more install options and instructions.

## Changelog