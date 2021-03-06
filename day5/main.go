package main

import (
	"fmt"
	"gin-demo/day5/router"
)

func main() {
	r := router.SetupRouter()
	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
