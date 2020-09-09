package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传文件出错")
			return
		}

		err = c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			c.String(500, "保存文件"+file.Filename+"出错")
			return
		}
		c.String(http.StatusOK, file.Filename)
	})

	r.POST("/uploadEx", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}
		//headers.Size 获取文件大小
		if headers.Size > 1024*1024*2 {
			fmt.Println("文件太大了")
			return
		}
		//headers.Header.Get("Content-Type")获取上传文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			c.String(http.StatusForbidden, "只允许上传png图片")
			return
		}
		c.SaveUploadedFile(headers, headers.Filename)
		c.String(http.StatusOK, headers.Filename)
	})
	r.Run(":8000")

}
