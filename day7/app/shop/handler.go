package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func goodsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com!",
	})
}
