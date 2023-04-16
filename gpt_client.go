package main

import (
	"cle/constants"
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"os"
)

func printResultWithGptCompletion(text string) error {
	AGptCli := openai.NewClientWithConfig(newClient())
	text = "简单解释这句命令行的作用: " + text
	message := make([]openai.ChatCompletionMessage, 0)
	message = append(message, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: text,
	})

	stream, err := AGptCli.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: message,
		})
	if err != nil {
		fmt.Printf("ChatCompletion error: %v", err)
		return err
	}
	defer stream.Close()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println("ChatCompletionStream failed, err: ", err.Error())
			return err
		}
		fmt.Print(response.Choices[0].Delta.Content)
	}
	return nil
}

func newClient() openai.ClientConfig {
	config := openai.DefaultConfig(os.Getenv(constants.ApiKeyEnvKey))
	return config
}
