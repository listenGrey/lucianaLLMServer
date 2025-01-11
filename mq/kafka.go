package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
	"lucianaLLMServer/model"
)

func PromptQueue(cid int64, qa *model.QA) error {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:  kafka.TCP(os.Getenv("KAFKA_ADDR") + ":" + os.Getenv("KAFKA_PORT")),
		Topic: "sendqa",
		//Balancer:               &kafka.Hash{},
		//WriteTimeout:           1 * time.Second,
		//RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", cid)) // key = cid
	value, err := json.Marshal(qa)        // value = qa
	if err != nil {
		return err
	}

	// 发送消息
	err = writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
