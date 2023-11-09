package als

import (
	"encoding/json"
	"fengqi/dodo-apex-robot/config"
	"fmt"
	"io"
	"net/http"
)

func init() {
	// todo 初始化http client
}

func GetBridge(player, platform string) (Bridge, error) {
	api := fmt.Sprintf("/bridge?auth=%s&player=%s&platform=PC", config.ALS.ApiKey, player)
	bytes, err := sendRequest(api)
	if err != nil {
		panic(err)
	}

	var bridge Bridge
	err = json.Unmarshal(bytes, &bridge)
	if err != nil {
		panic(err)
	}

	return bridge, nil
}

func GetMapRotation() (MapRotation, error) {
	api := fmt.Sprintf("/maprotation?auth=%s", config.ALS.ApiKey)
	bytes, err := sendRequest(api)
	if err != nil {
		panic(err)
	}

	var mapRotation MapRotation
	err = json.Unmarshal(bytes, &mapRotation)
	if err != nil {
		panic(err)
	}

	return mapRotation, nil
}

func sendRequest(api string) ([]byte, error) {
	//time.Sleep(time.Millisecond * 500)

	res, err := http.DefaultClient.Get(config.ALS.Host + api)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	return io.ReadAll(res.Body)
}
