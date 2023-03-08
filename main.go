package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// request struct
type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Response struct
type ChatCompletionResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
	Choices []Choice `json:"choices"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Choice struct {
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
	Index        int         `json:"index"`
}

func GetApiKey() (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if len(apiKey) == 0 {
		return "", errors.New("set env variable `OPENAI_API_KEY`")
	}
	return apiKey, nil
}

func main() {
	content := flag.String("c", "", "Content to use for chat prompt (Required)")
	system_content := flag.String("sc", "Your are chatbot. Always be more detail", "System context")
	flag.Parse()

	if *content == "" {
		fmt.Println("Content is required, pass using -c")
		return
	}

	// get api key
	apiKey, err := GetApiKey()

	if err != nil {
		fmt.Println(err)
		return
	}

	// make request body
	conversation := []ChatMessage{
		{Role: "system", Content: *system_content},
		{Role: "user", Content: *content},
	}

	requestBody, err := json.Marshal(ChatRequest{
		Model:    "gpt-3.5-turbo",
		Messages: conversation,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions",
		strings.NewReader(string(requestBody)))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	// get json and only print text response
	var completionResponse ChatCompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&completionResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, choice := range completionResponse.Choices {
		fmt.Println(choice.Message.Content)
	}
}
