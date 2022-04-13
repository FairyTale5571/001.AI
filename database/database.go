package database

import (
	"001.AI/config"
	"001.AI/logger"
	"fmt"
	"gorm.io/driver/mysql"
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

	dbc := config.GetDB()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbc.User, dbc.Password, dbc.Ip, dbc.Port, dbc.Database)
	db, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN: dsn,
		},
	))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(VerifiedRole{})
	db.AutoMigrate(ConLeaveChannels{})
	db.AutoMigrate(ConnectLogs{})
	db.AutoMigrate(FormsChannels{})
	db.AutoMigrate(Users{})
	db.AutoMigrate(Tickers{})

	return db, nil
}
