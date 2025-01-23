package components

import (
	"github.com/whookdev/cli/internal/ui"
)

// Modal represents a modal window component
type Modal struct {
	width  int
	height int
}

// NewModal creates a new modal component
func NewModal() Modal {
	return Modal{}
}

// SetSize updates the modal dimensions
func (m *Modal) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func (m Modal) RenderNewProject(input Input) string {
	titleContent := ui.TitleStyle.Render("Create New Project")
	inputContent := ui.InputStyle.Render(input.value + "█")

	return ui.ModalStyle.Render(
		titleContent + "\n\n" +
			"Enter project name:\n" +
			inputContent,
	)
}
