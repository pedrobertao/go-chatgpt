package modules

import (
	auth "github.com/pedrobertao/go-chatgpt/auth"
)

type Modules struct {
	auth auth.Config
}

func (m *Modules) sayHello() {

}
