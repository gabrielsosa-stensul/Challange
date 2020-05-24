package controller

import (
	"net/http"
	"strconv"

	"github.com/MarianoArias/Challange/server/internal/app/items/model"
	"github.com/MarianoArias/Challange/server/internal/app/items/validator"
	"github.com/MarianoArias/Challange/server/internal/pkg/image-uploader"
	"github.com/gin-gonic/gin"
)

// @Description Get Image
// @Produce image/png
// @Produce image/gif
// @Produce image/jpeg
// @Param fileName path string true "Image src"
// @Success 200
// @Failure 404
// @Router /images/ [get]
func GetImageHandler(c *gin.Context) {
	c.File(imageuploader.GetPath() + c.Param("fileName"))
	return
}

// @Description Get Items
// @Produce json
// @Success 200
// @Success 204
// @Failure 500
// @Router /items/ [get]
func CgetHandler(c *gin.Context) {
	_items, err := model.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	
	if len(*_items) == 0 {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	var items []model.Item
	for _, item := range *_items {
		item.Image = "/images/" + item.Image
		items = append(items, item)
	}

	c.JSON(http.StatusOK, items)
	return
}

// @Description Delete Item
// @Param id path int true "Item id"
// @Success 204
// @Failure 500
// @Failure 404
// @Router /items/{id} [get]
func DeleteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := model.Find(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	err = model.Delete(*item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusNoContent, nil)
	return
}

// @Description Get Item
// @Produce json
// @Param id path int true "Item id"
// @Success 200
// @Failure 404
// @Router /items/{id} [get]
func GetHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := model.Find(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	item.Image = "/images/" + item.Image

	c.JSON(http.StatusOK, item)
	return
}

// @Description Post Item
// @Accept multipart/form-data
// @Success 201
// @Failure 400
// @Failure 500
// @Router /items/ [post]
func PostHandler(c *gin.Context) {
	errorResponse := validator.ValidatePostHandler(c)
	if errorResponse != nil {
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	fileName, err := imageuploader.Upload(c, "image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var item model.Item
	item.Image = fileName
	item.Description = c.Request.FormValue("description")

	err = model.Persist(&item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Header("X-Resource-Id", strconv.Itoa(item.ID))
	c.Status(http.StatusCreated)
	return
}

// @Description Post Item
// @Accept multipart/form-data
// @Param id path int true "Item id"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /items/{id} [patch]
func PatchHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := model.Find(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	errorResponse := validator.ValidatePatchHandler(c)
	if errorResponse != nil {
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	file, _, _ := c.Request.FormFile("image")
	if file != nil {
		fileName, err := imageuploader.Upload(c, "image")
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		item.Image = fileName
	}

	description := c.Request.FormValue("description")
	if description != "" {
		item.Description = description
	}

	err = model.Update(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	order, _ := strconv.Atoi(c.Request.FormValue("order"))
	if order != 0 && item.Order != order {
		err = model.SwitchOrder(item.Order, order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	}

	c.JSON(http.StatusNoContent, item)
	return
}
