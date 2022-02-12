package twitch

import (
	"001.AI/config"
	"001.AI/logger"
	"fmt"
	"github.com/nicklaw5/helix/v2"
	"time"
)

var tw *helix.Client

func init() {

}

func CreateTwitchSession() {
	var err error
	tw, err = helix.NewClient(&helix.Options{
		ClientID: config.GetTwitchToken(),
	})
	if err != nil {
		logger.PrintLog("cant set new client %s\n", err.Error())
		return
	}
	logger.PrintLog("twitch initialized: %s\n", tw.GetAppAccessToken())
	go checkChannels()
}

func checkChannels() {
	for {
		time.Sleep(3000 * time.Millisecond)
		streams, err := tw.GetStreams(&helix.StreamsParams{
			UserLogins: []string{
				"zakvielchannel",
				"artem_wolff",
			},
		})
		if err != nil {
			logger.PrintLog("twitch error: get streams: %s\n", err.Error())
		} else {
			fmt.Println(streams.Data.Streams)
		}

	}

}
