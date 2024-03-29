# Coinmon-go

[![Go Report Card](https://goreportcard.com/badge/github.com/mr-pmillz/coinmon-go)](https://goreportcard.com/report/github.com/mr-pmillz/coinmon-go)
![GitHub all releases](https://img.shields.io/github/downloads/mr-pmillz/coinmon-go/total?style=social)
![GitHub repo size](https://img.shields.io/github/repo-size/mr-pmillz/coinmon-go?style=plastic)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mr-pmillz/coinmon-go?style=plastic)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/mr-pmillz/coinmon-go?style=plastic)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/mr-pmillz/coinmon-go?style=plastic)
[![codecov](https://codecov.io/gh/mr-pmillz/coinmon-go/branch/master/graph/badge.svg?token=1O7CY7MD6U)](https://codecov.io/gh/mr-pmillz/coinmon-go)

Table of Contents
=================

* [Coinmon\-go](#coinmon-go)
    * [Install](#install)
    * [Usage](#usage)

Get live Crypto Currency prices in the terminal with the quickness

## Install

If you have a version of golang >= 16.X you can install with

```shell
go install github.com/mr-pmillz/coinmon-go@latest
```

If using an older version of golang

```shell
go get github.com/mr-pmillz/coinmon-go@latest
```

To build from source, clone the repo and run

```bash
go build
```

## Coinmon-go now uses the CoinMarketCap API

- [Create a Free API Key](https://coinmarketcap.com/api/documentation/v1/#section/Quick-Start-Guide)
- add it as an env var (config file support coming soon)
- `COINMARKETCAP_API_KEY=foo-bar-baz`

Reason for switching is the data is more accurate than the coincap api and cmc api has more rich functionality.
If you don't add the COINMARKETCAP_API_KEY to your env vars, coin prices wont show.

## Usage

```bash
./coinmon-go -h
Options:

  -h, --help       display help information
  -t, --top[=10]   Amount of coins to show data for, Default is top 10, If -f|--find flag supplied, -t|--top is ignored
```

- Get Top 20 Coins

```bash
./coinmon-go -t 20
```


- By default, `./coinmon-go` with no arguments will just return the top 10 coins by market cap
