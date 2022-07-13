package main

import (
	"001.AI/bot"
	"001.AI/config"
	"001.AI/logger"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strings"
)

var router *gin.Engine

const (
	pathToRecords = "stream_records"
)

func mainRouter() {
	router.POST("/submit_form", unmarshallForm)
	router.GET("/streams/:filename", func(ctx *gin.Context) {
		fileName := ctx.Param("filename")
		targetPath := filepath.Join(pathToRecords, fileName)
		if !strings.HasPrefix(filepath.Clean(targetPath), pathToRecords) {
			ctx.String(403, "Look like you attacking me")
			return
		}
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Disposition", "attachment; filename="+fileName)
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.File(targetPath)
	})
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
