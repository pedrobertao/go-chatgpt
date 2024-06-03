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

	res, err := chat.SayHello()
	if err != nil {
		fmt.Printf("Ops..: %v \n", err)
		os.Exit(1)
	}

	fmt.Println(res)
	os.Exit(0)
}
