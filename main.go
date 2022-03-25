package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/mr-pmillz/coinmon-go/utils"

	"github.com/mkideal/cli"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type CMCCoinData struct {
	Status Status    `json:"status"`
	Data   []CMCData `json:"data"`
}
type Status struct {
	Timestamp    time.Time   `json:"timestamp"`
	ErrorCode    int         `json:"error_code"`
	ErrorMessage interface{} `json:"error_message"`
	Elapsed      int         `json:"elapsed"`
	CreditCount  int         `json:"credit_count"`
	Notice       interface{} `json:"notice"`
	TotalCount   int         `json:"total_count"`
}
type Usd struct {
	Price                 float64   `json:"price"`
	Volume24H             float64   `json:"volume_24h"`
	VolumeChange24H       float64   `json:"volume_change_24h"`
	PercentChange1H       float64   `json:"percent_change_1h"`
	PercentChange24H      float64   `json:"percent_change_24h"`
	PercentChange7D       float64   `json:"percent_change_7d"`
	PercentChange30D      float64   `json:"percent_change_30d"`
	PercentChange60D      float64   `json:"percent_change_60d"`
	PercentChange90D      float64   `json:"percent_change_90d"`
	MarketCap             float64   `json:"market_cap"`
	MarketCapDominance    float64   `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	LastUpdated           time.Time `json:"last_updated"`
}
type Quote struct {
	Usd Usd `json:"USD"`
}
type CMCData struct {
	ID                            int         `json:"id"`
	Name                          string      `json:"name"`
	Symbol                        string      `json:"symbol"`
	Slug                          string      `json:"slug"`
	NumMarketPairs                int         `json:"num_market_pairs"`
	DateAdded                     time.Time   `json:"date_added"`
	Tags                          []string    `json:"tags"`
	MaxSupply                     float64     `json:"max_supply"`
	CirculatingSupply             float64     `json:"circulating_supply"`
	TotalSupply                   float64     `json:"total_supply"`
	Platform                      interface{} `json:"platform"`
	CmcRank                       int         `json:"cmc_rank"`
	SelfReportedCirculatingSupply interface{} `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         interface{} `json:"self_reported_market_cap"`
	LastUpdated                   time.Time   `json:"last_updated"`
	Quote                         Quote       `json:"quote"`
}

type argT struct {
	cli.Helper
	Top string `cli:"!t,top" usage:"Amount of coins to show data for, Default is top 10, If -f|--find flag supplied, -t|--top is ignored" dft:"10"`
}

func formatFloat(number float64, percent, nearestThousandFMT, prec4 bool) (string, error) {
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

func redOrGreen(num float64) tablewriter.Colors {
	if num < 0 {
		return tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor}
	}
	return tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor}
}

func printTable(cmcData *CMCCoinData) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank", "Name", "Symbol", "Price (USD)", "Change 24H", "Change 7Day", "Change 30Day", "Market Cap", "Supply", "Volume 24H"})
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
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor, tablewriter.BgBlackColor},
	)

	for _, row := range cmcData.Data {
		vwap7DayString, _ := formatFloat(row.Quote.Usd.PercentChange7D, true, false, false)
		price, _ := formatFloat(row.Quote.Usd.Price, false, false, true)
		changePercent24Hr, _ := formatFloat(row.Quote.Usd.PercentChange24H, true, false, false)
		changePercent30Day, _ := formatFloat(row.Quote.Usd.PercentChange30D, true, false, false)
		marketCapUSD, _ := formatFloat(row.Quote.Usd.MarketCap, false, true, false)
		maxSupply, _ := formatFloat(row.MaxSupply, false, true, false)
		volume24Hr, _ := formatFloat(row.Quote.Usd.Volume24H, false, true, false)
		rank := strconv.FormatInt(int64(row.CmcRank), 10)

		colorData := []string{rank, row.Name, row.Symbol, "$" + price, changePercent24Hr, vwap7DayString, changePercent30Day, "$" + marketCapUSD, maxSupply, volume24Hr}

		table.Rich(colorData, []tablewriter.Colors{
			{},
			redOrGreen(row.Quote.Usd.PercentChange24H),
			{},
			{tablewriter.Normal},
			redOrGreen(row.Quote.Usd.PercentChange24H),
			redOrGreen(row.Quote.Usd.PercentChange7D),
			redOrGreen(row.Quote.Usd.PercentChange30D),
			{tablewriter.Normal},
			{tablewriter.Normal},
			{tablewriter.Normal},
		})
	}

	table.SetAutoMergeCells(true)
	table.Render()

	return nil
}

func printTotalMarketCap(coinData *GlobalMetrics) error {
	totalMarketCapFull, _ := formatFloat(coinData.Data.Quote.Usd.TotalMarketCap, false, false, false)
	fmt.Printf("\nTotal Crypto Market Cap \u2248 $%s \n", totalMarketCapFull)
	return nil
}

// getCoinMarketCapAPI ...
func getCoinMarketCapAPI(target interface{}, argv *argT) error {
	apiKey, _ := os.LookupEnv("COINMARKETCAP_API_KEY")
	if apiKey != "" {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}

		q := url.Values{}
		q.Add("start", "1")
		q.Add("limit", argv.Top)
		q.Add("convert", "USD")

		req.Header.Set("Accepts", "application/json")
		req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
		req.URL.RawQuery = q.Encode()

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request to server")
			os.Exit(1)
		}
		defer resp.Body.Close()

		return json.NewDecoder(resp.Body).Decode(target)
	}

	return nil
}

type GlobalMetrics struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       string      `json:"notice"`
	} `json:"status"`
	Data struct {
		Quote struct {
			Usd struct {
				TotalMarketCap float64 `json:"total_market_cap"`
			} `json:"USD"`
		} `json:"quote"`
		LastUpdated time.Time `json:"last_updated"`
	} `json:"data"`
}

// getTotalMarketCap ...
func getTotalMarketCap(target interface{}) error {
	apiKey, _ := os.LookupEnv("COINMARKETCAP_API_KEY")
	if apiKey != "" {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/global-metrics/quotes/latest", nil)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}

		req.Header.Set("Accepts", "application/json")
		req.Header.Add("X-CMC_PRO_API_KEY", apiKey)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request to server")
			os.Exit(1)
		}
		defer resp.Body.Close()

		return json.NewDecoder(resp.Body).Decode(target)
	}

	return nil
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		now := time.Now()
		fmt.Printf("%s\n", now.Format("01-02-2006 15:04 PM Monday"))

		argv := ctx.Argv().(*argT)
		cmcData := new(CMCCoinData)
		if err := getCoinMarketCapAPI(cmcData, argv); err != nil {
			log.Panic(err)
		}
		if err := printTable(cmcData); err != nil {
			log.Panic(err)
		}

		totalMarketCapCoinData := new(GlobalMetrics)
		if err := getTotalMarketCap(totalMarketCapCoinData); err != nil {
			log.Panic(err)
		}

		if err := printTotalMarketCap(totalMarketCapCoinData); err != nil {
			log.Panic(err)
		}

		return nil
	}))

}
