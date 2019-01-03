package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

// import "encoding/json"
// Blogs json data
type Blogs struct {
	Title      string `Title:"subject"`
	Author     string `Author:""`
	Content    string
	Label      []string `Label:"A Label"`
	Watch      string   `Watch:"Number of views"`
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

func Essay(c *gin.Context) {
	fmt.Println(c.Request.Body)

	con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	fmt.Println(string(con))
	fmt.Println("---------------")
	fmt.Println(con)
	// 创建连接
	var config pgx.ConnConfig

	// config.Host = "postgraSql"
	config.Host = "localhost"
	config.User = "postgres"
	config.Password = "postgres"
	config.Database = "postgres"

	conn, err := pgx.Connect(config)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}

	tx, err := conn.Exec(`INSERT INTO "public"."article" (
        "title",
        "context",
        "label",
        "create_time",
        "update_time" 
    )
    VALUES
        (
            'fdasfdsfdsafdsafdsaf',
            '这是英文',
            '{"hello":"world"}',
            '2019-01-02 19:31:16',
        '2019-01-02 19:31:21' 
        )`)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}

	num := tx.RowsAffected()
	if num >= 1 {
		c.JSON(http.StatusOK, gin.H{
			"message": "数据插入成功",
		})
	}
}
