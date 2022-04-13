package database

func SetWelcomeChannel(guildId, channel string) {
	db.Create(&ConLeaveChannels{
		GuildId:   guildId,
		ChannelId: channel,
	})
}

func RemoveWelcomeChannel(guildId, channel string) {
	db.Where("guild_id = ? and channel_id = ?", guildId, channel).Delete(&ConLeaveChannels{})
}

func SetConnectLog(guildId, userId, userName, userDiscriminator string, mode string) {
	db.Create(&ConnectLogs{
		Type:              mode,
		GuildId:           guildId,
		UserId:            userId,
		UserName:          userName,
		UserDiscriminator: userDiscriminator,
	})
}

func SetVerifiedRole(guildId, roleId string) {
	db.Create(&VerifiedRole{
		GuildId: guildId,
		RoleId:  roleId,
	})
}

func SetFormChannel(guildId, channelId string) {
	db.Create(&FormsChannels{
		GuildId:   guildId,
		ChannelId: channelId,
	})
}

func SetAlertChannel(guildId, channelId string) {
	db.Create(&Alerts{
		GuildId:   guildId,
		ChannelId: channelId,
	})
}

func RemoveVerifiedRole(guildId, roleId string) {
	db.Where("guild_id = ? and role_id = ?", guildId, roleId).Delete(&VerifiedRole{})
}

func InsertNewTicker(ticker, guildId string) {
	db.Create(&Tickers{
		GuildId: guildId,
		Symbol:  ticker,
	})
}
