package main

import (
	"001.AI/bot"
	"001.AI/config"
	"fmt"
	"github.com/getsentry/sentry-go"
	"log"
	"time"
)

func main() {
	go CreateGin()

	err := sentry.Init(sentry.ClientOptions{
		Dsn: config.GetSentry(),
		TracesSampler: sentry.TracesSamplerFunc(func(ctx sentry.SamplingContext) sentry.Sampled {
			return sentry.SampledTrue
		}),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	fmt.Printf("\n   ____  ____ ___  ___    ____\n  / __ \\/ __ <  / /   |  /  _/\n / / / / / / / / / /| |  / /  \n/ /_/ / /_/ / / / ___ |_/ /   \n\\____/\\____/_(_)_/  |_/___/   \n                              \n")
	log.Printf("001k AI will be started in few seconds\n")
	log.Printf("v%s\n", config.GetVersion())
	//InitSentry()
	bot.Start()
	return
}
