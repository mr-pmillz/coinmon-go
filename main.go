package main

import (
	"coinmon-go/utils"
	"encoding/json"
	"fmt"
	"github.com/mkideal/cli"
	"github.com/olekukonko/tablewriter"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type argT struct {
	cli.Helper
	Top string `cli:"!t,top" usage:"Amount of coins to show data for, Default is top 10, If -f|--find flag supplied, -t|--top is ignored" dft:"10"`
	Find []string `cli:"!f,find" usage:"Specific Coins to return. Example: bitcoin,cardano,ethereum,uniswap"`
}

type CoinData struct {
	Data      []Data `json:"data"`
	Timestamp int64  `json:"timestamp"`
}
type Data struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply,omitempty"`
	MarketCapUSD      string `json:"marketCapUsd"`
	VolumeUSD24Hr     string `json:"volumeUsd24Hr"`
	PriceUSD          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
	Explorer          string `json:"explorer"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func prettyFloatString(num string, percent, nearestThousandFMT, prec4 bool) (string, error) {
	number, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return "", err
	}
	if percent {
		return fmt.Sprintf("%.2f%%", number), nil
	} else if nearestThousandFMT {
		return utils.NearestThousandFormat(number), nil
	} else if prec4 {
		return fmt.Sprintf("%.4f", number), nil
	}
	return fmt.Sprintf("%.2f", number), nil
}

func printTable(coinData *CoinData) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank", "Name", "Symbol", "Price (USD)", "Change 24H", "VWAP 24H", "Market Cap", "Supply", "Volume 24H"})
	table.SetBorder(false)

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
	)

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.FgHiWhiteColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.FgHiWhiteColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor, tablewriter.BgBlackColor},
	)

	for _, row := range coinData.Data {
		price, _ := prettyFloatString(row.PriceUSD, false, false, true)
		changePercent24Hr, _ := prettyFloatString(row.ChangePercent24Hr, true, false, false)
		vwap24Hr, _ := prettyFloatString(row.Vwap24Hr, false, false, false)
		marketCapUSD, _ := prettyFloatString(row.MarketCapUSD, false, true, false)
		maxSupply, _ := prettyFloatString(row.MaxSupply, false, true, false)
		volume24Hr, _ := prettyFloatString(row.VolumeUSD24Hr, false, true, false)

		colorData := []string{row.Rank, row.Name, row.Symbol, "$" + price, changePercent24Hr, "$" + vwap24Hr, "$" + marketCapUSD, maxSupply, volume24Hr}

		if strings.HasPrefix(row.ChangePercent24Hr, "-") {
			table.Rich(colorData, []tablewriter.Colors{
				{},
				{tablewriter.FgHiRedColor},
				{},
				{tablewriter.Normal},
				{tablewriter.Bold, tablewriter.FgHiRedColor},
				{tablewriter.Normal},
				{tablewriter.Normal},
				{tablewriter.Normal},
				{tablewriter.Normal},
			})
		} else {
			table.Rich(colorData, []tablewriter.Colors{
				{},
				{},
				{},
				{tablewriter.Normal},
				{tablewriter.Normal, tablewriter.FgHiGreenColor},
				{tablewriter.Normal},
				{tablewriter.Normal},
				{tablewriter.Normal},
				{tablewriter.Normal},
			})
		}
	}

	table.SetAutoMergeCells(true)
	table.Render()

	return nil
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		var url string
		if len(argv.Find) >= 1 {
			coins := strings.Join(argv.Find, ",")
			url = fmt.Sprintf("https://api.coincap.io/v2/assets?ids=%s", coins)
		} else {
			url = fmt.Sprintf("https://api.coincap.io/v2/assets?limit=%s", argv.Top)
		}

		coinData := new(CoinData)
		if err := getJson(url, coinData); err != nil {
			panic(err)
		}
		if err := printTable(coinData); err != nil {
			panic(err)
		}

		return nil
	}))

}
