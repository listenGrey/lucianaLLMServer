package controller

import (
	"fmt"
	"lucianaLLMServer/logic"
)

func Response(errCh chan<- error) {
	fmt.Println("回答服务正在运行")
	for {
		if err := logic.Response(); err != nil {
			errCh <- err
			return
		}
	}
}
