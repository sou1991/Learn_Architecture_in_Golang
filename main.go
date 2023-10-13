package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/subject_x_golangs_api_/router"
	"github.com/sou1991/subject_x_golangs_api_/controller"
	"github.com/sou1991/subject_x_golangs_api_/repository"
)

var repo IAlbumsRepository = repository.NewAlbumRepository()
var ctrl IAlbumsController = controller.NewAlbumController(repo)
var ro IHandlerAlbumsApi = router.NewRouter(ctrl)

func main() {
	r := gin.Default()
	ro.AlbumsApiCall(&r)
	r.Run("0.0.0.0:80")
}
