package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rivo/tview"
	"github.com/whookdev/cli/internal/config"
	"github.com/whookdev/cli/internal/models"
	"github.com/whookdev/cli/internal/relay"
	"github.com/whookdev/cli/internal/state"
	"github.com/whookdev/cli/internal/ui"
	"github.com/whookdev/cli/internal/wsclient"
)

func main() {
	cfg := config.Load()
	app := tview.NewApplication()
	ui := ui.NewUI(app)
	state := state.NewAppState()
	relay := relay.NewRelay(cfg)

	ctx := context.Background()

	wsClient, err := wsclient.Connect(ctx, relay)
	if err != nil {
		log.Fatalf("Error creating websocket client: %v", err)
	}
	defer wsClient.Close()

	for {
		msg, err := wsClient.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("Recv: %s\n", string(msg))
	}

	state.OnUpdate(func(items []models.Item) {
		app.QueueUpdateDraw(func() {
			ui.UpdateList(items)
		})
	})

	// if err := app.SetRoot(ui.GetLayout(), true).EnableMouse(true).Run(); err != nil {
	// 	log.Fatalf("Error running application: %v", err)
	// }
}
