package ui

import (
	"encoding/json"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/whookdev/cli/internal/models"
)

type UI struct {
	app      *tview.Application
	list     *tview.List
	jsonView *tview.TextView
	layout   *tview.Flex
}

func NewUI(app *tview.Application) *UI {
	ui := &UI{
		app:      app,
		list:     tview.NewList(),
		jsonView: tview.NewTextView(),
	}

	ui.list.SetBorder(true)
	ui.list.SetTitle("Requests")

	ui.jsonView.SetDynamicColors(true)
	ui.jsonView.SetRegions(true)
	ui.jsonView.SetWordWrap(true)
	ui.jsonView.SetBorder(true)
	ui.jsonView.SetTitle("JSON Content")

	ui.list.SetChangedFunc(func(index int, _ string, _ string, _ rune) {
		ui.updateJSON(index)
	})

	ui.layout = tview.NewFlex().
		AddItem(ui.list, 0, 1, true).
		AddItem(ui.jsonView, 0, 2, false)

	app.SetInputCapture(ui.handleInput)

	return ui
}

func (ui *UI) GetLayout() *tview.Flex {
	return ui.layout
}

func (ui *UI) UpdateList(items []models.Item) {
	ui.list.Clear()
	for i, item := range items {
		ui.list.AddItem(item.Title, item.Description, rune('a'+i), nil)
	}
}

func (ui *UI) updateJSON(index int) {
	if ui.list.GetItemCount() == 0 {
		return
	}

	title, description := ui.list.GetItemText(index)
	data := map[string]interface{}{
		"selected_index": index,
		"seleceed_item": models.Item{
			Title:       title,
			Description: description,
		},
		"total_items": ui.list.GetItemCount(),
	}

	jsonBytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		ui.jsonView.SetText(fmt.Sprintf("Error: %v", err))
		return
	}

	ui.jsonView.SetText(string(jsonBytes))
}

func (ui *UI) handleInput(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEsc || event.Rune() == 'q' {
		ui.app.Stop()
		return nil
	}

	return event
}
