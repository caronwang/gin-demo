# 简介
上传文件

```go
r.POST("/upload", func(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.String(500, "上传图片出错")
    }
    // c.JSON(200, gin.H{"message": file.Header.Context})
    c.SaveUploadedFile(file, "uploadFile/"+file.Filename)
    c.String(http.StatusOK, file.Filename)
})
```