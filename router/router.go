package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/subject_x_golangs_api_/controller"
)

type IHandlerAlbumsApi interface{
	AlbumsApiCall()
}

type routerAlubum struct{
	Iac controller.IAlbumsController
}

func NewRouter(iac controller.IAlbumsController, c *gin.Context) IHandlerAlbumsApi{
	return &routerAlubum{Iac: iac, Ctx: c}
}

func (ra routerAlubum) AlbumsApiCall(r *gin.*Engine){
	r.GET("/albums", func(c *gin.Context) {
		ra.Iac.GetAll()
	})
	//r.GET("/albums/:id", model.GetAlbumByID)
	//r.POST("/albums", model.PostAlbums)
}
