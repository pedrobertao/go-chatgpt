package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *ChatGPT) Post(message string, model Models) (*DefaultResponse, error) {
	url := fmt.Sprintf("%s%s", API_V1_URL, "/chat/completions")
	data := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{"role": "user", "content": message},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil, nil
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, nil
	}
	defer response.Body.Close()

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
