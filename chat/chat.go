package chat

import (
	"errors"
	"fmt"
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

func (c *ChatGPT) SayHello() (string, error) {
	if err := c.check(); err != nil {
		return "", err
	}

	res, err := c.Post("Hello, how are you ?", GPT_35_TURBO)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New("failed to call api")
	}

	if len(res.Body.Choices) > 0 {
		strRe := fmt.Sprintf(res.Body.Choices[0].Message.Content + "\n")
		return strRe, nil
	}

	return "", nil
}
