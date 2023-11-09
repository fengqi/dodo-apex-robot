package main

import (
	"fengqi/dodo-apex-robot/cache"
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/webhook"
	"fmt"
	"net/http"
	"strconv"
)

func init() {
	config.Load("./config.json")
}

func main() {
	http.HandleFunc("/images/", cache.ImageHandler)
	http.HandleFunc("/webhook", webhook.Handler)

	fmt.Printf("Starting server at port %d\n", config.Port)
	if err := http.ListenAndServe(":"+strconv.Itoa(config.Port), nil); err != nil {
		panic(err)
	}
}
