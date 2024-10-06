package main

import (
	"fmt"
	"lucianaLLMServer/controller"
)

func main() {
	fmt.Println("对话服务正在运行")
	err := controller.Response()
	if err != nil {
		fmt.Printf("生成对话服务挂掉了, %s\n", err)
	}

}
