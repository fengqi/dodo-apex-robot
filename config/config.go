package config

import (
	"encoding/json"
	"os"
)

var (
	Dodo dodo
	ALS  apexLegendsStatus
	ALT  apexLegendsTracker
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

	Dodo = config.Dodo
	ALS = config.ALS
	ALT = config.ALT
}
