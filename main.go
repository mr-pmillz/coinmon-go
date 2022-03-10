package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mr-pmillz/coinmon-go/utils"

	"github.com/mkideal/cli"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const top2000 = "https://api.coincap.io/v2/assets?limit=2000"

type argT struct {
	cli.Helper
	Top  string   `cli:"!t,top" usage:"Amount of coins to show data for, Default is top 10, If -f|--find flag supplied, -t|--top is ignored" dft:"10"`
	Find []string `cli:"!f,find" usage:"Specific Coins to return. Example: bitcoin,cardano,ethereum,uniswap"`
}

// CoinData ...
type CoinData struct {
	Data      []Data `json:"data"`
	Timestamp int64  `json:"timestamp"`
}

// Data ...
type Data struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply,omitempty"`
	MaxSupply         string `json:"maxSupply,omitempty"`
	MarketCapUSD      string `json:"marketCapUsd,omitempty"`
	VolumeUSD24Hr     string `json:"volumeUsd24Hr,omitempty"`
	PriceUSD          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
	Explorer          string `json:"explorer"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode == 429 {
		fmt.Printf("[-] Sleeping for 10 seconds, too many requests HTTP Response: %d\n", r.StatusCode)
		time.Sleep(time.Second * 10)
		req, err := myClient.Get(url)
		if err != nil {
			return err
		}
		defer req.Body.Close()
		return json.NewDecoder(req.Body).Decode(target)
	}

	return json.NewDecoder(r.Body).Decode(target)
}

func prettyFloatString(num string, percent, nearestThousandFMT, prec4 bool) (string, error) {
	number, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return "", err
	}
	p := message.NewPrinter(language.English)
	if percent {
		return p.Sprintf("%.2f%%", number), nil
	} else if nearestThousandFMT {
		return utils.NearestThousandFormat(number), nil
	} else if prec4 {
		return p.Sprintf("%.4f", number), nil
	}
	return p.Sprintf("%.2f", number), nil
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

func printTotalMarketCap(coinData *CoinData) error {
	var total float64 = 0
	for _, row := range coinData.Data {
		if row.MarketCapUSD != "" {
			number, err := strconv.ParseFloat(row.MarketCapUSD, 64)
			if err != nil {
				return err
			}
			total += number
		}
	}

	totalString := strconv.FormatFloat(total, 'E', -1, 64)
	totalMarketCap, _ := prettyFloatString(totalString, false, true, false)
	totalMarketCapFull, _ := prettyFloatString(totalString, false, false, false)
	fmt.Printf("\nTotal Crypto Market Cap \u2248 $%s \u2248 $%s \n", totalMarketCap, totalMarketCapFull)
	return nil
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		now := time.Now()
		fmt.Printf("%s\n", now.Format("01-02-2006 15:04 PM Monday"))

		argv := ctx.Argv().(*argT)
		var url string
		if len(argv.Find) >= 1 {
			coins := strings.Join(argv.Find, ",")
			url = fmt.Sprintf("https://api.coincap.io/v2/assets?ids=%s", coins)
		} else {
			url = fmt.Sprintf("https://api.coincap.io/v2/assets?limit=%s", argv.Top)
		}

		coinData := new(CoinData)
		if err := getJSON(url, coinData); err != nil {
			log.Panic(err)
		}
		if err := printTable(coinData); err != nil {
			log.Panic(err)
		}

		// needed to prevent request rate limit
		time.Sleep(time.Second * 2)

		totalMarketCapCoinData := new(CoinData)
		if err := getJSON(top2000, totalMarketCapCoinData); err != nil {
			log.Panic(err)
		}
		if err := printTotalMarketCap(totalMarketCapCoinData); err != nil {
			log.Panic(err)
		}

		return nil
	}))

}
