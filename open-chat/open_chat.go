package open_chat

import "open-chat/adapter"

type Client[T any] struct {
	BaseClient[T, string]
}

func NewClient[T any, U any](baseUrl string, apiKey string) ClientInterface[T, U] {
	c := &Client[T]{}
	c.BaseURL = baseUrl
	c.Init(apiKey, &adapter.DefaultResponseHandler{})
	return c
}

func NewTextClient(url, key string) ClientInterface[adapter.TextBuilder, adapter.TextBuilder] {
	return NewClient[adapter.TextBuilder, adapter.TextBuilder](url, key)
}
