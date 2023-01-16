# currency-cli

This app is a example to use Makefile and Mage to compile Go projects.

# CLI Usage

```bash
Currency CLI é uma CLI para consultar a cotação de várias moedas. 
Exemplo: 
$ currency-cli eur
Cotação agora do Euro: R$ 5.5303

Usage:
  currency-cli [command]

Available Commands:
  btc         Cotação agora do Bitcoin.
  cad         Cotação agora do Dolár Canadense.
  completion  Generate the autocompletion script for the specified shell
  doge        Cotação agora do Dogecoin.
  eur         Cotação agora do Euro.
  gbp         Cotação agora da Libra Esterlina
  help        Help about any command
  usd         Cotação agora do Dolár.

Flags:
  -h, --help     help for currency-cli
  -t, --toggle   Mensagem de ajuda para o toogle

Use "currency-cli [command] --help" for more information about a command.
```

# Make

### Build

```bash
$ make build
$ ./bin/currency-cli usd #Example
```

### Compile
```bash
$ make compile
Compliling for every OS and Platform
Compliling to MacOS...
Compliling to FreeBSD...
Compliling to Linux...
Compliling to Windows...
```

### Clean
```bash
bin/currency-cli-freebsd-386
bin/currency-cli-freebsd-amd64
bin/currency-cli-linux-386
bin/currency-cli-linux-amd64
bin/currency-cli-macos-amd64
bin/currency-cli-macos-arm64
bin/currency-cli-windows-386
bin/currency-cli-windows-amd64
```

# Mage

Under implementation!