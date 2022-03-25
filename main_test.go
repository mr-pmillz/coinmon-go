package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/mkideal/cli"
	"github.com/olekukonko/tablewriter"
)

const coinJSON = `{"status":{"timestamp":"2022-03-25T19:43:42.566Z","error_code":0,"error_message":null,"elapsed":21,"credit_count":1,"notice":null,"total_count":9722},"data":[{"id":1,"name":"Bitcoin","symbol":"BTC","slug":"bitcoin","num_market_pairs":9274,"date_added":"2013-04-28T00:00:00.000Z","tags":["mineable","pow","sha-256","store-of-value","state-channel","coinbase-ventures-portfolio","three-arrows-capital-portfolio","polychain-capital-portfolio","binance-labs-portfolio","blockchain-capital-portfolio","boostvc-portfolio","cms-holdings-portfolio","dcg-portfolio","dragonfly-capital-portfolio","electric-capital-portfolio","fabric-ventures-portfolio","framework-ventures-portfolio","galaxy-digital-portfolio","huobi-capital-portfolio","alameda-research-portfolio","a16z-portfolio","1confirmation-portfolio","winklevoss-capital-portfolio","usv-portfolio","placeholder-ventures-portfolio","pantera-capital-portfolio","multicoin-capital-portfolio","paradigm-portfolio"],"max_supply":21000000,"circulating_supply":18993625,"total_supply":18993625,"platform":null,"cmc_rank":1,"self_reported_circulating_supply":null,"self_reported_market_cap":null,"last_updated":"2022-03-25T19:42:00.000Z","quote":{"USD":{"price":44273.52652845786,"volume_24h":30324483516.52537,"volume_change_24h":-2.5043,"percent_change_1h":-0.25340873,"percent_change_24h":0.58373623,"percent_change_7d":6.06744851,"percent_change_30d":17.39290612,"percent_change_60d":23.73294567,"percent_change_90d":-13.03873223,"market_cap":840914760309.0804,"market_cap_dominance":42.3248,"fully_diluted_market_cap":929744057097.62,"last_updated":"2022-03-25T19:42:00.000Z"}}},{"id":1027,"name":"Ethereum","symbol":"ETH","slug":"ethereum","num_market_pairs":5586,"date_added":"2015-08-07T00:00:00.000Z","tags":["mineable","pow","smart-contracts","ethereum-ecosystem","coinbase-ventures-portfolio","three-arrows-capital-portfolio","polychain-capital-portfolio","binance-labs-portfolio","blockchain-capital-portfolio","boostvc-portfolio","cms-holdings-portfolio","dcg-portfolio","dragonfly-capital-portfolio","electric-capital-portfolio","fabric-ventures-portfolio","framework-ventures-portfolio","hashkey-capital-portfolio","kenetic-capital-portfolio","huobi-capital-portfolio","alameda-research-portfolio","a16z-portfolio","1confirmation-portfolio","winklevoss-capital-portfolio","usv-portfolio","placeholder-ventures-portfolio","pantera-capital-portfolio","multicoin-capital-portfolio","paradigm-portfolio","injective-ecosystem","bnb-chain"],"max_supply":null,"circulating_supply":120108316.3115,"total_supply":120108316.3115,"platform":null,"cmc_rank":2,"self_reported_circulating_supply":null,"self_reported_market_cap":null,"last_updated":"2022-03-25T19:42:00.000Z","quote":{"USD":{"price":3096.36226228143,"volume_24h":17087268447.331795,"volume_change_24h":-7.8157,"percent_change_1h":-0.79385952,"percent_change_24h":-0.61268378,"percent_change_7d":5.43129681,"percent_change_30d":17.11251285,"percent_change_60d":31.57620683,"percent_change_90d":-24.11254865,"market_cap":371898858013.0897,"market_cap_dominance":18.7191,"fully_diluted_market_cap":371898858013.09,"last_updated":"2022-03-25T19:42:00.000Z"}}},{"id":825,"name":"Tether","symbol":"USDT","slug":"tether","num_market_pairs":30282,"date_added":"2015-02-25T00:00:00.000Z","tags":["payments","stablecoin","asset-backed-stablecoin","avalanche-ecosystem","solana-ecosystem","arbitrum-ecosytem","moonriver-ecosystem","injective-ecosystem","bnb-chain"],"max_supply":null,"circulating_supply":80958735861.26395,"total_supply":84405585147.17514,"platform":{"id":1027,"name":"Ethereum","symbol":"ETH","slug":"ethereum","token_address":"0xdac17f958d2ee523a2206206994597c13d831ec7"},"cmc_rank":3,"self_reported_circulating_supply":null,"self_reported_market_cap":null,"last_updated":"2022-03-25T19:42:00.000Z","quote":{"USD":{"price":1.0004115084826206,"volume_24h":65178909389.26961,"volume_change_24h":-16.1267,"percent_change_1h":-0.00230759,"percent_change_24h":-0.00903119,"percent_change_7d":-0.01885041,"percent_change_30d":-0.0077849,"percent_change_60d":0.00032218,"percent_change_90d":-0.03856683,"market_cap":80992051067.8131,"market_cap_dominance":4.0766,"fully_diluted_market_cap":84440318761.44,"last_updated":"2022-03-25T19:42:00.000Z"}}},{"id":1839,"name":"BNB","symbol":"BNB","slug":"bnb","num_market_pairs":769,"date_added":"2017-07-25T00:00:00.000Z","tags":["marketplace","centralized-exchange","payments","smart-contracts","alameda-research-portfolio","multicoin-capital-portfolio","moonriver-ecosystem","bnb-chain"],"max_supply":165116760,"circulating_supply":165116760.89,"total_supply":165116760.89,"platform":null,"cmc_rank":4,"self_reported_circulating_supply":null,"self_reported_market_cap":null,"last_updated":"2022-03-25T19:42:00.000Z","quote":{"USD":{"price":409.5078411286467,"volume_24h":1710069957.3542104,"volume_change_24h":-5.9208,"percent_change_1h":-0.31505865,"percent_change_24h":-1.33497668,"percent_change_7d":3.32239253,"percent_change_30d":10.11865877,"percent_change_60d":13.26950868,"percent_change_90d":-25.07237702,"market_cap":67616608286.21886,"market_cap_dominance":3.4033,"fully_diluted_market_cap":67616607921.76,"last_updated":"2022-03-25T19:42:00.000Z"}}},{"id":3408,"name":"USD Coin","symbol":"USDC","slug":"usd-coin","num_market_pairs":3394,"date_added":"2018-10-08T00:00:00.000Z","tags":["medium-of-exchange","stablecoin","asset-backed-stablecoin","fantom-ecosystem","arbitrum-ecosytem","moonriver-ecosystem","bnb-chain"],"max_supply":null,"circulating_supply":52450514469.16298,"total_supply":52450514469.16298,"platform":{"id":1027,"name":"Ethereum","symbol":"ETH","slug":"ethereum","token_address":"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"},"cmc_rank":5,"self_reported_circulating_supply":null,"self_reported_market_cap":null,"last_updated":"2022-03-25T19:42:00.000Z","quote":{"USD":{"price":0.9997558120438371,"volume_24h":4011091296.3792906,"volume_change_24h":-14.9026,"percent_change_1h":-0.01556295,"percent_change_24h":0.01051961,"percent_change_7d":0.00022441,"percent_change_30d":0.03336693,"percent_change_60d":-0.07129258,"percent_change_90d":-0.01023823,"market_cap":52437706685.23506,"market_cap_dominance":2.6393,"fully_diluted_market_cap":52437706685.24,"last_updated":"2022-03-25T19:42:00.000Z"}}},{"id":52,"name":"XRP","symbol":"XRP","slug":"xrp","num_market_pairs":695,"date_added":"2013-08-04T00:00:00.000Z","tags":["medium-of-exchange","enterprise-solutions","binance-chain","arrington-xrp-capital-portfolio","galaxy-digital-portfolio","a16z-portfolio","pantera-capital-portfolio"],"max_supply":100000000000,"circulating_supply":48121609012,"total_supply":99989656524,"platform":null,"cmc_rank":6,"self_reported_circulating_supply":null,"self_reported_market_cap":null,"last_updated":"2022-03-25T19:42:00.000Z","quote":{"USD":{"price":0.8238937816425282,"volume_24h":1650424450.8464234,"volume_change_24h":-26.8746,"percent_change_1h":-0.27547789,"percent_change_24h":-2.83283576,"percent_change_7d":3.94539805,"percent_change_30d":15.97268325,"percent_change_60d":39.09155654,"percent_change_90d":-10.93731819,"market_cap":39647094427.61985,"market_cap_dominance":1.9955,"fully_diluted_market_cap":82389378164.25,"last_updated":"2022-03-25T19:42:00.000Z"}}},{"id":2010,"name":"Cardano","symbol":"ADA","slug":"cardano","num_market_pairs":409,"date_added":"2017-10-01T00:00:00.000Z","tags":["mineable","dpos","pos","platform","research","smart-contracts","staking","cardano-ecosystem","cardano","bnb-chain"],"max_supply":45000000000,"circulating_supply":33726286575.782,"total_supply":34239675266.895,"platform":null,"cmc_rank":7,"self_reported_circulating_supply":null,"self_reported_market_cap":null,"last_updated":"2022-03-25T19:42:00.000Z","quote":{"USD":{"price":1.090102511168468,"volume_24h":2447842580.097247,"volume_change_24h":-33.7853,"percent_change_1h":-0.33370579,"percent_change_24h":-7.30916416,"percent_change_7d":27.85431874,"percent_change_30d":20.20496145,"percent_change_60d":3.33435961,"percent_change_90d":-24.1825107,"market_cap":36765109688.64735,"market_cap_dominance":1.8505,"fully_diluted_market_cap":49054613002.58,"last_updated":"2022-03-25T19:42:00.000Z"}}}]}`
const totalMarketCapJSON = `{"status":{"timestamp":"2022-03-25T22:48:24.172Z","error_code":0,"error_message":null,"elapsed":18,"credit_count":1,"notice":"You have used 163% of your plan's daily credit limit."},"data":{"active_cryptocurrencies":9730,"total_cryptocurrencies":18470,"active_market_pairs":56484,"active_exchanges":475,"total_exchanges":1611,"eth_dominance":18.741195405813,"btc_dominance":42.302311578511,"eth_dominance_yesterday":18.70945352,"btc_dominance_yesterday":41.81609104,"eth_dominance_24h_percentage_change":0.031741885813,"btc_dominance_24h_percentage_change":0.486220538511,"defi_volume_24h":11382725925.543377,"defi_volume_24h_reported":29354560140.08033,"defi_market_cap":142761093112.20792,"defi_24h_percentage_change":-17.879482370765,"stablecoin_volume_24h":75033093629.8481,"stablecoin_volume_24h_reported":210099128623.3004,"stablecoin_market_cap":180920888992.7633,"stablecoin_24h_percentage_change":-18.103599407868,"derivatives_volume_24h":162436548824.35928,"derivatives_volume_24h_reported":162436548824.35928,"derivatives_24h_percentage_change":-2.860356587268,"quote":{"USD":{"total_market_cap":1991441743015.1462,"total_volume_24h":91808749490.2,"total_volume_24h_reported":246160677760.9,"altcoin_volume_24h":61246128626.896126,"altcoin_volume_24h_reported":143359982983.54648,"altcoin_market_cap":1149015851980.3398,"defi_volume_24h":11382725925.543377,"defi_volume_24h_reported":29354560140.08033,"defi_24h_percentage_change":-17.879482370765,"defi_market_cap":142761093112.20792,"stablecoin_volume_24h":75033093629.8481,"stablecoin_volume_24h_reported":210099128623.3004,"stablecoin_24h_percentage_change":-18.103599407868,"stablecoin_market_cap":180920888992.7633,"derivatives_volume_24h":162436548824.35928,"derivatives_volume_24h_reported":162436548824.35928,"derivatives_24h_percentage_change":-2.860356587268,"last_updated":"2022-03-25T22:46:59.999Z","total_market_cap_yesterday":1999409870501.2666,"total_volume_24h_yesterday":108705693415.07,"total_market_cap_yesterday_percentage_change":-0.39852396467978224,"total_volume_24h_yesterday_percentage_change":-15.543752488062012}},"last_updated":"2022-03-25T22:46:59.999Z"}}`

func Test_formatFloat(t *testing.T) {
	type args struct {
		num                float64
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
		{name: "Test_formatFloat 1", args: args{
			num:                2.1387987031263619,
			percent:            false,
			nearestThousandFMT: false,
			prec4:              false,
		}, want: "2.14", wantErr: false},
		{name: "Test_formatFloat 2", args: args{
			num:                2.1387987031263619,
			percent:            true,
			nearestThousandFMT: false,
			prec4:              false,
		}, want: "2.14%", wantErr: false},
		{name: "Test_formatFloat 3", args: args{
			num:                1137176492787.6108331328990334,
			percent:            false,
			nearestThousandFMT: true,
			prec4:              false,
		}, want: "1.1T", wantErr: false},
		{name: "Test_formatFloat 4", args: args{
			num:                2.1387987031263619,
			percent:            false,
			nearestThousandFMT: false,
			prec4:              true,
		}, want: "2.1388", wantErr: false},
		{name: "Test_formatFloat 5 Nearest Thousand", args: args{
			num:                1137176492.6108331328990334,
			percent:            false,
			nearestThousandFMT: true,
			prec4:              false,
		}, want: "1.1B", wantErr: false},
		{name: "Test_formatFloat 6 Nearest Thousand", args: args{
			num:                1137176.6108331328990334,
			percent:            false,
			nearestThousandFMT: true,
			prec4:              false,
		}, want: "1.1M", wantErr: false},
		{name: "Test_formatFloat 6 Nearest Thousand", args: args{
			num:                1137.6108331328990334,
			percent:            false,
			nearestThousandFMT: true,
			prec4:              false,
		}, want: "1.1K", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatFloat(tt.args.num, tt.args.percent, tt.args.nearestThousandFMT, tt.args.prec4)
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

func Test_redOrGreen(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		args args
		want tablewriter.Colors
	}{
		{name: "Red", args: args{num: -10.33}, want: tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor}},
		{name: "Green", args: args{num: 25.43}, want: tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redOrGreen(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("redOrGreen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printTable(t *testing.T) {
	type args struct {
		cmcData *CMCCoinData
	}

	rawReader := strings.NewReader(coinJSON)
	coinData := new(CMCCoinData)
	if err := json.NewDecoder(rawReader).Decode(coinData); err != nil {
		t.Errorf("couldnt decode data: %v", err)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "PrintTable", args: args{cmcData: coinData}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := printTable(tt.args.cmcData); (err != nil) != tt.wantErr {
				t.Errorf("printTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_printTotalMarketCap(t *testing.T) {
	type args struct {
		coinData *GlobalMetrics
	}
	rawReader := strings.NewReader(totalMarketCapJSON)
	coinData := new(GlobalMetrics)
	if err := json.NewDecoder(rawReader).Decode(coinData); err != nil {
		t.Errorf("couldnt decode data: %v", err)
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "TotalMarketCap Test", args: args{coinData: coinData}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := printTotalMarketCap(tt.args.coinData); (err != nil) != tt.wantErr {
				t.Errorf("printTotalMarketCap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getCoinMarketCapAPI(t *testing.T) {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		type args struct {
			target interface{}
			argv   *argT
		}
		argv := ctx.Argv().(*argT)
		cmcData := new(CMCCoinData)
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{name: "getCoinMarketCapAPI", args: args{
				target: cmcData,
				argv:   argv,
			}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := getCoinMarketCapAPI(tt.args.target, tt.args.argv); (err != nil) != tt.wantErr {
					t.Errorf("getCoinMarketCapAPI() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
		return nil
	})
}

func Test_getTotalMarketCap(t *testing.T) {
	type args struct {
		target interface{}
	}
	rawReader := strings.NewReader(totalMarketCapJSON)
	coinData := new(GlobalMetrics)
	if err := json.NewDecoder(rawReader).Decode(coinData); err != nil {
		t.Errorf("couldnt decode data: %v", err)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "getTotalMarketCap", args: args{target: coinData}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getTotalMarketCap(tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("getTotalMarketCap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
