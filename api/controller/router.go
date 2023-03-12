package controller

import (
	docs "WordTree/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	wordController := NewWordController()
	api := r.Group("/api")
	{
		api.GET("/word", wordController.Search)
		api.POST("/word", wordController.Create)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":3000")
}
