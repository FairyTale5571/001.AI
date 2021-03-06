package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

func init() {
	err := readConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

type Cfg struct {
	Token         string `yaml:"token"`
	Prefix        string `yaml:"prefixCommand"`
	Version       string `yaml:"version"`
	SentryKey     string `yaml:"sentry_key"`
	Port          string `yaml:"port"`
	TwitchToken   string `yaml:"twitch_token"`
	TelegramToken string `yaml:"telegram_token"`

	Ip       string `yaml:"ip"`
	DbPort   int    `yaml:"dbPort"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

var conf = Cfg{}

func readConfig() error {
	f, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Fatalf("loadConfig -> %v", err)
		return err
	}
	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		log.Fatalf("loadConfig -> %v", err)
		return err
	}
	return nil
}

func saveConfig() {
	data, err := yaml.Marshal(&conf)
	if err != nil {
		log.Fatalf("saveConfig -> %v", err)
		return
	}
	err = ioutil.WriteFile("config.yml", data, 0)
	if err != nil {
		log.Fatal("saveConfig -> %v", err)
	}
}
