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

	YFIKey    = "aDRnN0KL7m52V8tTBhGOA2Z69brK92fD5XXIsSXI"
	SP500     = "%5EGSPC"
	YFIApiUrl = "https://yfapi.net/v7/finance/options/"
)

func sp500() *YFI {
	var client http.Client
	var y YFI
	req, err := http.NewRequest("GET", YFIApiUrl+SP500, nil)
	req.Header.Set("X-API-KEY", YFIKey)
	resp, err := client.Do(req)

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("status code: %d\n", resp.StatusCode)
		return nil
	} else {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return nil
		}
		if !json.Valid(bodyBytes) {
			fmt.Println("json is invalid")
			return nil
		}
		err = json.Unmarshal(bodyBytes, &y)

		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	return &y
}

func tickers(guildId string) *ExchangeInfos {

	var e ExchangeInfos
	var client http.Client

	symbols, err := database.GetTickers(guildId)
	fmt.Printf("%s%s\n", BinanceURLTickerArray, struct2JSON(symbols))

	resp, err := client.Get(fmt.Sprintf("%s%s", BinanceURLTickerArray, struct2JSON(symbols)))
	defer resp.Body.Close()
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

func ticker(ticker string) *ExchangeInfo {
	var e ExchangeInfo
	var client http.Client

	resp, err := client.Get(fmt.Sprintf("%s%s", BinanceURLTicker, ticker))
	defer resp.Body.Close()
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
	defer resp.Body.Close()
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

func printPrices(guildId, channelId string) {
	t := *tickers(guildId)
	fmt.Printf("tickers: %v\n", t)

	e := embed.NewEmbed()
	e.MessageEmbed.Type = discordgo.EmbedTypeImage
	e.SetColor(0xF0B90B)
	e.SetTitle("Утренний отчет")
	for _, i := range t {
		e.AddField(i.Symbol, strings.TrimRight(i.LastPrice, "0"))
	}
	sp := sp500().OptionChain.Result[0].Quote
	e.AddField(sp.ShortName, fmt.Sprintf("%.2f", sp.RegularMarketPrice))
	s.ChannelMessageSendEmbed(channelId, e.MessageEmbed)

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
