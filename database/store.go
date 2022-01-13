package database

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type VerifiedRole struct {
	gorm.Model
	Id      int64 `gorm:"primaryKey; not null"`
	GuildId string
	RoleId  string `gorm:"uniqueIndex"`
}

type ConLeaveChannels struct {
	gorm.Model

	Id        int64 `gorm:"primaryKey; not null"`
	GuildId   string
	ChannelId string `gorm:"uniqueIndex"`
}

type ConnectLogs struct {
	gorm.Model

	Id                int64 `gorm:"primaryKey; not null"`
	Type              string
	GuildId           string
	UserId            string
	UserName          string
	UserDiscriminator string
}

type Tickets struct {
	gorm.Model

	Id                int64 `gorm:"primaryKey; not null"`
	GuildId           string
	UserId            string
	UserName          string
	UserDiscriminator string
	ReasonToOpen      string
}
