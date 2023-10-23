package repository

import (
	"github.com/sou1991/subject_x_golangs_api_/entity"
	"github.com/gin-gonic/gin"
)

type IAlbumsRepository interface{
	GetAll(...*gin.Context)
}

type AlbumsRepository struct{
	Conn string
}

func NewAlbumRepository() IAlbumsRepository{
	return AlbumsRepository{Conn: "postgres"}
}

func (AlbumsRepository) GetAll(c ...*gin.Context){
	entity.GetAlbums(c[0])
}
