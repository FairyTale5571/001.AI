package database

import (
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type Users struct {
	gorm.Model

	GuildId          string
	MemberId         string `gorm:"uniqueIndex"`
	TimeConnect      time.Time
	TimeEndSubscribe time.Time
}

type VerifiedRole struct {
	gorm.Model

	GuildId string
	RoleId  string `gorm:"uniqueIndex"`
}

type ConLeaveChannels struct {
	gorm.Model

	GuildId   string
	ChannelId string `gorm:"uniqueIndex"`
}

type FormsChannels struct {
	gorm.Model

	GuildId   string
	ChannelId string `gorm:"uniqueIndex"`
}

type ConnectLogs struct {
	gorm.Model

	Type              string
	GuildId           string
	UserId            string
	UserName          string
	UserDiscriminator string
}

type Alerts struct {
	gorm.Model

	GuildId   string
	ChannelId string
}

type Tickets struct {
	gorm.Model

	GuildId           string
	UserId            string
	UserName          string
	UserDiscriminator string
	ReasonToOpen      string
}
