// Package sidebar provides a simple sidebar model.
package sidebar

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("62"))

type Model struct {
	width, height int
	table         table.Model
}

func New() Model {
	return Model{}
}

func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "enter":
			return m, tea.Batch(tea.Println("Selected", m.table.SelectedRow()[1]))
		}
	case tea.WindowSizeMsg:
		t := createTable(m.width)
		m.table = t

	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return style.
		Width(m.width - 4).
		Height(m.height - 2).
		Render(m.table.View() + "\n")
}

func createTable(width int) table.Model {
	// Width - Method column (6) - padding (12)
	endpointWidth := width - 6 - 12
	columns := []table.Column{
		{Title: "Method", Width: 6},
		{Title: "Endpoint", Width: endpointWidth},
	}

	rows := []table.Row{
		{"GET", "/api/v1/resource"},
		{"POST", "/api/v1/resource"},
		{"PUT", "/api/v1/resource/1"},
		{"DELETE", "/api/v1/resource/1"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("28")).
		Bold(true)
	t.SetStyles(s)

	return t
}
