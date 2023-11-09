package dodo

import "net/http"

var (
	ApiSendMessage = "/api/v2/channel/message/send"
)

var client http.Client

func init() {

}
