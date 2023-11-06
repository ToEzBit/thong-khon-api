package main

import (
	"thong-khon-api/config"
)

func main() {
	config.LoadEnvVariables()
	config.ConnectDB()
}
