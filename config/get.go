package config

// GetToken получить токен с конфига
func GetToken() string {
	return conf.Token
}

// GetPrefix получить префикс для вызова команды
func GetPrefix() string {
	return conf.Prefix
}

// GetVersion получить версию бота
func GetVersion() string {
	return conf.Version
}

func GetSentry() string {
	return conf.SentryKey
}

func GetPort() string {
	return conf.Port
}

func GetTwitchToken() string {
	return conf.TwitchToken
}

type dbStruct struct {
	Ip       string
	Port     int
	Database string
	User     string
	Password string
}

func GetDB() dbStruct {
	ret := dbStruct{
		Ip:       conf.Ip,
		Port:     conf.DbPort,
		Database: conf.Database,
		User:     conf.User,
		Password: conf.Password,
	}
	return ret
}

func GetTelegramToken() string {
	return conf.TelegramToken
}
