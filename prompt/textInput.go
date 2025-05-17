package prompt

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type textinputModel struct {
	promptMessage string
	textInput     textinput.Model
	err           error
}

func TextPromptModel(promptMessage string) *textinputModel {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return &textinputModel{
		textInput:     ti,
		promptMessage: promptMessage,
		err:           nil,
	}
}

func (t *textinputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (t *textinputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return t, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		t.err = msg
		return t, nil
	}

	t.textInput, cmd = t.textInput.Update(msg)
	return t, cmd
}

func (t *textinputModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		t.promptMessage,
		t.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
