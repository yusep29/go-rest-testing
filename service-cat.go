package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCatFact(c *gin.Context) {
	url := "https://catfact.ninja/fact"

	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	c.IndentedJSON(http.StatusOK, string(body))
}
