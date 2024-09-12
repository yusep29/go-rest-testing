package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	initDb()

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album-one", getOneAlbums)
	router.GET("/cat", getCatFact)
	router.GET("/user", getUser)
	router.POST("/user/add", addUser)
	router.DELETE("/user/delete", deleteUser)

	router.Run(":8081")
}

//source : https://go.dev/doc/tutorial/web-service-gin
