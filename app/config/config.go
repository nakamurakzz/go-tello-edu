package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfList struct {
	LogFileName string
	Address     string
	Port        int
}

var Config ConfList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Println("config.ini load error")
		os.Exit(1)
	}

	Config = ConfList{
		LogFileName: cfg.Section("gotello").Key("log_file_name").String(),
		Address:     cfg.Section("web").Key("address").String(),
		Port:        cfg.Section("web").Key("port").MustInt(),
	}
}
