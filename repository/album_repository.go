package repository

import (
	"github.com/sou1991/subject_x_golangs_api_/entity"
	"github.com/gin-gonic/gin"
)

type IAlbumsRepository interface{
	GetAll()
}

type AlbumsRepository struct{
	Conn string
}

func NewAlbumRepository(c *gin.Context) IAlbumsRepository{
	return AlbumsRepository{Conn: "postgres", Ctx: c}
}

func (AlbumsRepository) GetAll(){
	entity.GetAlbums()
}
