package bot

import (
	"001.AI/logger"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"fd_yes": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			yesAnswer(true, s, i)
			return
		},
	}
)

func yesAnswer(a bool, s *discordgo.Session, i *discordgo.InteractionCreate) {
	var cont string
	if !a {
		cont = "Вам нужно согласится с правилами чтобы получить полный доступ к серверу"
		removeVerifiedRoles(i.GuildID, i.Interaction.Member.User)
	} else {
		cont = "Спасибо! Доступ к серверу выдан!"
		giveVerifiedRoles(i.GuildID, i.Interaction.Member.User)
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: cont,
			Flags:   1 << 6,
		},
	})
	if err != nil {
		logger.PrintLog("error: cant create respond %s\n", err.Error())
		return
	}

	if err := s.InteractionResponseDelete(s.State.User.ID, i.Interaction); err != nil {
		fmt.Printf("error: cant delete response %s\n", err.Error())
	}

}

func printRules(chanelId string) {
	rules := fmt.Sprintf("```" +
		"Присоединяясь к серверу 001k.crypto, вы соглашаетесь следовать приведенным ниже правилам, а также соглашаетесь с тем, что любое воспроизведение / продажа контента / или привлечение участников может привести к судебному иску. Любое их нарушение приведет к предупреждению или бану. Основная цель - сохранить качество и эксклюзивность комнаты!\n\n" +
		"1. Относитесь к каждому участнику сервера с уважением, как к семье.\n\n" +
		"2. НЕ троллить, не флудить, не спамить нигде и никому. БЕЗ ИСКЛЮЧЕНИЙ!\n\n" +
		"3. Проявляйте должную осмотрительность в отношении своих сделок! Не следуй слепо! ТОЛЬКО ВЫ несете ответственность за свои сделки.\n\n" +
		"4. 001k не может и не скажет вам когда и по какой цене покупать или продавать активы в этой комнате.\n\n" +
		"5. Никакая информация, опубликованная и распространенная внутри этого сообщества, не является обещанием или гарантией какого-либо конкретного результата.\n\n" +
		"6. НЕ ЗАНИМЕЙТЕСЬ САМОПРОДВИЖЕНИЕМ. Это означает, что нельзя публиковать ссылки социальных сетей, телеграм каналов, материалы других трейдеров и т.д.\n\n" +
		"7. НЕ пытайтесь продавать какие-либо продукты участникам.\n\n" +
		"8. НЕ публикуйте такие вещи, как «Покупайте монету XYZ, она взорвется / она будет расти!» «Не пропустите!» И т д. Комментарии, подобные этим, будут истолкованы как имеющие скрытый мотив или создающие FOMO." +
		"```")

	data := &discordgo.MessageSend{
		Content: rules,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						// Label is what the user will see on the button.
						Label: "Согласен",
						// Style provides coloring of the button. There are not so many styles tho.
						Style: discordgo.SuccessButton,
						// Disabled allows bot to disable some buttons for users.
						Disabled: false,
						// CustomID is a thing telling Discord which data to send when this button will be pressed.
						CustomID: "fd_yes",
						Emoji: discordgo.ComponentEmoji{
							Name: "👍",
						},
					},
				},
			},
		},
	}

	_, err := s.ChannelMessageSendComplex(chanelId, data)
	if err != nil {
		logger.PrintLog("cant print rules %s\n", err.Error())
		return
	}
}

func printRules2(chanelId string) {
	rules := fmt.Sprintf("```" +
		"ПРАВИЛА\n" +
		"1 - При попытке заскамить кого-либо будет БАН.\n" +
		"2 - За оскорбления таймаут, потом БАН.\n" +
		"3 - Без политики! за разжигание ненависти будет БАН.\n" +
		"4 - Без NSFW.\n" +
		"5 - Если модератор говорит вам перестать делать что-либо, вы должны перестать.\n" +
		"Чтобы получить доступ к каналам, нажмите на ✅ в реакциях снизу" +
		"```")

	data := &discordgo.MessageSend{
		Content: rules,
	}

	msg, err := s.ChannelMessageSendComplex(chanelId, data)
	if err != nil {
		logger.PrintLog("cant print rules %s\n", err.Error())
		return
	}
	err = s.MessageReactionAdd(msg.ChannelID, msg.ID, "✅")
	if err != nil {
		logger.PrintLog("cant print rules %s\n", err.Error())
		return
	}
}
