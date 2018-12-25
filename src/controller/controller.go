package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// import "encoding/json"
// Blogs json data
type Blogs struct {
	Title      string `Title:"subject"`
	Author     string `Author:""`
	Content    string
	Label      []string `Label:"A Label"`
	Watch      string
	CreateTime int
	UpdateTime int
}

// the blog list
type BlogList []Blog
type Blog struct {
	Title      string
	Author     string
	Sketch     string
	Label      []string
	Watch      string
	CreateTime int
	UpdateTime int
}

// Return to the list of articles
func List(c *gin.Context) {
	// 列表查询与返回
	fmt.Println(c)
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

// Returns the article by ID
func Article(c *gin.Context) {
	c.Request.ParseForm()
	id := c.Param("id")

	c.JSON(200, gin.H{
		"message": id,
	})
}
