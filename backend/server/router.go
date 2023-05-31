package server

import (
	"github.com/LebrancWorkshop/ShortenURL-Workshop-BorntoDev/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(controller *controllers.Controller) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	api := router.Group("/api")
	{
		api.GET("/ping", controller.Ping)
		v1 := api.Group("/v1")
		{
			v1.POST("/url", controller.CreateShortURL)
			v1.GET("/url", controller.GetShortURL)
		}
	}



	return router
}
