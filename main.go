package main

import (
	"context"
	"fmt"
	"log"
	open_chat "open-chat"
	"open-chat/adapter"
	"open-chat/config"
)

func main() {
	conf, err := config.LoadConfig[config.AppConfig]("./config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	client := open_chat.NewTextClient(conf.API.APIUrl, conf.API.APIKey)

	client.SetParams(&adapter.TextBuilder{
		Model:        "qwen-max-2025-01-25",
		SystemPrompt: "我是小安, 让我协助你编写前端代码吧!请描述你的需求, 我将会为你编写tailwindcss风格的代码.",
		UserPrompt:   `基于如下:<div class="p-4 bg-blue-500 text-white">Hello World</div>编写一个影视页面。`,
	})

	// client.SetParams(&adapter.VisionBuilder{
	// 	Model:      "qwen-vl-plus",
	// 	UserPrompt: "这是什么?",
	// 	ImageURL:   "https://dashscope.oss-cn-beijing.aliyuncs.com/images/dog_and_girl.jpeg",
	// })

	client.OnResponse(func(ctx context.Context, response any) {
		log.Println(response.(string))
	})

	if err := client.Start(); err != nil {
		log.Println(err)
	}
}
