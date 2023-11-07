package main

import (
	"thong-khon-api/config"
	"thong-khon-api/models"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()

}

func main() {
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Thong{})
	config.DB.AutoMigrate(&models.ThongRoom{})
	config.DB.AutoMigrate(&models.Reserve{})
}
