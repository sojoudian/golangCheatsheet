package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hi from RESTful API \n")
	})

	r.GET("/v2", response)
	r.Run(":8081")
	fmt.Println("Server is running")

}

func response(c *gin.Context) {
	c.String(http.StatusOK, "Hi from version 2 RESTful API \n")
}
