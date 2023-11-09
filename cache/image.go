package cache

import (
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	path := "data" + r.URL.Path

	fi, err := os.Stat(path)
	if err != nil || fi.IsDir() {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, path)
}

func CacheImage(url string) string {
	hash := utils.Md5(url)
	path := fmt.Sprintf("/%s/%s/%s%s", hash[0:2], hash[2:4], hash, filepath.Ext(url))
	if utils.FileExist(config.ImagePath + path) {
		return config.ImageDomain + "/images" + path
	}

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	if resp.StatusCode != 200 {
		return url
	}

	utils.CheckPath(config.ImagePath + path)
	f, err := os.OpenFile(config.ImagePath+path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		panic(err)
	}

	return config.ImageDomain + "/images" + path
}
