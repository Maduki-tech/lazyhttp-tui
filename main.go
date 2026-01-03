package main

import (
	"fmt"
	"os"

	"lazyhttp-tui/internal/mainview"
	"lazyhttp-tui/internal/sidebar"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width    int
	height   int
	sidebar  sidebar.Model
	mainview mainview.Model
}

func initialModel() model {
	return model{
		sidebar:  sidebar.New(),
		mainview: mainview.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		sidebarWidth := m.width / 3
		mainViewWidth := m.width - sidebarWidth

		m.sidebar.SetSize(sidebarWidth, m.height)
		m.mainview.SetSize(mainViewWidth, m.height)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.sidebar, cmd = m.sidebar.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, m.sidebar.View(), m.mainview.View())
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
