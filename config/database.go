package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := os.Getenv("DB_URL")

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

}

func CloseDB() {
	db, err := DB.DB()
	if err != nil {
		panic("Failed to get the underlying DB connection: " + err.Error())
	}
	db.Close()
}
