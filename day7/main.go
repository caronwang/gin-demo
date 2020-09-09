package main

import (
	"fmt"
	"gin-demo/day7/app/blog"
	"gin-demo/day7/app/shop"
	"gin-demo/day7/router"
)

func main() {
	// 加载多个APP的路由配置
	router.Include(shop.Routers, blog.Routers)
	// 初始化路由
	r := router.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
