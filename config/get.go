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