package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	chat "github.com/pedrobertao/go-chatgpt/chat"
)

func main() {
	godotenv.Load()

	fmt.Println("go-chatgpt-sdk say hello!")

	chat := chat.ChatGPT{
		BaseURI: chat.API_V1_URL,
		ApiKey:  os.Getenv("API_KEY"),
		Model:   chat.GPT_35_TURBO,
	}

	// Example Test
	res, err := chat.SayHello()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)

	fmt.Println("=================")
	// Example Promtp
	res, err = chat.SendPrompt("Tell me the story of golang with 200 characters")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)

	os.Exit(0)
}
