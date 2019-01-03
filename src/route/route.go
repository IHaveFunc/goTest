package route

import (
	"controller"

	"github.com/gin-gonic/gin"
)

// func GetDataB(c *gin.Context) {
//     var b StructB
//     c.Bind(&b)
//     c.JSON(200, gin.H{
//         "a": b.NestedStruct,
//         "b": b.FieldB,
//     })
// }
func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()

	route.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	route.GET("/list", controller.List)
	route.GET("/article/:id", controller.Article)

	route.POST("/save/article", controller.Essay)

	return route
}
