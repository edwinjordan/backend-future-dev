package configs

import (
	"net/http"

	"go-heroku/repositories"
	"go-heroku/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(wilayahRepository *repositories.WilayahRepository) *gin.Engine {
	route := gin.Default()

	route.GET("/show/:link", func(context *gin.Context) {
		link := context.Param("link")

		code := http.StatusOK

		response := services.FindOneWilayahById(link, *wilayahRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	return route
}
