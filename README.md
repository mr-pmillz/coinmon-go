# Coinmon-go

Get Crypto Currency prices in the terminal with the quickness

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

- Get Top 15 Coins

```bash
./coinmon-go -t 15
```

<img src="https://github.com/mr-pmillz/coinmon-go/blob/master/img/top20.png" />

- Get Specified Coins

```bash
./coinmon-go -f bitcoin,ethereum,cardano,uniswap,dogecoin
```

<img src="https://github.com/mr-pmillz/coinmon-go/blob/master/img/find.png" />