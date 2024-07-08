package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id    int    `form:"id"`
	Name  string `form:"name"`
	Email string `form:"email"`
}

func main() {
	r := gin.Default()
	r.POST("/user", func(c *gin.Context) {
		if err := c.ShouldBindQuery(&userObj); err == nil {
			fmt.Printf("user obj - %+v \n", userObj)
		} else {
			fmt.Printf("error - %+v \n", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userObj,
		})
	})
	r.Run("localhost:8080")
}
