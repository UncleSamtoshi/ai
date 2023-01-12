package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	c := gogpt.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx := context.Background()

	//get all the text after the command as the text
	text := strings.Join(os.Args[1:], " ")

	// Create a new request
	req := gogpt.CompletionRequest{
		Model:     "text-davinci-003",
		MaxTokens: 1000,
		Prompt:    text,
	}

	// Send request to OpenAI
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(resp.Choices[0].Text + "\n")
}
