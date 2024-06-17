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
		c.String(http.StatusOK, "API")
	})

	r.GET("/:value", func(c *gin.Context) {
		var val = c.Param("value")
		c.JSON(http.StatusOK, gin.H{
			"value": val,
		})
	})

	r.POST("/add", func(c *gin.Context) {
		var data TestModel
		if err := c.ShouldBind(&data); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Spintf("%v", err),
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": data,
			})
		}
	})

	r.run("localhost:8080")
}

