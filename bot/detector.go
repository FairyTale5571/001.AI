package bot

import (
	"001.AI/database"
	"001.AI/embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	ErrorInvalidTicker          = "тикер не обнаружен"
	ErrorInvalidTickerSymbols   = "введенные символы не соответствуют тикерам на бирже, прим. BTCUSDT/C98USDT/UAHUSD"
	ErrorInvalidTickerUndefined = "неизвестная ошибка, обратитесь к администратору"

	BinanceURLTickerArray = "https://api3.binance.com/api/v3/ticker/24hr?symbols="
	BinanceURLTicker      = "https://api3.binance.com/api/v3/ticker/24hr?symbol="
)

func ticker(ticker string) *ExchangeInfo {
	var e ExchangeInfo
	var client http.Client

	resp, err := client.Get(fmt.Sprintf("%s%s", BinanceURLTicker, ticker))

	if err != nil {
		log.Println(err)
		return nil
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return nil
		}
		if !json.Valid(bodyBytes) {
			fmt.Println("json is invalid")
			return nil
		}
		err = json.Unmarshal(bodyBytes, &e)

		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	return &e
}

func checkTicker(ticker string) error {

	var client http.Client

	const url = "https://api.binance.com/api/v3/exchangeInfo?symbol="
	resp, err := client.Get(url + strings.ToUpper(ticker))

	if err != nil {
		log.Println(err)
		return err
	}
	bodyBytes, err := io.ReadAll(resp.Body)

	fmt.Printf("body: %v \n", string(bodyBytes))
	fmt.Printf("resp code: %d\n", resp.StatusCode)
	switch resp.StatusCode {
	case http.StatusOK:
		break
	case 400:
		return errors.New(ErrorInvalidTicker)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("status code: %d\n", resp.StatusCode)
		return errors.New(ErrorInvalidTickerUndefined)
	}
	return nil
}

func addTicker(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ticker := i.Interaction.ApplicationCommandData().Options[0].StringValue()

	if err := checkTicker(ticker); err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Тикер имеет ошибку: " + err.Error(),
			},
		})
		return
	}
	database.InsertNewTicker(strings.ToUpper(ticker), i.Interaction.GuildID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Тикер: " + ticker + " добавлен в список",
		},
	})
}

func printTicker(s *discordgo.Session, i *discordgo.InteractionCreate) {
	tick := i.Interaction.ApplicationCommandData().Options[0].StringValue()

	if err := checkTicker(tick); err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Тикер имеет ошибку: " + err.Error(),
				Flags:   1 << 6,
			},
		})
		return
	}
	info := ticker(tick)
	e := embed.NewEmbed()
	e.MessageEmbed.Type = discordgo.EmbedTypeLink
	e.SetDescription("Информация по тикеру: " + tick)
	e.SetColor(0xF0B90B)
	e.AddField("Последняя цена:", trimZeros(info.LastPrice))
	e.AddField("Максимум за сутки: ", trimZeros(info.HighPrice))
	e.AddField("Минимум за сутки: ", trimZeros(info.LowPrice))

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				e.MessageEmbed,
			},
		},
	})
}
