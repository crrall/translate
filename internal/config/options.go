package config

import (
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Options struct {
	Text  *string
	Lang  *string
	Model *string
	Cli   *bool
}

func (o *Options) Init() {
	execName := filepath.Base(os.Args[0])
	defaultLang := "zh"
	if execName == "ten" {
		defaultLang = "en"
	}

	o.Text = flag.String("t", "", "Text to process")
	o.Lang = flag.String("l", defaultLang, "Language code")
	o.Model = flag.String("m", "gemma3:4b", "Model to use")
	flag.Parse()

	if *o.Text == "" && flag.NArg() > 0 {
		*o.Text = strings.Join(flag.Args(), " ")
	}

	o.Cli = new(bool)
	*o.Cli = true
}

func (o *Options) GetText() {
	if *o.Text == "" {
		cmd := exec.Command("xclip", "-o", "-selection", "primary")
		if out, err := cmd.Output(); err == nil {
			*o.Text = string(out)
			*o.Cli = false
		}
	}
}
