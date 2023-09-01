package model

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type albam struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price int `json:"price"`
}

//albums slice to seed record album data.
var albams = []albam{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 5000},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 2000},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 3000},
}

func GetAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albams)
}
