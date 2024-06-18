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
	r.GET("/user", func(c *gin.Context) {
		var userObj user
		if err := c.ShouldBindQuery(&userObj); err == nil {
			fmt.Printf("user obj - %+v \n", userObj)
		}
	})
	r.Run("localhost:8080")
}
