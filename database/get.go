package database

func GetWelcomeChannelId(guildId string) ([]string, error) {
	var ret []string
	db.Table("con_leave_channels").Select("channel_id").Where("deleted_at is null and guild_id = ?", guildId).Scan(&ret)
	return ret, nil
}

func GetVerifiedRoles(guildId string) ([]string, error) {
	var ret []string
	db.Table("verified_roles").Select("role_id").Where("deleted_at is null and guild_id = ?", guildId).Scan(&ret)
	return ret, nil
}

func GetFormsChannelId(guildId string) ([]string, error) {
	var ret []string
	db.Table("forms_channels").Select("channel_id").Where("deleted_at is null and guild_id = ?", guildId).Scan(&ret)
	return ret, nil
}

func GetAlertsChannel(guildId string) ([]string, error) {
	var ret []string
	db.Table("alerts").Select("channel_id").Where("deleted_at is null and guild_id = ?", guildId).Scan(&ret)
	return ret, nil
}
