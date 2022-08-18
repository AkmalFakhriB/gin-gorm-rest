package config

import (
	"github.com/akmalfakhrib/gin-gorm-rest/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "golang:golang@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Album{})
	DB = db
}
