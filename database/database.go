package database

import (
	"001.AI/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	var err error
	db, err = ConnectDatabase()
	if err != nil {
		logger.PrintLog("cant open database %s\n", err.Error())
	}
	logger.PrintLog("Database connected")
}

func ConnectDatabase() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("001_database.db"), &gorm.Config{})
	if err != nil {
		logger.PrintLog("cant open database %s\n", err.Error())
		return db, err
	}
	db.AutoMigrate(VerifiedRole{})
	db.AutoMigrate(ConLeaveChannels{})
	db.AutoMigrate(ConnectLogs{})

	return db, nil
}
