package healthhandler

import (
	"net/http"

	"github.com/MarianoArias/challange-api/internal/pkg/mongodb"
	"github.com/gin-gonic/gin"
)

type Response map[string]interface{}

const (
	UP   = "UP"
	DOWN = "DOWN"
)

func HealthHandler(c *gin.Context) {
	generalStatus, mongodbStatus := UP, UP

	if err := mongodb.Ping(); err != nil {
		generalStatus = DOWN
		mongodbStatus = DOWN
	}

	a := Response{
		"status": Response{
			"code": generalStatus,
		},
		"details": Response{
			"mongodb": Response{
				"status": Response{
					"code": mongodbStatus,
				},
			},
		},
	}

	c.JSON(http.StatusOK, a)
}
