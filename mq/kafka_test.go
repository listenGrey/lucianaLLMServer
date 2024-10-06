package mq

import (
	"lucianaLLMServer/model"
	"testing"
)

func TestPromptQueue(t *testing.T) {
	var cid int64 = 98765
	qa := &model.QA{
		Request:  "test question",
		Response: "test answer",
	}
	err := PromptQueue(cid, qa)
	if err != nil {
		t.Error(err)
	}
}
