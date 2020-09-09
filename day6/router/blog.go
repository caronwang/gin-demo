package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func commentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com!",
	})
}

func LoadBlog(e *gin.Engine) {
	e.GET("/comment", commentHandler)
}
