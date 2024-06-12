package chat

import (
	"errors"
	"fmt"

	sanitize "github.com/mrz1836/go-sanitize"
)

func (c *ChatGPT) SayHello() (string, error) {
	if err := c.check(); err != nil {
		return "", err
	}

	res, err := c.Post("Hello, how are you ?")
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New("failed to call api")
	}

	if len(res.Body.Choices) > 0 {
		strRe := fmt.Sprintf(res.Body.Choices[0].Message.Content)
		return strRe, nil
	}

	return "", fmt.Errorf("no message content")
}

func (c *ChatGPT) SendPrompt(prompt string) (string, error) {
	if err := c.check(); err != nil {
		return "", err
	}

	res, err := c.Post(sanitize.AlphaNumeric(prompt, true))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New("failed to call api")
	}

	if len(res.Body.Choices) > 0 {
		strRe := fmt.Sprintf(res.Body.Choices[0].Message.Content)
		return strRe, nil
	}

	if len(res.Body.Choices) > 0 {
		strRe := fmt.Sprintf(res.Body.Choices[0].Message.Content)
		return strRe, nil
	}

	return "", fmt.Errorf("no message content")

}
