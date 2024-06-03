package chat

import (
	"errors"
	"fmt"
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

	return "", nil
}
