package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func authorize() {

}

func refreshAuth() {

}

func callChatGpt() {

}

/*
* ENDPOINT	LATEST MODELS
/v1/assistants	All GPT-4 and GPT-3.5 Turbo models except gpt-3.5-turbo-0301 supported. The retrieval tool requires gpt-4-turbo-preview (and subsequent dated model releases) or gpt-3.5-turbo-1106 (and subsequent versions).
/v1/audio/transcriptions	whisper-1
/v1/audio/translations	whisper-1
/v1/audio/speech	tts-1, tts-1-hd
/v1/chat/completions	gpt-4 and dated model releases, gpt-4-turbo-preview and dated model releases, gpt-4-vision-preview, gpt-4-32k and dated model releases, gpt-3.5-turbo and dated model releases, gpt-3.5-turbo-16k and dated model releases, fine-tuned versions of gpt-3.5-turbo
/v1/completions (Legacy)	gpt-3.5-turbo-instruct, babbage-002, davinci-002
/v1/embeddings	text-embedding-3-small, text-embedding-3-large, text-embedding-ada-002
/v1/fine_tuning/jobs	gpt-3.5-turbo, babbage-002, davinci-002
/v1/moderations	text-moderation-stable, text-moderation-latest
/v1/images/generations	dall-e-2, dall-e-3*
*/
func main() {
	values := map[string]string{"foo": "baz"}
	jsonData, err := json.Marshal(values)

	req, err := http.NewRequest(http.MethodPost, "https://httpbin.org/post", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	// appending to existing query args
	q := req.URL.Query()
	q.Add("foo", "bar")

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))
}
