package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

// lookup7Day
func lookup7Day(data *CMCCoinData, coinIndex string) (string, error) {
	coinID, err := strconv.Atoi(coinIndex)
	if err != nil {
		return "", err
	}
	for _, row := range data.Data {
		if coinID == row.CmcRank {
			p := message.NewPrinter(language.English)
			percent := p.Sprintf("%.2f", row.Quote.Usd.PercentChange7D)
			return percent, nil
		}
	}

	return "", nil
}

func printTable(coinData *CoinData, cmcData *CMCCoinData) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank", "Name", "Symbol", "Price (USD)", "Change 24H", "Change 7Day", "VWAP 24H", "Market Cap", "Supply", "Volume 24H"})
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

	for _, row := range coinData.Data {
		vwap7Day, err := lookup7Day(cmcData, row.Rank)
		if err != nil {
			return err
		}

		vwap7DayString, _ := prettyFloatString(vwap7Day, true, false, false)
		price, _ := prettyFloatString(row.PriceUSD, false, false, true)
		changePercent24Hr, _ := prettyFloatString(row.ChangePercent24Hr, true, false, false)
		vwap24Hr, _ := prettyFloatString(row.Vwap24Hr, false, false, false)
		marketCapUSD, _ := prettyFloatString(row.MarketCapUSD, false, true, false)
		maxSupply, _ := prettyFloatString(row.MaxSupply, false, true, false)
		volume24Hr, _ := prettyFloatString(row.VolumeUSD24Hr, false, true, false)

		colorData := []string{row.Rank, row.Name, row.Symbol, "$" + price, changePercent24Hr, vwap7DayString, "$" + vwap24Hr, "$" + marketCapUSD, maxSupply, volume24Hr}

		if strings.HasPrefix(row.ChangePercent24Hr, "-") && strings.HasPrefix(vwap7DayString, "-") {
			table.Rich(colorData, []tablewriter.Colors{
				{},
				{tablewriter.FgHiRedColor},
				{},
				{tablewriter.Normal},
				{tablewriter.Bold, tablewriter.FgHiRedColor},
				{tablewriter.Bold, tablewriter.FgHiRedColor},
				{tablewriter.Normal},
				{tablewriter.Normal},
				{tablewriter.Normal},
				{tablewriter.Normal},
			})
		} else if strings.HasPrefix(row.ChangePercent24Hr, "-") && !strings.HasPrefix(vwap7DayString, "-") {
			table.Rich(colorData, []tablewriter.Colors{
				{},
				{tablewriter.FgHiRedColor},
				{},
				{tablewriter.Normal},
				{tablewriter.Normal, tablewriter.FgHiRedColor},
				{tablewriter.Normal, tablewriter.FgHiGreenColor},
				{tablewriter.Normal},
				{tablewriter.Normal},
				{tablewriter.Normal},
				{tablewriter.Normal},
			})
		} else if !strings.HasPrefix(row.ChangePercent24Hr, "-") && strings.HasPrefix(vwap7DayString, "-") {
			table.Rich(colorData, []tablewriter.Colors{
				{},
				{tablewriter.Normal, tablewriter.FgHiGreenColor},
				{},
				{tablewriter.Normal},
				{tablewriter.Normal, tablewriter.FgHiGreenColor},
				{tablewriter.Normal, tablewriter.FgHiRedColor},
				{tablewriter.Normal},
				{tablewriter.Normal},
				{tablewriter.Normal},
				{tablewriter.Normal},
			})
		} else {
			table.Rich(colorData, []tablewriter.Colors{
				{},
				{tablewriter.FgHiGreenColor},
				{},
				{tablewriter.Normal},
				{tablewriter.Normal, tablewriter.FgHiGreenColor},
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
	totalMarketCapFull, _ := prettyFloatString(totalString, false, false, false)
	fmt.Printf("\nTotal Crypto Market Cap \u2248 $%s \n", totalMarketCapFull)
	return nil
}

// printTopMovers prints the largest gaining and losing coin from past 24 hours
func printTopMovers(data *CoinData, cmcData *CMCCoinData) error {
	var highest float64 = 0
	var lowest float64 = 0
	topGainers := &CoinData{}
	topLosers := &CoinData{}

	for _, row := range data.Data {
		if row.ChangePercent24Hr != "" {
			num, err := strconv.ParseFloat(row.ChangePercent24Hr, 64)
			if err != nil {
				return err
			}
			if num > highest {
				highest = num
				topGainers.Data = []Data{row}
			}
			if num < lowest {
				lowest = num
				topLosers.Data = []Data{row}
			}
		}
	}

	fmt.Println("====================================Top Gainers Past 24 Hours========================================")
	if err := printTable(topGainers, cmcData); err != nil {
		return err
	}
	fmt.Println("====================================Top Losers Past 24 Hours=========================================")
	if err := printTable(topLosers, cmcData); err != nil {
		return err
	}

	return nil
}

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
		//fmt.Println(resp.Status)
		//respBody, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(respBody))
	}

	return nil
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		now := time.Now()
		fmt.Printf("%s\n", now.Format("01-02-2006 15:04 PM Monday"))

		argv := ctx.Argv().(*argT)
		var defaultURL string
		if len(argv.Find) >= 1 {
			coins := strings.Join(argv.Find, ",")
			defaultURL = fmt.Sprintf("https://api.coincap.io/v2/assets?ids=%s", coins)
		} else {
			defaultURL = fmt.Sprintf("https://api.coincap.io/v2/assets?limit=%s", argv.Top)
		}

		cmcData := new(CMCCoinData)
		if err := getCoinMarketCapAPI(cmcData, argv); err != nil {
			log.Panic(err)
		}
		coinData := new(CoinData)
		if err := getJSON(defaultURL, coinData); err != nil {
			log.Panic(err)
		}
		if err := printTable(coinData, cmcData); err != nil {
			log.Panic(err)
		}

		// needed to prevent request rate limit
		time.Sleep(time.Second * 2)

		totalMarketCapCoinData := new(CoinData)
		if err := getJSON(top2000, totalMarketCapCoinData); err != nil {
			log.Panic(err)
		}
		//if err := printTopMovers(totalMarketCapCoinData, cmcData); err != nil {
		//	log.Panic(err)
		//}
		if err := printTotalMarketCap(totalMarketCapCoinData); err != nil {
			log.Panic(err)
		}

		return nil
	}))

}
