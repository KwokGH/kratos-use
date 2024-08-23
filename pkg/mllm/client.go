package mllm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	clientPool sync.Pool

	ErrInternalServerError = errors.New("InternalServerError")
)

func init() {
	clientPool = sync.Pool{
		New: func() interface{} {
			return &http.Client{
				Transport: &http.Transport{
					MaxIdleConns:        100,
					MaxIdleConnsPerHost: 100,
					IdleConnTimeout:     30 * time.Minute,
				},
			}
		},
	}
}
func getClient() *http.Client {
	return clientPool.Get().(*http.Client)
}

func putClient(client *http.Client) {
	clientPool.Put(client)
}

type Client struct {
	apiKey  string
	BaseUrl string
	//HttpClient     *http.Client
	requestBuilder *HTTPRequestBuilder
}

func NewClient(apiKey, baseURL string) *Client {
	return &Client{
		apiKey:  apiKey,
		BaseUrl: baseURL,
		//HttpClient:     &http.Client{},
		requestBuilder: NewRequestBuilder(),
	}
}

type Response struct {
	Code int    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"varint,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

type requestOptions struct {
	body   any
	header http.Header
}
type requestOption func(*requestOptions)

func withBody(body any) requestOption {
	return func(args *requestOptions) {
		args.body = body
	}
}

func withContentType(contentType string) requestOption {
	return func(args *requestOptions) {
		args.header.Set("Content-Type", contentType)
	}
}

func (c *Client) newRequest(ctx context.Context, method, url string, setters ...requestOption) (*http.Request, error) {
	// Default Options
	args := &requestOptions{
		body:   nil,
		header: make(http.Header),
	}
	for _, setter := range setters {
		setter(args)
	}
	req, err := c.requestBuilder.Build(ctx, method, url, args.body, args.header)
	if err != nil {
		return nil, err
	}
	c.setCommonHeaders(req)
	return req, nil
}
func (c *Client) sendRequest(ctx context.Context, req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json")

	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := getClient()
	defer putClient(client)

	res, err := client.Do(req)

	if err != nil {
		log.Context(ctx).Errorw("msg", "请求失败", "err", err)
		return err
	}

	defer res.Body.Close()

	if isFailureStatusCode(res) {
		log.Context(ctx).Errorw("msg", "错误的状态", "res", res)
		return ErrInternalServerError
	}

	return decodeResponse(res.Body, v)
}

func decodeResponse(body io.Reader, v any) error {
	if v == nil {
		return nil
	}

	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func (c *Client) setCommonHeaders(req *http.Request) {
	if c.apiKey != "" {
		// OpenAI or Azure AD authentication
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}
}

func (c *Client) fullURL(suffix string) string {
	return c.BaseUrl + suffix
}

type streamable interface {
	ChatCompletionStreamResponse
}
type ChatCompletionStreamResponse struct {
}

type JSONUnmarshaler struct{}

func (jm *JSONUnmarshaler) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

type HTTPRequestBuilder struct {
	marshaller Marshaller
}

type Marshaller interface {
	Marshal(value any) ([]byte, error)
}

type JSONMarshaller struct{}

func (jm *JSONMarshaller) Marshal(value any) ([]byte, error) {
	return json.Marshal(value)
}

func NewRequestBuilder() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{
		marshaller: &JSONMarshaller{},
	}
}

func (b *HTTPRequestBuilder) Build(
	ctx context.Context,
	method string,
	url string,
	body any,
	header http.Header,
) (req *http.Request, err error) {
	var bodyReader io.Reader
	if body != nil {
		if v, ok := body.(io.Reader); ok {
			bodyReader = v
		} else {
			var reqBytes []byte
			reqBytes, err = b.marshaller.Marshal(body)
			if err != nil {
				return
			}
			bodyReader = bytes.NewBuffer(reqBytes)
		}
	}
	req, err = http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return
	}
	if header != nil {
		req.Header = header
	}
	return
}

func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}
