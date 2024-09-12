package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getOneAlbums(c *gin.Context) {
	var albumOne = album{ID: "123", Title: "yuseo and the gang", Artist: "yuseo", Price: 12.5}
	c.IndentedJSON(http.StatusOK, albumOne)
}
