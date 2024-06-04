package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define the album struct and albums slice
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// GetAlbums returns the list of albums

func GetAlbums(c *gin.Context) {
	// champion := c.Query("championName")
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		c.Error(err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		c.Error(errors.New("failed to fetch albums"))
		return
	}

	var raw json.RawMessage
	err = json.NewDecoder(res.Body).Decode(&raw)
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, raw)
}

// PostAlbums adds a new album to the list
func PostAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID returns the album with the specified ID
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// DeleteAlbumByID deletes the album with the specified ID
func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			// Delete the album from the slice
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func UpdateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	for i, a := range albums {
		if a.ID == id {
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
