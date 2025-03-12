package adapter

import (
	"io"
	"net/http"
)

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content,omitempty"`
}

type Content struct {
	Type     string       `json:"type"`
	Text     string       `json:"text,omitempty"`
	ImageURL ContentImage `json:"image_url,omitempty"`
}

type ContentImage struct {
	URL string `json:"url"`
}

type PayloadBuilder interface {
	Build() (any, error)
}

type ResponseHandler interface {
	Handle(response *http.Response) (string, error)
}

type DefaultResponseHandler struct{}

func (h *DefaultResponseHandler) Handle(response *http.Response) (string, error) {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
