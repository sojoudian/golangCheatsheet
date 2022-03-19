package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Numerics struct {
	Num1   float64 `json:"num1"`
	Num2   float64 `json:"num2"`
	Result float64 `json:"result"`
}

func main() {
	r := gin.Default()

	r.POST("/add", add)
	r.POST("/subtract", subtract)
	r.POST("/multiply", multiply)
	r.POST("/devide", devide)

	r.Run(":8010")
	//fmt.Println("Server is running on port 8010")
}

func add(c *gin.Context) {
	var num Numerics
	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	num.Result = num.Num1 + num.Num2
	c.IndentedJSON(http.StatusOK, num)
}

func subtract(c *gin.Context) {
	var num Numerics
	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	num.Result = num.Num1 - num.Num2
	c.IndentedJSON(http.StatusOK, num)
}

func multiply(c *gin.Context) {
	var num Numerics
	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	num.Result = num.Num1 * num.Num2
	c.IndentedJSON(http.StatusOK, num)
}

func devide(c *gin.Context) {
	var num Numerics
	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	num.Result = num.Num1 / num.Num2
	c.IndentedJSON(http.StatusOK, num)
}

// func add(c *gin.Context){

// }

// func add(c *gin.Context){

// }

// func add(c *gin.Context){

// }
