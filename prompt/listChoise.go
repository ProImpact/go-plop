package prompt

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type listModel struct {
	choices      []string
	promtMessage string
	cursor       int
	selected     map[int]struct{}
}

func OptionListModel(promtMessage string, choises ...string) *listModel {
	return &listModel{
		choices:      choises,
		promtMessage: promtMessage,
		selected:     make(map[int]struct{}),
	}
}

func (l *listModel) Init() tea.Cmd {
	return nil
}

func (l *listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return l, tea.Quit

		case "up", "k":
			if l.cursor > 0 {
				l.cursor--
			}

		case "down", "j":
			if l.cursor < len(l.choices)-1 {
				l.cursor++
			}

		case "enter", " ":
			_, ok := l.selected[l.cursor]
			if ok {
				delete(l.selected, l.cursor)
			} else {
				l.selected[l.cursor] = struct{}{}
			}
		}
	}

	return l, nil
}

func (l *listModel) View() string {
	s := fmt.Sprintf("%s\n\n", l.promtMessage)

	for i, choice := range l.choices {

		cursor := " " // no cursor
		if l.cursor == i {
			cursor = ">" // cursor!
		}

		checked := " " // not selected
		if _, ok := l.selected[i]; ok {
			checked = "x" // selected!
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"
	return s
}
