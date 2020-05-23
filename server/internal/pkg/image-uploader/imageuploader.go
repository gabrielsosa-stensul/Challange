package imageuploader

import (
	"errors"
	"image"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context, field string) (string, error) {
	file, header, _ := c.Request.FormFile(field)

	fileName := strings.TrimSuffix(header.Filename, filepath.Ext(header.Filename))
	fileName = strings.Replace(fileName, " ", "-", -1)
	fileName = fileName + "_" + strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(header.Filename)

	filePath := GetPath() + fileName
	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func GetPath() string {
	return "./tmp/"
}

func Validate(c *gin.Context, field string, width int, height int) error {
	file, _, _ := c.Request.FormFile(field)

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		return errors.New("Image extension must jpg, gif or png")
	}

	if image.Width != width || image.Height != height {
		return errors.New("Image size must be 320px x 320px")
	}

	return nil
}
