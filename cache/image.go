package cache

import (
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/logger"
	"fengqi/dodo-apex-robot/utils"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/image/draw"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
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

// CacheImage 缓存远程图片到本地
func CacheImage(url string) string {
	hash := utils.Md5(url)
	path := fmt.Sprintf("/%s/%s/%s%s", hash[0:2], hash[2:4], hash, filepath.Ext(url))
	if utils.FileExist(config.ImagePath + path) {
		return config.ImageDomain + "/images" + path
	}

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		logger.Client.Error("download image err", zap.String("url", url), zap.Error(err))
		return url
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Client.Error("CacheImage() close body err", zap.Error(err))
		}
	}(resp.Body)
	if resp.StatusCode != 200 {
		return url
	}

	utils.CheckPath(config.ImagePath + path)
	f, err := os.OpenFile(config.ImagePath+path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		logger.Client.Error("create file err", zap.Error(err))
		return url
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			logger.Client.Error("CacheImage() close file err", zap.Error(err))
		}
	}(f)

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		logger.Client.Error("write file err", zap.Error(err))
		return url
	}

	return config.ImageDomain + "/images" + path
}

// ResizeImage 缩放图片
func ResizeImage(inputUrl string, targetSize int64) string {
	hash := filepath.Base(inputUrl)
	inputPath := fmt.Sprintf("%s/%s/%s/%s", config.ImagePath, hash[0:2], hash[2:4], hash)
	outputPath := filepath.Dir(inputPath) + "/resized_" + hash
	if utils.FileExist(outputPath) {
		return fmt.Sprintf("%s/images/%s/%s/resized_%s", config.ImageDomain, hash[0:2], hash[2:4], hash)
	}

	// 打开图像文件
	file, err := os.Open(inputPath)
	if err != nil {
		return inputUrl
	}
	defer file.Close()

	// 解码图像
	img, _, err := image.Decode(file)
	if err != nil {
		return inputUrl
	}

	// 获取当前图像大小
	currentSize, err := file.Seek(0, os.SEEK_END)
	if err != nil {
		return inputUrl
	}

	// 计算缩放比例
	scaleFactor := float64(targetSize) / float64(currentSize) * 2
	newWidth := int(float64(img.Bounds().Dx()) * scaleFactor)
	newHeight := int(float64(img.Bounds().Dy()) * scaleFactor)

	// 调整图像大小
	resizedImg := resize(img, newWidth, newHeight)

	// 创建输出文件
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return inputUrl
	}
	defer outputFile.Close()

	// 写入输出文件
	ext := filepath.Ext(inputPath)
	switch ext {
	case ".png":
		err = png.Encode(outputFile, resizedImg)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(outputFile, resizedImg, nil)
	case ".gif":
		err = gif.Encode(outputFile, resizedImg, nil)
	}

	return fmt.Sprintf("%s/images/%s/%s/resized_%s", config.ImageDomain, hash[0:2], hash[2:4], hash)
}

// 自定义缩放函数
func resize(img image.Image, width, height int) image.Image {
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(newImg, newImg.Bounds(), img, img.Bounds(), draw.Over, nil)
	return newImg
}
