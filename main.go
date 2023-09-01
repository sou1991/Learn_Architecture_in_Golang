package main

import (
	"github.com/sou1991/subject_x_golangs_api_/model"
	"github.com/gin-gonic/gin"
)
func main(){
	router := gin.Default()

	router.GET("/albums", model.GetAlbums)
	router.Run("0.0.0.0:80")
}
