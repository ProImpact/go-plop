package main

import "github.com/ProImpact/goplop/prompt"

func main() {
	prompt.NewPrompt(prompt.TextPromptModel(
		"Whats your name",
	))
}
