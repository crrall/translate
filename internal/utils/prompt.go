package utils

import "strings"

func BuildPrompt(lang *string, text *string) string {
	var builder strings.Builder
	builder.WriteString("Translate the following into ")
	builder.WriteString(*lang)
	builder.WriteString(". ")
	builder.WriteString("Only output the translated sentence. No explanations or emojis:\n\"")
	builder.WriteString(*text)
	builder.WriteString("\"")
	return builder.String()
}
