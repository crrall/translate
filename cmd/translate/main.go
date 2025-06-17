package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os/exec"

	"translate/internal/api"
	"translate/internal/config"
	"translate/internal/utils"
)

func main() {
	opts := &config.Options{}
	opts.Init()
	opts.GetText()

	promptStr := utils.BuildPrompt(opts.Lang, opts.Text)

	params := api.Params{}
	params.BuildParams(*opts.Model, promptStr, false)

	jsonData, err := json.Marshal(params)
	if err != nil {
		fmt.Println("json marshal error:", err)
		return
	}

	resp, err := api.SendRequest(jsonData)
	if err != nil {
		fmt.Println("API request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	var result api.OllamaResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}

	exec.Command("zenity", "--info", "--title", "翻译("+*opts.Lang+")", "--text", result.Response).Run()
}
