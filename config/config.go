package config

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type cfg struct {
	token   string `yaml:"token"`
}

var config = cfg{}

func ReadConfig() error {
	f, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Fatalf("loadConfig -> %v", err)
		return err
	}
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		log.Fatalf("loadConfig -> %v", err)
		return err
	}
	return nil
}

func saveConfig() {
	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("saveConfig -> %v", err)
		return
	}
	err = ioutil.WriteFile("config.yml", data, 0)
	if err != nil {
		log.Fatal("saveConfig -> %v", err)
	}
}