package routes

import (
	"thong-khon-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	thongGroup := router.Group("/thong")
	{
		thongGroup.POST("/", handlers.CreateThong)
		thongGroup.POST("/room", handlers.CreateThongRoom)

		thongGroup.GET("/", handlers.GetAllThong)
		thongGroup.GET("/:id", handlers.GetThongById)
	}

}
