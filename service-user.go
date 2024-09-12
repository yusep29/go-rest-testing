package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/lib/pq"
)

func getUser(c *gin.Context) {
	arrUser := dbGetUser()
	c.IndentedJSON(http.StatusOK, arrUser)
}

func addUser(c *gin.Context) {
	request := User{}
	if errA := c.ShouldBindBodyWith(&request, binding.JSON); errA == nil {
		dbAddUser(&request)
		c.IndentedJSON(http.StatusOK, "ok")
	}
}

func deleteUser(c *gin.Context) {
	request := User{}
	if errA := c.ShouldBindBodyWith(&request, binding.JSON); errA == nil {
		dbDeleteUser(request.Id)
		c.IndentedJSON(http.StatusOK, "ok")
	}

}
