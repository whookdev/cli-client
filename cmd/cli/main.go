package main

import (
	"log"

	"github.com/rivo/tview"
	"github.com/whookdev/cli/internal/models"
	"github.com/whookdev/cli/internal/state"
	"github.com/whookdev/cli/internal/ui"
)

func main() {
	app := tview.NewApplication()
	ui := ui.NewUI(app)
	state := state.NewAppState()

	state.OnUpdate(func(items []models.Item) {
		app.QueueUpdateDraw(func() {
			ui.UpdateList(items)
		})
	})

	if err := app.SetRoot(ui.GetLayout(), true).EnableMouse(true).Run(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
