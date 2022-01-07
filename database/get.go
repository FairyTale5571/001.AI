package database

import "fmt"

func GetWelcomeChannelId(guildId string) (string,error) {
	var ret string
	tx := db.Exec("SELECT channel_id from con_leave_channels WHERE guild_id = ? LIMIT 1",guildId)

	tx.Scan(&ret)
	if ret != "" {
		return ret, nil
	}
	return "", fmt.Errorf("welcome channel is undefined")
}

func GetRuleChannelId(guildId string) (string,error) {

	return "", nil
}
