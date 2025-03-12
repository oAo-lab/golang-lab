package adapter

type VisionPayload struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type VisionBuilder struct {
	Model      string
	UserPrompt string
	ImageURL   string
}

func (b *VisionBuilder) Build() (any, error) {
	return VisionPayload{
		Model: b.Model,
		Messages: []Message{
			{
				Role: "user",
				Content: []Content{
					{Type: "text", Text: b.UserPrompt},
					{Type: "image_url", ImageURL: ContentImage{URL: b.ImageURL}},
				},
			},
		},
	}, nil
}
