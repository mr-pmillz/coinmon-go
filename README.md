# Coinmon-go

Get live Crypto Currency prices in the terminal with the quickness

## Install

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

<img src="https://github.com/mr-pmillz/coinmon-go/blob/master/img/top20.png" />

- Get Specified Coins

```bash
./coinmon-go --find bitcoin,ethereum,cardano,uniswap,dogecoin,chainlink,monero,filecoin,tron,bittorrent
```

<img src="https://github.com/mr-pmillz/coinmon-go/blob/master/img/find.png" />

- By default, `./coinmon-go` with no arguments will just return the top 10 coins by market cap