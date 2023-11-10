package als

import (
	"encoding/json"
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/logger"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func init() {
	// todo 初始化http client
}

func GetBridge(player, platform string) (*Bridge, error) {
	api := fmt.Sprintf("/bridge?auth=%s&player=%s&platform=PC", config.ALS.ApiKey, player)
	bytes, err := sendRequest(api)
	if err != nil {
		return nil, err
	}

	bridge := &Bridge{}
	err = json.Unmarshal(bytes, bridge)
	if err != nil {
		return nil, err
	}

	return bridge, nil
}

func GetMapRotation() (*MapRotation, error) {
	api := fmt.Sprintf("/maprotation?auth=%s&version=2", config.ALS.ApiKey)
	bytes, err := sendRequest(api)
	if err != nil {
		return nil, err
	}

	mapRotation := &MapRotation{}
	err = json.Unmarshal(bytes, mapRotation)
	if err != nil {
		return nil, err
	}

	return mapRotation, nil
}

func GetCrafting() ([]Bundle, error) {
	api := fmt.Sprintf("/crafting?auth=%s", config.ALS.ApiKey)
	bytes, err := sendRequest(api)
	if err != nil {
		return nil, err
	}

	var list []Bundle
	err = json.Unmarshal(bytes, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func sendRequest(api string) ([]byte, error) {
	//time.Sleep(time.Millisecond * 500)
	logger.Zap().Debug("request ALS", zap.String("api", api))

	res, err := http.DefaultClient.Get(config.ALS.Host + api)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Zap().Error("close response body error", zap.String("api", api), zap.Error(err))
		}
	}(res.Body)

	return io.ReadAll(res.Body)
}
