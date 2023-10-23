package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/subject_x_golangs_api_/controller"
)

type IHandlerAlbumsApi interface{
	AlbumsApiCall(...*gin.Engine)
}

type routerAlubum struct{
	Iac controller.IAlbumsController
}

func NewRouter(iac controller.IAlbumsController) IHandlerAlbumsApi{
	return &routerAlubum{Iac: iac}
}

func (ra routerAlubum) AlbumsApiCall(r ...*gin.Engine){
	r[0].GET("/albums", func(c *gin.Context) {
		ra.Iac.GetAll(c)
	})
	//r.GET("/albums/:id", model.GetAlbumByID)
	//r.POST("/albums", model.PostAlbums)
}
