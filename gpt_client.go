package main

import (
	"cle/constants"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/Xuanwo/go-locale"
	"github.com/sashabaranov/go-openai"
)

func printResultWithGptCompletion(text string) error {
	AGptCli := openai.NewClientWithConfig(newClient())
	prefix := constants.QuestionPrefix
	tag, _ := locale.Detect()

	text = fmt.Sprintf(prefix, runtime.GOOS, tag.String(), text)
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
