package bot

import (
	"001.AI/database"
	"001.AI/logger"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func isAdmin(user *discordgo.User, channel string) bool {
	_, err := s.UserChannelPermissions(user.ID, channel)
	if err != nil {
		logger.PrintLog("cant check permission %s\n", err.Error())
		return false
	}
	return false
}

func giveVerifiedRoles(guildId string, user *discordgo.User) {
	roles, err := database.GetVerifiedRoles(guildId)
	if err != nil {
		logger.PrintLog("get rule channel error: %s\n", err.Error())
	}
	if len(roles) == 0 {
		logger.PrintLog("connect messages not configured\n")
		return
	}
	for _, role := range roles {
		giveRole(guildId, user.ID, role)
	}
}

func removeVerifiedRoles(guildId string, user *discordgo.User) {
	roles, err := database.GetVerifiedRoles(guildId)
	if err != nil {
		logger.PrintLog("get rule channel error: %s\n", err.Error())
	}
	if len(roles) == 0 {
		logger.PrintLog("connect messages not configured\n")
		return
	}
	for _, role := range roles {
		removeRole(guildId, user.ID, role)
	}
}

func giveRole(guildId string, uid string, roleId string) {
	fmt.Printf("giving role \n")
	_, err := s.GuildMember(guildId, uid)
	if err != nil {
		logger.PrintLog("error: give role - user undefined\n")
		return
	}
	err = s.GuildMemberRoleAdd(guildId, uid, roleId)
	if err != nil {
		logger.PrintLog("error: give role - %s\n", err.Error())
		return
	}
}

func removeRole(guildId string, uid string, roleId string) {
	_, err := s.GuildMember(guildId, uid)
	if err != nil {
		logger.PrintLog("error: remove role - user undefined\n")
		return
	}
	err = s.GuildMemberRoleRemove(guildId, uid, roleId)
	if err != nil {
		logger.PrintLog("error: remove role - %s\n", err.Error())
		return
	}
}

// сохранить роль
func setVerifiedRole(s *discordgo.Session, i *discordgo.InteractionCreate) {
	role := i.ApplicationCommandData().Options[0].RoleValue(s, "")
	database.SetVerifiedRole(i.GuildID, role.ID)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Роль " + role.Mention() + " установлена",
			Flags:   1 << 6,
		},
	})
}

// удалить роль
func removeVerifiedRole(s *discordgo.Session, i *discordgo.InteractionCreate) {
	role := i.ApplicationCommandData().Options[0].RoleValue(s, "")
	database.RemoveVerifiedRole(i.GuildID, role.ID)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Роль " + role.Mention() + " удалена",
			Flags:   1 << 6,
		},
	})
}

// установить канал с логами
func setWelcomeChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channel := i.ApplicationCommandData().Options[0].ChannelValue(s)
	database.SetWelcomeChannel(i.GuildID, channel.ID)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Канал " + channel.Mention() + " установлен",
			Flags:   1 << 6,
		},
	})
}

// установить канал с логами
func setFormChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channel := i.ApplicationCommandData().Options[0].ChannelValue(s)
	database.SetFormChannel(i.GuildID, channel.ID)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Канал " + channel.Mention() + " установлен",
			Flags:   1 << 6,
		},
	})
}

// убрать канал с логами
func removeWelcomeChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channel := i.ApplicationCommandData().Options[0].ChannelValue(s)
	database.RemoveWelcomeChannel(i.GuildID, channel.ID)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Канал " + channel.Mention() + " удален",
			Flags:   1 << 6,
		},
	})
}
