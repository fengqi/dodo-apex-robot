package utils

import (
	"encoding/json"
	"fengqi/dodo-apex-robot/message"
	"fmt"
	"net/http"
)

func Response(w http.ResponseWriter) {
	bytes, _ := json.Marshal(&message.Response{})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}

func ResponseData(w http.ResponseWriter, data interface{}) {
	fmt.Printf("response: %v\n", data)
	bytes, _ := json.Marshal(&message.Response{Data: data})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}
