package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		fmt.Println(action)
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
	//默认为监听8080端口

	r.GET("/prod", func(c *gin.Context) {
		//指定默认值
		//http://localhost:8080/prod才会打印出来默认的值
		name := c.DefaultQuery("name", "tv")
		c.String(http.StatusOK, fmt.Sprintf("product name = %s", name))
	})

	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("name")
		password := c.PostForm("pwd")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})

	r.Run(":8000")
}
