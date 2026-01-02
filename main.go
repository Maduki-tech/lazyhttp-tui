package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	requestView textinput.Model
	resultView  textinput.Model
}

func newTextInput(placeholder string) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return ti
}

func NewModel() model {
	return model{
		requestView: newTextInput("Request"),
		resultView:  newTextInput("Result"),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, m.requestView.View(), m.resultView.View())
}

func main() {
	model := NewModel()
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
