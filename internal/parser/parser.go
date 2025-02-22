package parser

import (
	"encoding/json"

	"github.com/whookdev/cli/internal/models"
)

func ParseWSMessage(data []byte) ([]models.Item, error) {
	var items []models.Item
	err := json.Unmarshal(data, &items)
	return items, err
}
