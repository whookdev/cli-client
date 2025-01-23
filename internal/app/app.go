package app

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/whookdev/cli/internal/ui"
	"github.com/whookdev/cli/internal/ui/components"
)

func Run() error {
	p := tea.NewProgram(NewModel(), tea.WithAltScreen())
	_, err := p.Run()
	return err
}

func (m Model) Init() tea.Cmd {
	// No initial commands needed
	return nil
}

func (m Model) View() string {
	var content string
	var helpText string

	switch m.viewState {
	case ViewStateProjects:
		content = m.list.Render()
		helpText = "↑/k: Move up   ↓/j: Move down   c: Create new project   enter: Select project   q/esc: Quit"
	case ViewStateNewProject:
		content = m.modal.RenderNewProject(m.input)
		helpText = "enter: Save project   esc: Cancel"
	}

	if m.windowSize.Width == 0 {
		return content
	}

	modalContent := content
	modalHeight := lipgloss.Height(modalContent)
	verticalPadding := (m.windowSize.Height - modalHeight - 2) / 2

	return lipgloss.JoinVertical(lipgloss.Center,
		strings.Repeat("\n", verticalPadding),
		lipgloss.Place(
			m.windowSize.Width,
			modalHeight,
			lipgloss.Center,
			lipgloss.Center,
			modalContent,
		),
		strings.Repeat("\n", m.windowSize.Height-modalHeight-verticalPadding-2),
		ui.HelpStyle.Width(m.windowSize.Width).Render(helpText),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowSize = msg
		m.help.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	}

	return m, nil
}

func (m Model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.viewState {
	case ViewStateProjects:
		return m.handleProjectListKeys(msg)
	case ViewStateNewProject:
		return m.handleNewProjectKeys(msg)
	default:
		return m, nil
	}
}

func (m Model) handleProjectListKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q", "esc":
		return m, tea.Quit

	case "up", "k":
		m.list.MoveCursor(true)

	case "down", "j":
		m.list.MoveCursor(false)

	case "c":
		m.viewState = ViewStateNewProject
		m.input.Clear()

	case "enter", " ":
		// Handle project selection
		return m, tea.Quit
	}

	return m, nil
}

func (m Model) handleNewProjectKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEsc:
		m.viewState = ViewStateProjects
		return m, nil

	case tea.KeyEnter:
		if m.input.Value() != "" {
			if err := m.projectManager.AddProject(m.input.Value()); err == nil {
				m.list = components.NewProjectList(m.projectManager.Projects())
				m.viewState = ViewStateProjects
			}
		}
		return m, nil

	case tea.KeyBackspace:
		m.input.Backspace()

	case tea.KeyRunes:
		m.input.HandleInput(string(msg.Runes))
	}

	return m, nil
}
