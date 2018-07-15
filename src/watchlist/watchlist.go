package watchlist

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"encoding/json"
	"fmt"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

type WatchListItem struct {
	Symbol string
	Price  float64

	Priority     string
	Fundamentals string
	Status       string
	//Note         string

	PositionSize []int `bson:"-" json:"-"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func (w *WatchListItem) JsonDecode(buffer []byte) error {

	err := json.Unmarshal(buffer, w)
	if err != nil {
		return err
	}

	if w.Symbol == "" {
		return fmt.Errorf("Symbol cannot be empty")
	}

	if w.Price == 0 {
		return fmt.Errorf("Price cannot be empty")
	}

	if w.Priority == "" {
		return fmt.Errorf("Priority cannot be empty")
	}

	if w.Fundamentals == "" {
		return fmt.Errorf("Fundamentals cannot be empty")
	}

	if w.Status == "" {
		return fmt.Errorf("Status cannot be empty")
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (w *WatchListItem) Caculate(capital float64) {

	w.PositionSize = make([]int, len(config.WatchListPosition))

	for i, p := range config.WatchListPosition {
		c := capital * p
		share := int(c / w.Price)

		w.PositionSize[i] = share
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
