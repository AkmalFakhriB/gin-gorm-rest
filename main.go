package main

import (
	"github.com/akmalfakhrib/gin-gorm-rest/config"
	"github.com/akmalfakhrib/gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.AlbumRoute(router)
	router.Run("127.0.0.1:8080")
}
