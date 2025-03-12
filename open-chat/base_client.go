package open_chat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type BaseClient[T any, U any] struct {
	BaseURL         string
	APIKey          string
	Proxy           string
	Timeout         time.Duration
	Headers         map[string]string
	RequestBuilder  RequestBuilder[T]
	ResponseHandler ResponseHandler[U]
	lock            sync.Mutex
	payload         []byte
	onRequest       []func(context.Context, *http.Request)
	onResponse      []func(context.Context, any)
}

func (c *BaseClient[T, U]) Init(apiKey string, responseHandler ResponseHandler[U]) {
	c.APIKey = apiKey
	c.Timeout = 60 * time.Second
	c.Headers = map[string]string{
		"Authorization": "Bearer " + apiKey,
	}
	c.ResponseHandler = responseHandler
	c.onRequest = []func(context.Context, *http.Request){}
	c.onResponse = []func(context.Context, any){}
}

func (c *BaseClient[T, U]) SetProxy(proxy string) {
	c.Proxy = proxy
}

func (c *BaseClient[T, U]) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}

func (c *BaseClient[T, U]) SetParams(params RequestBuilder[T]) error {
	c.RequestBuilder = params
	payload, err := c.RequestBuilder.Build()
	if err != nil {
		return err
	}

	c.payload, err = json.Marshal(payload)
	if err != nil {
		return err
	}

	fmt.Println("Payload:", string(c.payload))
	return nil
}

func (c *BaseClient[T, U]) OnRequest(f func(ctx context.Context, req *http.Request)) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.onRequest = append(c.onRequest, f)
}

func (c *BaseClient[T, U]) OnResponse(f func(ctx context.Context, response any)) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.onResponse = append(c.onResponse, f)
}

func (c *BaseClient[T, U]) Start() error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	req, err := c.prepareRequest(ctx)
	if err != nil {
		return err
	}

	resp, err := c.executeRequest(req)
	if err != nil {
		return err
	}

	responseStr, err := c.ResponseHandler.Handle(resp)
	if err != nil {
		return err
	}

	c.lock.Lock()
	for _, f := range c.onResponse {
		f(ctx, responseStr)
	}
	c.lock.Unlock()

	return nil
}

func (c *BaseClient[T, U]) prepareRequest(ctx context.Context) (*http.Request, error) {
	urlStr, err := c.constructURL()
	if err != nil {
		return nil, err
	}

	if len(c.payload) == 0 {
		return nil, fmt.Errorf("should have params, using SetParams(params T) to set")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", urlStr, bytes.NewBuffer(c.payload))
	if err != nil {
		return nil, err
	}

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	c.lock.Lock()
	for _, f := range c.onRequest {
		f(ctx, req)
	}
	c.lock.Unlock()

	return req, nil
}

func (c *BaseClient[T, U]) constructURL() (string, error) {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (c *BaseClient[T, U]) executeRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{Timeout: c.Timeout}
	return client.Do(req)
}
