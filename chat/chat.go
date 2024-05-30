package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ChatGPT struct {
	BaseURI URLVersion
	ApiKey  string
	Model   Models
}

func (c ChatGPT) check() error {
	if c.ApiKey == "" || c.BaseURI == "" || c.Model == "" {
		return fmt.Errorf("package not initialize correctly")
	}
	return nil
}

func (c *ChatGPT) New(apiKey string, model Models, urlVersion URLVersion) {
	c.BaseURI = API_V1_URL
	c.ApiKey = apiKey
	c.Model = model
}

func (c *ChatGPT) NewFree(apiKey string) {
	c.BaseURI = API_V1_URL
	c.ApiKey = apiKey
	c.Model = GPT_35_TURBO
}

func (c *ChatGPT) SayHello() (*DefaultResponse, error) {
	if err := c.check(); err != nil {
		return nil, err
	}

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
		return nil, nil
	}

	// Create a new POST request with the JSON data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, nil
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)
	// Perform the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, nil
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, nil
	}

	var responseBody DefaultResponseBody

	if err := json.Unmarshal(body, &responseBody); err != nil {
		return nil, err
	}
	return &DefaultResponse{
		Body:  responseBody,
		Code:  response.StatusCode,
		Error: nil,
	}, nil
}
