package adapter

import "encoding/json"

type DocumentPayload struct {
	Model      string        `json:"model"`
	Input      DocumentInput `json:"input"`
	Parameters struct {
		ResultFormat string `json:"result_format"`
	} `json:"parameters"`
}

type DocumentInput struct {
	Messages []Message `json:"messages"`
}

type DocumentBuilder struct {
	Model      string
	UserPrompt string
	DocumentID string
}

func (b *DocumentBuilder) Build() (any, error) {
	documentInfo, err := json.Marshal(map[string]string{
		"content":   "百炼X1 搭载6.7英寸1440 x 3200像素超清屏幕...",
		"file_type": "docx",
		"filename":  "百炼系列手机产品介绍",
		"title":     "百炼手机产品介绍",
	})
	if err != nil {
		return nil, err
	}

	return DocumentPayload{
		Model: b.Model,
		Input: DocumentInput{
			Messages: []Message{
				{Role: "system", Content: []Content{{Type: "text", Text: string(documentInfo)}}},
				{Role: "user", Content: []Content{{Type: "text", Text: b.UserPrompt}}},
			},
		},
		Parameters: struct {
			ResultFormat string `json:"result_format"`
		}{ResultFormat: "message"},
	}, nil
}
