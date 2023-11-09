package main

import (
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/webhook"
	"fmt"
	"net/http"
)

func init() {
	config.Load("./config.json")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/webhook", webhook.Handler)

	fmt.Printf("Starting server at port 2080\n")
	if err := http.ListenAndServe(":2080", nil); err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello, World!"))
}
