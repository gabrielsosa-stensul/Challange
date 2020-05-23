package validator

import (
	"github.com/MarianoArias/challange-api/internal/pkg/image-uploader"
	"github.com/gin-gonic/gin"
)

func ValidatePostHandler(c *gin.Context) gin.H {
	var responseError gin.H

	_, _, err := c.Request.FormFile("image")
	if err != nil {
		responseError = gin.H{"image": "required"}
	}

	err = imageuploader.Validate(c, "image", 320, 320)
	if err != nil {
		responseError = gin.H{"image": err.Error()}
	}

	description := c.Request.FormValue("description")
	if description == "" {
		responseError = gin.H{"description": "required"}
	}

	return responseError
}

func ValidatePatchHandler(c *gin.Context) gin.H {
	var responseError gin.H

	file, _, _ := c.Request.FormFile("image")
	if file != nil {
		err := imageuploader.Validate(c, "image", 320, 320)
		if err != nil {
			responseError = gin.H{"image": err.Error()}
		}
	}

	return responseError
}