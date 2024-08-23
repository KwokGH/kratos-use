package mllm

import (
	"context"
	"os"
	"testing"
)

func TestRequestKouZi(t *testing.T) {
	ctx := context.Background()
	key := os.Getenv("LINGMU_COZE_KEY")
	t.Log(key)
	c := NewCozeClient(ctx, key, "https://api.coze.cn/open_api/v2")

	request := &CozeReq{
		BotId:          "7380211195858665524",
		ConversationId: "1",
		User:           "test",
		Query:          "请根据以下信息，给我讲一个故事：\n主题： 中国历史故事；\n相关内容：秦始皇\n字数：800字",
		Stream:         false,
	}
	resp, err := c.SendRequest(ctx, request)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range resp.Messages {
		t.Log(item.Type)
		t.Log(item.ContentType)
		t.Log(item.Content)
	}
}
