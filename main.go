package main

import (
	"encoding/json"
	flag "github.com/spf13/pflag"
	"log"
	"net/http"
	"os"
	"strings"
	"cgpt/config"
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
	apiKey := os.Getenv("OPENAI_API_KEYS")
	if len(apiKey) == 0 {
		apiKey = config.SaveOrGetToken()
		// return "", errors.New("set env variable `OPENAI_API_KEY`")
	}
	return apiKey, nil
}

func main() {
	var query string
	var system_content_final string
	content := flag.StringP("query", "q", "", "Query for chat prompt. You can pass just Args to, first Args is taken as query.")
	system_content := flag.StringP("system_context", "s", "Your are chatbot. Always be more detail.", "System context")
	code_bool := flag.BoolP("code", "c", false, "Flag to just get code.")
	json_bool := flag.BoolP("json", "j", false, "Flag to just get json.")
	grammar_bool := flag.BoolP("grammar", "g", false, "Flag to fix grammar of text.")
	debug := flag.BoolP("debug", "d", false, "Print query, System context and Response.")
	flag.Parse()
	
	args := flag.Args()

	if len(args) > 0 {
		query = args[0]
	} else if *content == "" {
		log.Println("Content is required, pass using -q or send first argument")
		return
	} else {
		query = *content
	}


	//make System context base on flag
	if *code_bool {
		system_content_final = "Just get the code no explaination or other text"
	} else if *json_bool {
		system_content_final = "Do infromation extraction and make it detailed. Just Json output"
	} else if  *grammar_bool {
		system_content_final = "Fixed the grammar and make it better for formal conversation."
	} else {
		system_content_final = *system_content
	}

	// get api key
	apiKey, err := GetApiKey()

	if err != nil {
		log.Println(err)
		return
	}

	// make request body
	conversation := []ChatMessage{
		{Role: "system", Content: system_content_final},
		{Role: "user", Content: query},
	}

	requestBody, err := json.Marshal(ChatRequest{
		Model:    "gpt-3.5-turbo",
		Messages: conversation,
	})

	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions",
		strings.NewReader(string(requestBody)))
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	// get json and only print text response
	var completionResponse ChatCompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&completionResponse)
	if err != nil {
		log.Println(err)
		return
	}

	// print for debug 
	if *debug {
		log.Println("Query: ", query)
		log.Println("System Content: ", system_content_final)
		log.Println("Response: ", completionResponse)
		}

	for _, choice := range completionResponse.Choices {
		log.Println(choice.Message.Content)
	}
}
