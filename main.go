package main

import (
	"fmt"
	"lucianaLLMServer/controller"
)

func main() {
	fmt.Println("正在运行")
	errCh := make(chan error)

	go controller.Response(errCh)

	for {
		select {
		case err := <-errCh:
			fmt.Printf("生成对话服务挂掉了, %s\n", err)
		}
	}
}
