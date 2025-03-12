package adapter

type TextPayload struct {
	Model        string    `json:"model"`
	Messages     []Message `json:"messages"`
	EnableSearch bool      `json:"enable_search,omitempty"`
}

type TextBuilder struct {
	Model        string
	SystemPrompt string
	UserPrompt   string
	EnableSearch bool
}

func (b *TextBuilder) Build() (any, error) {
	return TextPayload{
		Model: b.Model,
		Messages: []Message{
			{Role: "system", Content: []Content{{Type: "text", Text: b.SystemPrompt}}},
			{Role: "user", Content: []Content{{Type: "text", Text: b.UserPrompt}}},
		},
		EnableSearch: b.EnableSearch,
	}, nil
}
