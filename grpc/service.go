package grpc

import (
	"context"
	"github.com/tmc/langchaingo/llms/ollama"
	"google.golang.org/grpc/peer"
	"log"
)

type ResponseService struct {

}

func (s *ResponseService) Response(ctx context.Context) () {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("生成回答")
	}

	llm, err := ollama.New(ollama.WithModel("llama3"))
	if err != nil {
		return "", err
	}

	res, err := llm.Call(context.Background(), //request.Prompt)
	if err != nil {
		return "", err
	}
}