package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"config"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

func connInit(c *gin.Context) *pgx.Conn {
	var connInfo pgx.ConnConfig
	connInfo.Host = config.Instance.Db.Host
	connInfo.Port = config.Instance.Db.Port
	connInfo.Database = config.Instance.Db.Database
	connInfo.User = config.Instance.Db.Username
	connInfo.Password = config.Instance.Db.Password

	conn, err := pgx.Connect(connInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
	return conn
}

// import "encoding/json"
// Blogs json data
type Blogs struct {
	Title      string `Title:"subject"`
	Author     string `Author:""`
	Content    string
	Label      []string `Label:"A Label, "`
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

	result, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	var f *Blogs
	err := json.Unmarshal(result, &f)

	nowTime := time.Now().Format(config.Instance.Time)
	// 连接
	conn := connInit(c)
	// 执行语句
	tx, err := conn.Exec(`INSERT INTO "public"."article" (
        "title", "context", "label", "create_time", "update_time"  )
    VALUES
    ($1,$2,'{"hello":"world"}', $3, $4)`,
		f.Title, f.Content, nowTime, nowTime)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
	// 返回受影响的行数，否则返回0
	num := tx.RowsAffected()
	if num >= 1 {
		c.JSON(http.StatusOK, gin.H{
			"message": "数据插入成功",
		})
	}
}
