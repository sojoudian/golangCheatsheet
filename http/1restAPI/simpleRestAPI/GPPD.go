package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	//
)

type Product struct {
	Id    int     `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Stock int     `json:"stock"`
	Price float64 `json:"price"`
}

func main() {
	r := gin.Default()
	// HHTP GET, POST, PUT, DELETE
	r.GET("/", getFunc)
	r.POST("/", postFunc)
	r.PUT("/", putFunc)
	r.DELETE("/", deleteFunc)

	r1 := r.Group("/api")
	{
		r1.GET("/v1", getFunc) // http://<server>/api/v1
		r1.POST("/v1", postFunc)
		r1.PUT("/v1", putFunc)
		r1.DELETE("/v1", deleteFunc)
	}

	// handle path param
	r.GET("/productz/:id", getProductBYID)   // http://<server>/
	r.GET("/profile/:username", showProfile) // http://<server>/
	r.GET("/compute/:num1/add/:num2", compute)

	// handling query string params
	// /employee?firstname=zzzz&last=cccc&id=4510041
	r.GET("/employee", showEmployee) // http://<server>//employee?firstname=zzzz&last=cccc&id=4510041

	// Binding posd data
	r.POST("/product", performProduct)   // http://<server>
	r.POST("/products", performProducts) // http://<server>

	// read from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loadinf .env file", err)
	}
	PORT := os.Getenv("GINPORT")

	r.Run(PORT)

}

func performProduct(c *gin.Context) {
	var product Product
	if err := c.BindJSON(&product); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalide request",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, product)
}

func performProducts(c *gin.Context) {
	var products []Product
	if err := c.BindJSON(&products); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalide request",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, products)
}

func showEmployee(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "")
	lastname := c.DefaultQuery("lastname", "")
	id, _ := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 0)
	c.IndentedJSON(http.StatusOK, gin.H{
		"id":        id,
		"firstname": firstname,
		"las	tname": lastname,
	})
}

func getProductBYID(c *gin.Context) {
	id := c.Param("id") // string
	idn, _ := strconv.ParseInt(id, 10, 0)
	c.IndentedJSON(http.StatusOK, gin.H{
		"id":   idn,
		"name": "Product A",
	})
}

func showProfile(c *gin.Context) {
	username := c.Param("username") //string

	c.IndentedJSON(http.StatusOK, gin.H{
		"username": username,
	})
}

func compute(c *gin.Context) {
	num1, _ := strconv.ParseInt(c.Param("num1"), 10, 0)
	num2, _ := strconv.ParseInt(c.Param("num2"), 10, 0)
	result := num1 + num2
	c.IndentedJSON(http.StatusOK, gin.H{
		"num1":   num1,
		"num2":   num2,
		"result": result,
	})
}

func getFunc(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello RESP API - HTTP GET"})
}

func postFunc(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello RESP API - HTTP POST"})
}

func putFunc(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello RESP API - HTTP PUT"})
}

func deleteFunc(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello RESP API - HTTP DELETE"})
}
