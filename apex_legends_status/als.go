package als

import (
	"encoding/json"
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/logger"
	"fengqi/dodo-apex-robot/translate"
	"fengqi/dodo-apex-robot/utils"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	cache = &sync.Map{}
)

func init() {
	// todo 初始化http client
}

func GetBridge(player, platform string) (*Bridge, error) {
	api := fmt.Sprintf("/bridge?auth=%s&player=%s&platform=PC", config.ALS.ApiKey, player)
	bytes, err := requestApi(api)
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
	var rankMap = make(map[string]RankData)

	t, ok := cache.Load("rank-dist-time")
	if ok {
		if t.(int64)+86400 > time.Now().In(utils.GetLocation()).Unix() {
			rank, ok := cache.Load("rank-dist")
			if ok {
				rankMap = rank.(map[string]RankData)
			}
		}
	}

	if rankMap == nil || len(rankMap) == 0 {
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
				rankMap[v.Name] = RankData{
					Name:    v.Name,
					Percent: v.Data[index].(float64),
					Total:   v.TotalCount,
				}
			}
		}

		cache.Store("rank-dist-time", time.Now().In(utils.GetLocation()).Unix())
		cache.Store("rank-dist", rankMap)
	}

	if merge {
		var rankMap2 = make(map[string]RankData)
		for k, _ := range translate.RankMap {
			for k2, v2 := range rankMap {
				if strings.Contains(strings.ToUpper(k2), k) {
					if _, ok := rankMap2[k]; ok {
						rankMap2[k] = RankData{
							Name:    k,
							Percent: rankMap2[k].Percent + v2.Percent,
							Total:   rankMap2[k].Total + v2.Total,
						}
					} else {
						rankMap2[k] = v2
					}
				}
			}
		}
		rankMap = rankMap2
	}

	return rankMap
}

func requestApi(api string) ([]byte, error) {
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

func requestALS(api string) ([]byte, error) {
	//time.Sleep(time.Millisecond * 500)
	logger.Zap().Debug("request ALS", zap.String("api", api))

	res, err := http.DefaultClient.Get("https://apexlegendsstatus.com" + api)
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
