package als

import (
	"encoding/json"
	"fengqi/dodo-apex-robot/cache"
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/logger"
	"fengqi/dodo-apex-robot/translate"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
	"time"
)

func init() {
	// todo 初始化http client
}

func GetBridge(player, platform string) (*Bridge, error) {
	bridge := &Bridge{}
	if val, ok := cache.Client.Get("als-bridge-" + player); ok {
		bridge = val.(*Bridge)
	}

	if bridge == nil || bridge.Global.Uid == "" {
		api := fmt.Sprintf("/bridge?auth=%s&player=%s&platform=PC", config.ALS.ApiKey, player)
		bytes, err := requestApi(api)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bytes, bridge)
		if err != nil {
			return nil, err
		}

		cache.Client.Set("als-bridge-"+player, bridge, 10*time.Minute)
	}

	return bridge, nil
}

func GetMapRotation() (*MapRotation, error) {
	api := fmt.Sprintf("/maprotation?auth=%s&version=2", config.ALS.ApiKey)
	bytes, err := requestApi(api)
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
	bytes, err := requestApi(api)
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

func GetRankDistribution(merge bool) map[string]RankData {
	var rankDist = make(map[string]RankData)

	if val, ok := cache.Client.Get("als-rank-dist"); ok {
		rankDist = val.(map[string]RankData)
	}

	if rankDist == nil || len(rankDist) == 0 {
		bytes, err := requestALS("/lib/php/rankdistrib.php?unranked=yes")
		if err != nil {
			panic(err)
		}

		var list []RankDistribution
		err = json.Unmarshal(bytes, &list)
		if err != nil {
			panic(err)
		}

		var indexMap = make(map[string]int)
		for _, v := range list {
			if v.Name == "Rank" {
				for k2, v2 := range v.Data {
					indexMap[v2.(string)] = k2
				}
				break
			}
		}

		for _, v := range list {
			if v.Name == "Rank" {
				continue
			}
			if index, ok := indexMap[v.Name]; ok {
				rankDist[v.Name] = RankData{
					Name:    v.Name,
					Percent: v.Data[index].(float64),
					Total:   v.TotalCount,
				}
			}
		}

		cache.Client.Set("als-rank-dist", rankDist, 24*time.Hour)
	}

	if merge {
		var rankDist2 = make(map[string]RankData)
		for k, _ := range translate.RankMap {
			for k2, v2 := range rankDist {
				if strings.Contains(strings.ToUpper(k2), k) {
					if _, ok := rankDist2[k]; ok {
						rankDist2[k] = RankData{
							Name:    k,
							Percent: rankDist2[k].Percent + v2.Percent,
							Total:   rankDist2[k].Total + v2.Total,
						}
					} else {
						rankDist2[k] = v2
					}
				}
			}
		}
		rankDist = rankDist2
	}

	return rankDist
}

func requestApi(api string) ([]byte, error) {
	//time.Sleep(time.Millisecond * 500)
	logger.Client.Debug("request ALS Api", zap.String("api", api))

	res, err := http.DefaultClient.Get(config.ALS.Host + api)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Client.Error("close response body error", zap.String("api", api), zap.Error(err))
		}
	}(res.Body)

	return io.ReadAll(res.Body)
}

func requestALS(api string) ([]byte, error) {
	//time.Sleep(time.Millisecond * 500)
	logger.Client.Debug("request ALS", zap.String("api", api))

	res, err := http.DefaultClient.Get("https://apexlegendsstatus.com" + api)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Client.Error("close response body error", zap.String("api", api), zap.Error(err))
		}
	}(res.Body)

	return io.ReadAll(res.Body)
}
