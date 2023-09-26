package main

import (
	"context"
	"errors"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"io"
)

//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

func main() {
	c := openai.NewClient("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")
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 20,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "重复我刚才说的话",
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}
}
