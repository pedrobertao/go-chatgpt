package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	chat "github.com/pedrobertao/go-chatgpt/chat"
)

func main() {
	godotenv.Load()

	fmt.Println("go-chatgpt-sdk SAY HELLO !")
	chat := chat.ChatGPT{
		BaseURI: chat.API_V1_URL,
		ApiKey:  os.Getenv("API_KEY"),
		Model:   "gpt-3.5-turbo-0125",
	}
	res, err := chat.SayHello()
	if err != nil {
		fmt.Printf("BOOOM: %v \n", err)
		os.Exit(1)
	}

	if res.Code == 200 {
		fmt.Printf(res.Body.Choices[0].Message.Content + "\n")
		//Hello! I'm just a computer program, so I don't have feelings,
		//but I'm here to help you. How can I assist you today?
	}

	os.Exit(0)
}
