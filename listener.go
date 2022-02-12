package main

import (
	"001.AI/bot"
	"001.AI/config"
	"001.AI/logger"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func mainRouter() {
	router.POST("/submit_form", unmarshallForm)
}

func CreateGin() {
	router = gin.Default()
	logger.PrintLog("gin opened %v\n", router.AppEngine)
	mainRouter()
	router.Run("0.0.0.0" + config.GetPort())
}

func unmarshallForm(c *gin.Context) {
	_body, err := c.GetRawData()
	if err != nil {
		logger.PrintLog("cant get raw data %s\n", err.Error())
		c.JSON(500, nil)
		return
	}
	fmt.Println(_body)

	var body map[string]string

	if err := json.Unmarshal(_body, &body); err != nil {
		logger.PrintLog("cant unmarshall %s\n", err.Error())
		c.JSON(500, nil)
		return
	}
	fmt.Println(body)
	bot.CreateNewForm(body)
	c.JSON(200, nil)
}
