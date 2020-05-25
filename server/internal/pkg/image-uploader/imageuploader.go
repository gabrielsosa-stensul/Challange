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

// Upload recieves a gin context and a field name to manage the upload of that
// field to a folder.
// It returns the filename or any write error encountered.
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

// GetPath returns the path where the images will be stored.
func GetPath() string {
	return "./tmp/"
}

// Validate returns an error if the image does not respect the defined
// extensions (jpg, gif or png), or the width and height that receives as a
// parameters.
func Validate(c *gin.Context, field string, width int, height int) error {
	file, _, _ := c.Request.FormFile(field)

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		return errors.New("Extension must be jpg, gif or png")
	}

	if image.Width != width || image.Height != height {
		return errors.New("Size must be " + strconv.Itoa(width) + "px x " + strconv.Itoa(height) + "px")
	}

	return nil
}
