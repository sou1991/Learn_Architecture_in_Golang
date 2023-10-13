package controller

import (
	"github.com/sou1991/subject_x_golangs_api_/repository"
	"github.com/gin-gonic/gin"
)

type IAlbumsController interface {
	//Create()
	GetAll()
	//GetByID()
}

func NewAlbumController(ar repository.IAlbumsRepository, c *gin.Context) IAlbumsController{
	return AlbumController{Ar: ar, Ctx: c}
}

type AlbumController struct{
	Ar repository.IAlbumsRepository
}

func (ac AlbumController) GetAll(){
	ac.Ar.GetAll()
}
