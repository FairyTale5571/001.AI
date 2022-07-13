package database

import (
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type Users struct {
	gorm.Model

	UserId            string
	UserName          string
	UserDiscriminator string
	ConnectDate       time.Time
}

type VerifiedRole struct {
	gorm.Model

	GuildId string
	RoleId  string `gorm:"uniqueIndex"`
}

type TotalMembersChannel struct {
	gorm.Model

	GuildId   string `gorm:"type:varchar(255);not null"`
	ChannelID string `gorm:"type:varchar(255);not null"`
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

type Tickers struct {
	gorm.Model

	GuildId string
	Symbol  string
}

type Tickets struct {
	gorm.Model

	GuildId           string
	UserId            string
	UserName          string
	UserDiscriminator string
	ReasonToOpen      string
}

type Settings struct {
	gorm.Model

	GuildId string
	Setting string
	Value   string
}
