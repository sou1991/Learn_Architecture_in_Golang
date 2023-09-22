package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type albam struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  int    `json:"price"`
}

type response struct {
	Albams []albam `json:"albams"`
}

// albums slice to seed record album data.
var albams = response{
	Albams: []albam{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 5000},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 2000},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 3000},
	},
}

func GetAlbumByID(c *gin.Context){
	id := c.Param("id")

	for _, a := range albams.Albams {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albams)
}

func PostAlbums(c *gin.Context) {
	var newAlbums albam

	if err := c.BindJSON(&newAlbums); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	// Add the new album to the slice
	albams.Albams = append(albams.Albams, newAlbums)
	c.IndentedJSON(http.StatusCreated, newAlbums)
}
