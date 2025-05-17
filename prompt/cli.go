package prompt

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Main() {
}

func NewPrompt(model tea.Model, opt ...tea.ProgramOption) {
	p := tea.NewProgram(model, opt...)
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
