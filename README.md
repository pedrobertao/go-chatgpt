
# go-chatGPT

A simple wrapper for the ChatGPT API 

Only text-based prompts for now


## Example
```go
// Example Promtp
res, err = chat.SendPrompt("Tell me the story of golang with 200 characters")
if err != nil {
  fmt.Println(err)
  os.Exit(1)
}
fmt.Println(res)
```

> ***Response:***
> Go, also known as Golang, was created by Google engineers in 2007. It is a statically typed, compiled language designed for efficiency and simplicity, gaining popularity for its clean syntax and powerful concurrency features.
## Licença

[MIT](https://choosealicense.com/licenses/mit/)

