package validator

import (
	"github.com/MarianoArias/Challange/server/internal/pkg/image-uploader"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidatePostHandler does all the necessary validations to post an item.
// It returns a gin interface with errors encountered.
func ValidatePostHandler(c *gin.Context) []Error {
	var errors []Error
	var error Error

	_, _, err := c.Request.FormFile("image")
	if err != nil {
		error.Field = "image"
		error.Message = "required"
		errors = append(errors, error)
	} else {
		err = imageuploader.Validate(c, "image", 320, 320)
		if err != nil {
			error.Field = "image"
			error.Message = err.Error()
			errors = append(errors, error)
		}
	}

	description := c.Request.FormValue("description")
	if description == "" {
		error.Field = "description"
		error.Message = "required"
		errors = append(errors, error)
	} else {
		if len(description) > 300 {
			error.Field = "description"
			error.Message = "Must be maximum 300 characters long"
			errors = append(errors, error)
		}
	}

	return errors
}

// ValidatePostHandler does all the necessary validations to patch an item.
// It returns a gin interface with errors encountered.
func ValidatePatchHandler(c *gin.Context) []Error {
	var errors []Error
	var error Error

	file, _, _ := c.Request.FormFile("image")
	if file != nil {
		err := imageuploader.Validate(c, "image", 320, 320)
		if err != nil {
			error.Field = "image"
			error.Message = err.Error()
			errors = append(errors, error)
		}
	}

	description := c.Request.FormValue("description")
	if len(description) > 300 {
		error.Field = "description"
		error.Message = "Must be maximum 150 characters long"
		errors = append(errors, error)
	}

	return errors
}
