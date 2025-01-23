package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/whookdev/cli/internal/projects"
	"github.com/whookdev/cli/internal/ui/components"
)

type ViewState int

const (
	ViewStateProjects ViewState = iota
	ViewStateNewProject
)

type Model struct {
	projectManager *projects.Manager
	list           components.ProjectList
	modal          components.Modal
	help           components.Help
	viewState      ViewState
	windowSize     tea.WindowSizeMsg
	input          components.Input
}

func NewModel() Model {
	pm := projects.NewManager()
	return Model{
		projectManager: pm,
		list:           components.NewProjectList(pm.Projects()),
		modal:          components.NewModal(),
		help:           components.NewHelp(),
		viewState:      ViewStateProjects,
		input:          components.NewInput(),
	}
}
