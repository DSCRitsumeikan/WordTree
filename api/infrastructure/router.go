package infrastructure

import (
	"WordTree/controller"
	docs "WordTree/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	api := r.Group("/api")
	{
		api.GET("/word/search", controller.WordSearch)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":3000")
}
