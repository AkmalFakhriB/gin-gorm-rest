package routes

import (
	"github.com/akmalfakhrib/gin-gorm-rest/controller"
	"github.com/gin-gonic/gin"
)

func AlbumRoute(router *gin.Engine) {
	router.GET("/album", controller.GetAlbum)
	router.POST("/album", controller.PostAlbum)
	router.GET("/album/:id", controller.GetAlbumById)
	router.PUT("/album/:id", controller.UpdateAlbum)
	router.DELETE("/album/:id", controller.DeleteAlbum)
}
