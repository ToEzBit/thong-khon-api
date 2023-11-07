package main

import (
	"os"
	"thong-khon-api/api/routes"
	"thong-khon-api/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnvVariables()
	config.ConnectDB()

	router := gin.Default()

	routes.SetupRoutes(router)
	router.Run(os.Getenv("PORT"))
}
