package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/subject_x_golangs_api_/model"
)

func main() {
	router := gin.Default()

	router.GET("/albums", model.GetAlbums)
	router.POST("/albums", model.PostAlbums)
	router.Run("0.0.0.0:80")
}
