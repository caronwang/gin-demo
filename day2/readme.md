# 说明
一个简单的API接口，以Url中的信息为参数

```go
r.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    //截取/
    action = strings.Trim(action, "/")
    c.String(http.StatusOK, name+" is "+action)
})
```
测试接口
```shell script
GET http://127.0.0.1:8000/user/wang/say

wang is say
```

GET获取参数

```go
r.GET("/prod", func(c *gin.Context) {
    name := c.DefaultQuery("name", "tv")
    c.String(http.StatusOK, fmt.Sprintf("product name = %s", name))
})
```


POST获取参数
```go

r.POST("/form", func(c *gin.Context) {
    types := c.DefaultPostForm("type", "post")
    username := c.PostForm("username")
    password := c.PostForm("userpassword")
    // c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
    c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
})

```