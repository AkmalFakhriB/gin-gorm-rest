package controller

import (
	"fmt"
	"net/http"

	"github.com/akmalfakhrib/gin-gorm-rest/config"
	"github.com/akmalfakhrib/gin-gorm-rest/models"
	"github.com/gin-gonic/gin"
)

func GetAlbum(c *gin.Context) {
	albums := []models.Album{}
	config.DB.Find(&albums)
	c.JSON(http.StatusOK, &albums)
}

func PostAlbum(c *gin.Context) {
	var albums models.Album
	c.BindJSON(&albums)
	result := config.DB.Create(&albums)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	c.JSON(http.StatusCreated, albums)

}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	albums := []models.Album{}
	config.DB.First(&albums, id)
	c.JSON(200, &albums)
}

func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var albums models.Album
	config.DB.First(&albums, id)
	c.BindJSON(&albums)
	result := config.DB.Save(&albums)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	c.JSON(200, &albums)
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	var albums models.Album
	config.DB.Where("id = ?", id).Delete(&albums)
	c.JSON(200, &albums)
}
