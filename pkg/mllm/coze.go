package mllm

import (
	"context"
	"net/http"
)

type CozeReq struct {
	BotId          string `json:"bot_id"`
	ConversationId string `json:"conversation_id"`
	User           string `json:"user"`
	Query          string `json:"query"`
	Stream         bool   `json:"stream"`
}

type CozeResp struct {
	ConversationId string         `json:"conversation_id"`
	Messages       []*CozeMessage `json:"messages"`
	Code           int            `json:"code"`
	Msg            string         `json:"msg"`
}
type CozeMessage struct {
	Role        string `json:"role"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
}

type CozeMLLMClient struct {
	client  *Client
	key     string
	baseUrl string
}

func NewCozeClient(ctx context.Context, key, baseUrl string) *CozeMLLMClient {
	// os.Getenv("LINGMU_COZE_KEY")
	// "https://api.coze.cn/open_api/v2"
	client := NewClient(key, baseUrl)
	return &CozeMLLMClient{
		client:  client,
		key:     key,
		baseUrl: baseUrl,
	}
}

func (c *CozeMLLMClient) SendRequest(ctx context.Context, req *CozeReq) (*CozeResp, error) {
	request, err := c.client.newRequest(ctx, http.MethodPost, c.client.fullURL("/chat"), withBody(req))
	if err != nil {
		return nil, err
	}

	resp := &CozeResp{}
	err = c.client.sendRequest(ctx, request, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
