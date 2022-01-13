package bot

import (
	"001.AI/logger"
	"github.com/bwmarrin/discordgo"
	"time"
)

type Ticket struct {
	TicketId  string
	ChannelId string
	CreatorId string
	Timer     time.Time
}

var (
	ticketCommands = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"create-ticket": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			createTicketCommand(s, i)
		},
		"close-ticket": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			closeTicketCommand(s, i)
		},
		"add-user-to-ticket": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			addUserCommand(s, i)
		},
	}
	// глобальный объект с тикетом
	ticket *Ticket
)

func addUserCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	chanel := i.Interaction.ChannelID
	ticket, err := findTicket(chanel)
	if err != nil {
		logger.PrintLog("cant find ticket %s\n", err.Error())
		return
	}

	ticket.addUser(i.Interaction.User.ID)
}

func createTicketCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse
	response = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Выберите причину открытия тикета",
			Flags:   1 << 6,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID:    "select-ticket",
							Placeholder: "Выберите из списка ниже",
							Options: []discordgo.SelectMenuOption{
								{
									Label:       "Оплата обучения",
									Value:       "payment",
									Default:     false,
									Description: "Решение проблем с оплатой обучения",
									Emoji: discordgo.ComponentEmoji{
										Name: "",
									},
								},
								{
									Label:       "Проверка ДЗ",
									Value:       "home-work",
									Default:     false,
									Description: "Сдать на проверку домашнее задание",
									Emoji: discordgo.ComponentEmoji{
										Name: "",
									},
								},
								{
									Label:       "Проблемы нет в списке",
									Value:       "etc",
									Description: "моей проблемы нет в списке",
									Default:     false,
									Emoji: discordgo.ComponentEmoji{
										Name: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		logger.PrintLog("cant create ticket: %s\n", err.Error())
		return
	}
}

func closeTicketCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	chanel := i.Interaction.ChannelID
	ticket, err := findTicket(chanel)
	if err != nil {
		logger.PrintLog("cant find ticket %s\n", err.Error())
		return
	}

	ticket.closeTicket()

}

// open ticket
func openTicket() *Ticket {

	ticket := &Ticket{
		TicketId:  "",
		ChannelId: "",
		CreatorId: "",
	}

	return ticket
}

// method insert ticket
func (t *Ticket) insertTicket() {

}

// add user to ticket
func (t *Ticket) addUser(userId string) {

}

// find ticket
func findTicket(chanelId string) (*Ticket, error) {

	return ticket, nil
}

// method close ticket
func (t *Ticket) closeTicket() {

}

// set timer refresh
func (t *Ticket) setTimer() {

}
