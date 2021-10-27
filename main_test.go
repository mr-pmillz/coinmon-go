package main

import (
	"encoding/json"
	"strings"
	"testing"
)

const rawData = `{"data":[{"id":"bitcoin","rank":"1","symbol":"BTC","name":"Bitcoin","supply":"18855306.0000000000000000","maxSupply":"21000000.0000000000000000","marketCapUsd":"1137176492787.6108331328990334","volumeUsd24Hr":"18687028375.6651281279008956","priceUsd":"60310.6888208343493939","changePercent24Hr":"-4.3370911053518709","vwap24Hr":"61967.9294595589205303","explorer":"https://blockchain.info/"},{"id":"ethereum","rank":"2","symbol":"ETH","name":"Ethereum","supply":"118089599.7490000000000000","maxSupply":null,"marketCapUsd":"487694290607.3243381700760187","volumeUsd24Hr":"8793930845.2528218020785789","priceUsd":"4129.8665728728088499","changePercent24Hr":"-2.0061025911046535","vwap24Hr":"4202.9589825799117033","explorer":"https://etherscan.io/"},{"id":"binance-coin","rank":"3","symbol":"BNB","name":"Binance Coin","supply":"166801148.0000000000000000","maxSupply":"166801148.0000000000000000","marketCapUsd":"79666706419.0832257472907276","volumeUsd24Hr":"842700214.4413421823133572","priceUsd":"477.6148568179112637","changePercent24Hr":"-1.4874930394008333","vwap24Hr":"484.5904963369062042","explorer":"https://etherscan.io/token/0xB8c77482e45F1F44dE1745F52C74426C631bDD52"},{"id":"cardano","rank":"4","symbol":"ADA","name":"Cardano","supply":"33258168699.5310000000000000","maxSupply":"45000000000.0000000000000000","marketCapUsd":"71132528082.9146648956862663","volumeUsd24Hr":"1219975849.2584226365046248","priceUsd":"2.1387987031263619","changePercent24Hr":"-0.3470502412514219","vwap24Hr":"2.1703121516156345","explorer":"https://cardanoexplorer.com/"},{"id":"tether","rank":"5","symbol":"USDT","name":"Tether","supply":"69663609914.1070700000000000","maxSupply":null,"marketCapUsd":"69679818291.0731586052256609","volumeUsd24Hr":"45922758891.2681863172556371","priceUsd":"1.0002326663373614","changePercent24Hr":"0.0039927020598999","vwap24Hr":"0.9999550075476314","explorer":"https://www.omniexplorer.info/asset/31"},{"id":"solana","rank":"6","symbol":"SOL","name":"Solana","supply":"301095835.7624625600000000","maxSupply":null,"marketCapUsd":"59984379324.4837688927071085","volumeUsd24Hr":"752760032.6985592494853863","priceUsd":"199.2202222677235247","changePercent24Hr":"-4.7090141804476989","vwap24Hr":"207.2249564422103869","explorer":"https://explorer.solana.com/"},{"id":"xrp","rank":"7","symbol":"XRP","name":"XRP","supply":"45404028640.0000000000000000","maxSupply":"100000000000.0000000000000000","marketCapUsd":"50165986624.5641102845452160","volumeUsd24Hr":"2204528899.1464717477929469","priceUsd":"1.1048796357327844","changePercent24Hr":"0.8857520734115551","vwap24Hr":"1.1172707016032134","explorer":"https://xrpcharts.ripple.com/#/graph/"},{"id":"polkadot","rank":"8","symbol":"DOT","name":"Polkadot","supply":"1044027894.5256100000000000","maxSupply":null,"marketCapUsd":"46917029981.1538592033009432","volumeUsd24Hr":"690277443.4266932660075236","priceUsd":"44.9384831834136254","changePercent24Hr":"1.0636276819510248","vwap24Hr":"44.4042110481044716","explorer":"https://polkascan.io/polkadot"},{"id":"dogecoin","rank":"9","symbol":"DOGE","name":"Dogecoin","supply":"131861817247.6350900000000000","maxSupply":null,"marketCapUsd":"33672487947.2963476878307151","volumeUsd24Hr":"1284645248.4417108191855590","priceUsd":"0.2553619284956446","changePercent24Hr":"-3.9418578392844293","vwap24Hr":"0.2662957661907547","explorer":"http://dogechain.info/chain/Dogecoin"},{"id":"usd-coin","rank":"10","symbol":"USDC","name":"USD Coin","supply":"32509827232.2937470000000000","maxSupply":null,"marketCapUsd":"32509803296.0226941667642316","volumeUsd24Hr":"705140395.9593562466598831","priceUsd":"0.9999992637219853","changePercent24Hr":"-0.0224547154094733","vwap24Hr":"1.0003773984768458","explorer":"https://etherscan.io/token/0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"}],"timestamp":1635293152930}`

func Test_getJSON(t *testing.T) {
	type args struct {
		url    string
		target interface{}
	}
	coinData := new(CoinData)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test_getJSON 1", args: args{
			url:    "https://api.coincap.io/v2/assets?limit=10",
			target: coinData,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getJSON(tt.args.url, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("getJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_prettyFloatString(t *testing.T) {
	type args struct {
		num                string
		percent            bool
		nearestThousandFMT bool
		prec4              bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Test_prettyFloatString 1", args: args{
			num:                "2.1387987031263619",
			percent:            false,
			nearestThousandFMT: false,
			prec4:              false,
		}, want: "2.14", wantErr: false},
		{name: "Test_prettyFloatString 2", args: args{
			num:                "2.1387987031263619",
			percent:            true,
			nearestThousandFMT: false,
			prec4:              false,
		}, want: "2.14%", wantErr: false},
		{name: "Test_prettyFloatString 3", args: args{
			num:                "1137176492787.6108331328990334",
			percent:            false,
			nearestThousandFMT: true,
			prec4:              false,
		}, want: "1.1T", wantErr: false},
		{name: "Test_prettyFloatString 4", args: args{
			num:                "2.1387987031263619",
			percent:            false,
			nearestThousandFMT: false,
			prec4:              true,
		}, want: "2.1388", wantErr: false},
		{name: "Test_prettyFloatString 5 Nearest Thousand", args: args{
			num:                "1137176492.6108331328990334",
			percent:            false,
			nearestThousandFMT: true,
			prec4:              false,
		}, want: "1.1B", wantErr: false},
		{name: "Test_prettyFloatString 6 Nearest Thousand", args: args{
			num:                "1137176.6108331328990334",
			percent:            false,
			nearestThousandFMT: true,
			prec4:              false,
		}, want: "1.1M", wantErr: false},
		{name: "Test_prettyFloatString 6 Nearest Thousand", args: args{
			num:                "1137.6108331328990334",
			percent:            false,
			nearestThousandFMT: true,
			prec4:              false,
		}, want: "1.1K", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := prettyFloatString(tt.args.num, tt.args.percent, tt.args.nearestThousandFMT, tt.args.prec4)
			if (err != nil) != tt.wantErr {
				t.Errorf("prettyFloatString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("prettyFloatString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printTable(t *testing.T) {
	type args struct {
		coinData *CoinData
	}

	rawReader := strings.NewReader(rawData)
	coinData := new(CoinData)
	if err := json.NewDecoder(rawReader).Decode(coinData); err != nil {
		t.Errorf("couldnt decode data: %v", err)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test_printTable 1", args: args{coinData: coinData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := printTable(tt.args.coinData); (err != nil) != tt.wantErr {
				t.Errorf("printTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_printTotalMarketCap(t *testing.T) {
	type args struct {
		coinData *CoinData
	}

	rawReader := strings.NewReader(rawData)
	coinData := new(CoinData)
	if err := json.NewDecoder(rawReader).Decode(coinData); err != nil {
		t.Errorf("couldnt decode data: %v", err)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test_printTotalMarketCap 1", args: args{coinData: coinData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := printTotalMarketCap(tt.args.coinData); (err != nil) != tt.wantErr {
				t.Errorf("printTotalMarketCap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

