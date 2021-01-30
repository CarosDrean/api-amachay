package models

type Configuration struct {
	Engine   string
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

type ConfigurationTelegram struct {
	Token  string
	ChatId string
}
