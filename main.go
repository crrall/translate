package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type OllamaResponse struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Response           string    `json:"response"`
	Done               bool      `json:"done"`
	Context            []int     `json:"context"`
	TotalDuration      int64     `json:"total_duration"`
	LoadDuration       int64     `json:"load_duration"`
	PromptEvalCount    int       `json:"prompt_eval_count"`
	PromptEvalDuration int64     `json:"prompt_eval_duration"`
	EvalCount          int       `json:"eval_count"`
	EvalDuration       int64     `json:"eval_duration"`
}

func main() {
	// 获取命名参数
	text := flag.String("t", "", "")
	lang := flag.String("l", "zh", "")
	model := flag.String("m", "gemma3:4b", "")

	flag.Parse()

	if *text == "" && flag.NArg() > 0 {
		*text = strings.Join(flag.Args(), " ")
	}

	// 构建prompt
	var builder strings.Builder
	builder.WriteString("Translate the following into ")
	builder.WriteString(*lang)
	builder.WriteString(". ")
	builder.WriteString("Only output the translated sentence. No explanations or emojis:")
	builder.WriteString("\n\"")
	builder.WriteString(*text)
	builder.WriteString("\"")
	prompt := builder.String()

	// 封装请求参数
	data := struct {
		Model  string `json:"model"`
		Prompt string `json:"prompt"`
		Stream bool   `json:"stream"`
	}{
		Model:  *model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json marshal error:", err)
		return
	}

	// 创建请求
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:11434/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("NewRequest error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("do error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	var result OllamaResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Unmarshal error:", err)
	}

	fmt.Println(result.Response)

}
