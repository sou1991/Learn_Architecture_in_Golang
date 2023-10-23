package controller

import (
	"github.com/sou1991/subject_x_golangs_api_/repository"
	"github.com/gin-gonic/gin"
)

type IAlbumsController interface {
	//Create()
	GetAll(...*gin.Context)
	//GetByID()
}

func NewAlbumController(ar repository.IAlbumsRepository) IAlbumsController{
	return AlbumController{Ar: ar}
}

type AlbumController struct{
	Ar repository.IAlbumsRepository
}

func (ac AlbumController) GetAll(c ...*gin.Context){
	ac.Ar.GetAll(c[0])
}
