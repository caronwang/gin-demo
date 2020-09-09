package main

import (
	"fmt"
	"gin-demo/day6/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.LoadBlog(r)
	router.LoadShop(r)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
