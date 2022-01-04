package database

import (
	"001.AI/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type RulesChannels struct {
	gorm.Model
	Id uint32
	GuildId string
	ChannelId string
}

type ConLeaveChannels struct {
	gorm.Model

	Id uint32
	GuildId string
	ChannelId string
}

type ConnectLogs struct {
	gorm.Model

	Id uint32
	GuildId string
	UserId string
	UserName string
	UserDiscriminator string
}

func ConnectDatabase() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("001_database.db"), & gorm.Config{})
	if err != nil {
		logger.PrintLog("cant open database %s\n", err.Error())
		return db, err
	}
	db.AutoMigrate(RulesChannels{})
	db.AutoMigrate(ConLeaveChannels{})
	db.AutoMigrate(ConnectLogs{})

	db.Create(&RulesChannels{})
	db.Create(&RulesChannels{})
	db.Create(&ConnectLogs{})

	return db, nil
}
