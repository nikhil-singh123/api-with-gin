package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue", Artist: "Honey", Price: 57.99},
	{ID: "2", Title: "Red", Artist: "Singh", Price: 59.99},
	{ID: "3", Title: "Green", Artist: "Arjit", Price: 654.99},
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	newAlbum.ID = uuid.New().String()

	for _, a := range albums {
		if a.ID == newAlbum.ID {
			newAlbum.ID = uuid.New().String()
			return
		}
	}
	albums = append(albums, newAlbum)

	c.JSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
}

func patchAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			var updatedAlbum album

			if err := c.BindJSON(&updatedAlbum); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if updatedAlbum.Title != "" {
				albums[i].Title = updatedAlbum.Title
			}

			if updatedAlbum.Artist != "" {
				albums[i].Artist = updatedAlbum.Artist
			}

			if updatedAlbum.Price != 0 {
				albums[i].Price = updatedAlbum.Price
			}

			c.JSON(http.StatusOK, albums[i])
		}
	}
}

func main() {
	r := gin.Default()

	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbumById)
	r.POST("/albums", postAlbum)
	r.PATCH("albums/:id", patchAlbumByID)
	r.Run(":8040")

}
