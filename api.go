package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestModel struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"responseData": "API data"
		})
		// c.String(http.StatusOK, "API")
	})

	r.GET("/:name", func(c *gin.Context) {
		var val = c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"value": val,
		})
	})

	r.POST("/add", func(c *gin.Context) {
		var data TestModel
		if err := c.ShouldBind(&data); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": data,
			})
		}
	})

	r.Run("localhost:8080")
}
