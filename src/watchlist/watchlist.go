package watchlist

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"encoding/json"
	"fmt"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

type WatchListItem struct {
	Symbol string
	Price  float64

	//Operation string

	Priority string

	GRS string
	RS  string

	Fundamentals string

	Status string
	Note   string

	Flag bool

	PositionSize int `bson:"-" json:"-"`
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

	if w.RS == "" {
		return fmt.Errorf("Relative Strength cannot be empty")
	}

	if w.GRS == "" {
		return fmt.Errorf("Group Relative Strength cannot be empty")
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

func (w *WatchListItem) Caculate(capital, size float64) {
	w.PositionSize = int((capital * (size / 100.0)) / w.Price)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
