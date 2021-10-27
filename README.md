# Coinmon-go

[![Donate](https://img.shields.io/badge/Donate-PayPal-yellow.svg)](https://www.paypal.com/donate?business=YR6C4WB5CDZZL&no_recurring=0&item_name=contribute+to+open+source&currency_code=USD)
[![Donate with Bitcoin](https://en.cryptobadges.io/badge/micro/3Cd54T1EB6WHRcechq1dRCGF6vY2HHhkdk)](https://en.cryptobadges.io/donate/3Cd54T1EB6WHRcechq1dRCGF6vY2HHhkdk)
[![Donate with Ethereum](https://en.cryptobadges.io/badge/micro/0x064AA753EF36e5641E2Ee3C9BbC117F6aFe35F62)](https://en.cryptobadges.io/donate/0x064AA753EF36e5641E2Ee3C9BbC117F6aFe35F62)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/12750b2a6fa44b429f8a3f64d3e7f1b6)](https://www.codacy.com/gh/mr-pmillz/coinmon-go/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mr-pmillz/coinmon-go&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/mr-pmillz/coinmon-go)](https://goreportcard.com/report/github.com/mr-pmillz/coinmon-go)
![GitHub all releases](https://img.shields.io/github/downloads/mr-pmillz/coinmon-go/total?style=social)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/mr-pmillz/coinmon-go/CI?style=plastic)
![GitHub repo size](https://img.shields.io/github/repo-size/mr-pmillz/coinmon-go?style=plastic)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mr-pmillz/coinmon-go?style=plastic)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/mr-pmillz/coinmon-go?style=plastic)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/mr-pmillz/coinmon-go?style=plastic)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/12750b2a6fa44b429f8a3f64d3e7f1b6)](https://www.codacy.com/gh/mr-pmillz/coinmon-go/dashboard?utm_source=github.com&utm_medium=referral&utm_content=mr-pmillz/coinmon-go&utm_campaign=Badge_Coverage)

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

## Usage

```bash
./coinmon-go -h
Options:

  -h, --help       display help information
  -t, --top[=10]   Amount of coins to show data for, Default is top 10, If -f|--find flag supplied, -t|--top is ignored
  -f, --find       Specific Coins to return. Example: bitcoin,cardano,ethereum,uniswap
```

- Get Top 20 Coins

```bash
./coinmon-go -t 20
```

![top20.png](https://github.com/mr-pmillz/coinmon-go/blob/master/img/top20.png?raw=true)

- Get Specified Coins

```bash
./coinmon-go --find bitcoin,ethereum,cardano,uniswap,dogecoin,chainlink,monero,filecoin,tron,bittorrent
```

![find.png](https://github.com/mr-pmillz/coinmon-go/blob/master/img/find.png?raw=true)

- By default, `./coinmon-go` with no arguments will just return the top 10 coins by market cap