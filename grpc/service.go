package grpc

import (
	"context"
	"github.com/listenGrey/lucianagRpcPKG/ask"
	"github.com/tmc/langchaingo/llms/ollama"
	"google.golang.org/grpc/peer"
	"log"
)

type RequestServer struct {
	ask.UnimplementedRequestServer
}

func (r *RequestServer) Request(ctx context.Context, re *ask.Prompt) (*ask.Response, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("生成回答")
	}

	llm, err := ollama.New(ollama.WithModel("llama3"))
	if err != nil {
		return nil, err
	}

	res, err := llm.Call(context.Background(), re.GetPrompt())
	if err != nil {
		return nil, err
	}

	return &ask.Response{Response: res}, nil
}
