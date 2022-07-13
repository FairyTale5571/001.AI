package bot

import (
	"github.com/bwmarrin/discordgo"
)

var voiceChannels = map[string]*discordgo.VoiceConnection{}

const (
	pathToRecords = "stream_records"
)

type ExchangeInfo struct {
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	CloseTime          int    `json:"closeTime"`
	Count              int    `json:"count"`
	FirstID            int    `json:"firstId"`
	HighPrice          string `json:"highPrice"`
	LastID             int    `json:"lastId"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	LowPrice           string `json:"lowPrice"`
	OpenPrice          string `json:"openPrice"`
	OpenTime           int    `json:"openTime"`
	PrevClosePrice     string `json:"prevClosePrice"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	QuoteVolume        string `json:"quoteVolume"`
	Symbol             string `json:"symbol"`
	Volume             string `json:"volume"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
}

type ExchangeInfos []struct {
	ExchangeInfo
}

type YFI struct {
	OptionChain struct {
		Error  interface{} `json:"error"`
		Result []struct {
			Quote struct {
				Ask                               float64 `json:"ask"`
				AskSize                           int     `json:"askSize"`
				AverageDailyVolume10Day           int     `json:"averageDailyVolume10Day"`
				AverageDailyVolume3Month          int     `json:"averageDailyVolume3Month"`
				Bid                               float64 `json:"bid"`
				BidSize                           int     `json:"bidSize"`
				Currency                          string  `json:"currency"`
				CustomPriceAlertConfidence        string  `json:"customPriceAlertConfidence"`
				EsgPopulated                      bool    `json:"esgPopulated"`
				Exchange                          string  `json:"exchange"`
				ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
				ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
				ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
				FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
				FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
				FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
				FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
				FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
				FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
				FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
				FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
				FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
				FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
				FirstTradeDateMilliseconds        int     `json:"firstTradeDateMilliseconds"`
				FullExchangeName                  string  `json:"fullExchangeName"`
				GmtOffSetMilliseconds             int     `json:"gmtOffSetMilliseconds"`
				Language                          string  `json:"language"`
				Market                            string  `json:"market"`
				MarketState                       string  `json:"marketState"`
				MessageBoardID                    string  `json:"messageBoardId"`
				PriceHint                         int     `json:"priceHint"`
				QuoteSourceName                   string  `json:"quoteSourceName"`
				QuoteType                         string  `json:"quoteType"`
				Region                            string  `json:"region"`
				RegularMarketChange               float64 `json:"regularMarketChange"`
				RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
				RegularMarketDayHigh              float64 `json:"regularMarketDayHigh"`
				RegularMarketDayLow               float64 `json:"regularMarketDayLow"`
				RegularMarketDayRange             string  `json:"regularMarketDayRange"`
				RegularMarketOpen                 float64 `json:"regularMarketOpen"`
				RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
				RegularMarketPrice                float64 `json:"regularMarketPrice"`
				RegularMarketTime                 int     `json:"regularMarketTime"`
				TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
				TypeDisp                          string  `json:"typeDisp"`
				ShortName                         string  `json:"shortName"`
			} `json:"quote"`
		} `json:"result"`
	} `json:"optionChain"`
}
