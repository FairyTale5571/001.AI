package bot

import (
	"fmt"
	"time"
)

func startRoutine() {
	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				refreshChannels()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func refreshChannels() {
	fmt.Printf("refresh channels\n")
	//go refreshBTC()
	go refreshSP()
}

const (
	ChannelSP500 = "961293558114041866"
	ChannelBTC   = "961353697571524639"
)

func refreshSP() {
	res := sp500().OptionChain.Result[0].Quote
	_, err := s.ChannelEdit(ChannelSP500, fmt.Sprintf("%s: %.2f", res.ShortName, res.RegularMarketPrice))
	if err != nil {
		fmt.Printf("%s | SP500\n", err.Error())
	}
}

func refreshBTC() {
	res := trimZeros(ticker("BTCUSDT").LastPrice)
	_, err := s.ChannelEdit(ChannelBTC, fmt.Sprintf("BTC: %s", res))
	if err != nil {
		fmt.Printf("%s | BTC\n", err.Error())
	}
}
