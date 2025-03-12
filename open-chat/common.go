package open_chat

import (
	"context"
	"net/http"
	"time"
)

type RequestBuilder[T any] interface {
	Build() (any, error)
}

type ResponseHandler[U any] interface {
	Handle(response *http.Response) (U, error)
}

type ClientInterface[T any, U any] interface {
	Start() error
	SetProxy(proxy string)
	SetParams(params RequestBuilder[T]) error
	SetTimeout(timeout time.Duration)
	OnRequest(f func(ctx context.Context, req *http.Request))
	OnResponse(f func(ctx context.Context, response any))
}
