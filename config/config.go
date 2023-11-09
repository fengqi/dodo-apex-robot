package config

import (
	"encoding/json"
	"os"
)

var (
	Port        int
	ImageDomain string
	ImagePath   string
	Dodo        dodo
	ALS         apexLegendsStatus
	ALT         apexLegendsTracker
)

func Load(file string) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		panic(err)
	}

	Port = config.Port
	ImageDomain = config.ImageDomain
	ImagePath = config.ImagePath
	Dodo = config.Dodo
	ALS = config.ALS
	ALT = config.ALT
}
