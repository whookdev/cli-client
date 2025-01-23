package ui

import "github.com/charmbracelet/lipgloss"

const (
	// Layout constants
	ModalWidth             = 65
	ModalHorizontalPadding = 2
)

var (
	// TitleStyle defines the style for modal titles
	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("62")).
			Bold(true).
			MarginBottom(1).
			Width(ModalWidth - (ModalHorizontalPadding * 2)).
			Align(lipgloss.Center)

	// SelectedItemStyle defines the style for selected items in lists
	SelectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("170")).
				Bold(true)

	// NormalItemStyle defines the style for unselected items in lists
	NormalItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245"))

	// ModalStyle defines the base style for modals
	ModalStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Padding(1, ModalHorizontalPadding).
			Width(ModalWidth)

	// InputStyle defines the style for text input fields
	InputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Padding(0, 1).
			MarginTop(1).
			Width(ModalWidth - (ModalHorizontalPadding * 4))

	// HelpStyle defines the style for the help bar
	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Border(lipgloss.RoundedBorder()).
			BorderTop(true).
			BorderBottom(false).
			BorderLeft(false).
			BorderRight(false).
			BorderForeground(lipgloss.Color("240")).
			Padding(0, 2).
			Width(100).
			Align(lipgloss.Center)
)
