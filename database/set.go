package database

import "time"

func SetWelcomeChannel(guildId, channel string) {
	db.Create(&ConLeaveChannels{
		GuildId:   guildId,
		ChannelId: channel,
	})
}

func SetConnectLog(guildId, userId, userName, userDiscriminator string, mode string) {
	db.Create(&ConnectLogs{
		Id: 			   time.Now().Unix(),
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
