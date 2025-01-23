package components

import "github.com/whookdev/cli/internal/ui"

type Help struct {
	width int
}

func NewHelp() Help {
	return Help{}
}

func (h Help) Render(text string) string {
	style := ui.HelpStyle
	if h.width > 0 {
		style = style.Width(h.width)
	}
	return style.Render(text)
}

func (h *Help) SetWidth(width int) {
	h.width = width
}
