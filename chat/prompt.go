package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Body  ResponseBody
	Code  int
	Error error
}

type ChatGPT struct {
	BaseURI string
	ApiKey  string
	Model   string
}

func (c *ChatGPT) SayHello() (*Response, error) {
	url := fmt.Sprintf("%s%s", API_V1_URL, "/chat/completions")
	data := map[string]interface{}{
		"model": c.Model,
		"messages": []map[string]string{
			{"role": "user", "content": "Hello, how are you?"},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return &Response{}, nil
	}

	// Create a new POST request with the JSON data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return &Response{}, nil
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)
	// Perform the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return &Response{}, nil
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return &Response{}, nil
	}

	var responseBody ResponseBody

	if err := json.Unmarshal(body, &responseBody); err != nil {
		return &Response{}, err
	}
	return &Response{
		Body:  responseBody,
		Code:  response.StatusCode,
		Error: nil,
	}, nil
}
