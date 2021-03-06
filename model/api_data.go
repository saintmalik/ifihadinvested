package model
import "time"
var (
	Id                string
	Vscoinid          string
	Fromdate          string
	Fund              string
	Historical_data   string
	Current_price     string
	Check_response    string
	Crypto_coin       string
	Expected_amount   string
	Check_profit_loss string
)
const (
	LayoutISO = "02-01-2006"
	LayoutUS  = "January 2, 2006"
)
type CoinHistoricalData struct {
	ID         string     `json:"id"`
	Symbol     string     `json:"symbol"`
	Name       string     `json:"name"`
	MarketData MarketData `json:"market_data"`
} //struct to store data
type CoinCurrentData struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Symbol        string     `json:"symbol"`
	MarketCapRank int        `json:"market_cap_rank"`
	MarketData    MarketData `json:"market_data"`
} //struct for the json data
type MarketData struct {
	CurrentPrice struct {
		Usd  float64 `json:"usd"`
		Sats float64 `json:"sats"`
	} `json:"current_price"`
	Ath struct {
		Usd float64 `json:"usd"`
	} `json:"ath"`
	AthDate struct {
		Usd time.Time `json:"usd"`
	} `json:"ath_date"`
	// Atl struct {
	// 	Usd float64 `json:"usd"`
	// } `json:"atl"`
	// AtlDate struct {
	// 	Usd time.Time `json:"usd"`
	// } `json:"atl_date"`
}
type CoinDatas struct {
	CoinName             string
	CoinSymbol           string
	LatestCoinPrice      float64
	TotalCoinOwned       float64
	TotalValueNow        float64
	PurchasePrice        float64
	ChangePercent        float64
	ProfitLoss           float64
	InitialCapital       float64
	InvestmentDate       string
	CheckResponse        string
	CheckValueProfitLoss string
}
type CompareDatas struct {
	CoinName        string
	CoinSymbol      string
	LatestCoinPrice float64
	AllTimeHigh float64
	AllTimeHighDate time.Time
	CoinRank int
	VsCoinName       string
	VsCoinSymbol      string
	VsLatestCoinPrice float64
	VsAllTimeHigh float64
	VsAllTimeHighDate time.Time
	VsCoinRank int
}
type IfHodlDatas struct {
	IfHodl                      float64
	IfYouSold                   float64
	ProfitLoss                  float64
	CoinSymbol                  string
	CoinValueOnSellDate         float64
	CoinOwned                   float64
	AmountYouFeelWorthInSomeDay float64
	SellingDate                 string
	CheckResponse               string
	CoinName                    string
}
type CoinTrendingData struct {
	Coins []struct {
		Item struct {
			ID            string  `json:"id"`
			CoinID        int     `json:"coin_id"`
			Name          string  `json:"name"`
			Symbol        string  `json:"symbol"`
			MarketCapRank int     `json:"market_cap_rank"`
			Thumb         string  `json:"thumb"`
			Small         string  `json:"small"`
			Large         string  `json:"large"`
			Slug          string  `json:"slug"`
			PriceBtc      float64 `json:"price_btc"`
			Score         int     `json:"score"`
		} `json:"item"`
	} `json:"coins"`
}
