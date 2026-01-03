// Package mainview provides the main view model for the application.
package mainview

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("63")).
	Padding(1, 2)

type Model struct {
	width, height int
}

func New() Model {
	return Model{}
}

func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func (m Model) View() string {
	content := "Main View\n"

	return style.Width(m.width - 4).Height(m.height - 2).Render(content)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}
