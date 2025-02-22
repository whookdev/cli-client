package state

import "github.com/whookdev/cli/internal/models"

type UpdateCallback func([]models.Item)

type AppState struct {
	items     []models.Item
	callbacks []UpdateCallback
}

func NewAppState() *AppState {
	return &AppState{
		items:     make([]models.Item, 0),
		callbacks: make([]UpdateCallback, 0),
	}
}

func (s *AppState) OnUpdate(callback UpdateCallback) {
	s.callbacks = append(s.callbacks, callback)
}

func (s *AppState) UpdateItems(items []models.Item) {
	s.items = items
	for _, callback := range s.callbacks {
		callback(items)
	}
}
