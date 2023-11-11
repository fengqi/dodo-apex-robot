package config

import (
	"encoding/json"
	"os"
)

var (
	Port        int
	ImageDomain string
	ImagePath   string
	CachePath   string
	Dodo        dodo
	ALS         apexLegendsStatus
	ALT         apexLegendsTracker
	Log         logger
	Season      season
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
	CachePath = config.CachePath
	Dodo = config.Dodo
	ALS = config.ALS
	ALT = config.ALT
	Log = config.Logger
	Season = config.Season
}
