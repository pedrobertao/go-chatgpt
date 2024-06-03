package chat

import (
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
