package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
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

type options struct {
	Text  *string
	Lang  *string
	Model *string
}

func (o *options) Init() {
	execName := filepath.Base(os.Args[0])

	var defaultLang string

	switch execName {
	case "tzh":
		defaultLang = "zh"
	case "ten":
		defaultLang = "en"
	default:
		defaultLang = "zh"
	}

	o.Text = flag.String("t", "", "Text to process")
	o.Lang = flag.String("l", defaultLang, "Language code")
	o.Model = flag.String("m", "gemma3:4b", "Model to use")
	flag.Parse()

	if *o.Text == "" && flag.NArg() > 0 {
		*o.Text = strings.Join(flag.Args(), " ")
	}
}

func (o *options) GetText() {
	if *o.Text == "" {
		cmd := exec.Command("xclip", "-o", "-selection", "primary")
		out, err := cmd.Output()
		if err == nil {
			*o.Text = string(out)
		}
	}
}

func BuildPrompt(lang *string, text *string) string {
	var builder strings.Builder
	builder.WriteString("Translate the following into ")
	builder.WriteString(*lang)
	builder.WriteString(". ")
	builder.WriteString("Only output the translated sentence. No explanations or emojis:")
	builder.WriteString("\n\"")
	builder.WriteString(*text)
	builder.WriteString("\"")
	return builder.String()
}

type params struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

func (p *params) BuildParams(model string, prompt string, stream bool) {
	p.Model = model
	p.Prompt = prompt
	p.Stream = stream
}

func main() {
	// 获取命名参数
	opts := &options{}
	opts.Init()
	opts.GetText()

	// 构建prompt
	prompt := BuildPrompt(opts.Lang, opts.Text)

	// 封装请求参数
	data := struct {
		Model  string `json:"model"`
		Prompt string `json:"prompt"`
		Stream bool   `json:"stream"`
	}{
		Model:  *opts.Model,
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

	exec.Command("zenity", "--info", "--title", "翻译("+*opts.Lang+")", "--text", result.Response).Run()

}
