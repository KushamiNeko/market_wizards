package transaction

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"encoding/json"
	"fmt"
	"hashutils"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Open struct {
	ID       string
	Etag     string
	Order    string
	Date     int
	Symbol   string
	Price    float64
	Share    int
	BuyPoint string
	Cost     float64
	Stage    float64
	Note     string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (b *Open) JsonDecode(buffer []byte) error {

	err := json.Unmarshal(buffer, b)
	if err != nil {
		return err
	}

	if b.Order != "buy" && b.Order != "sell" {
		return fmt.Errorf("Invalid Order")
	}

	b.ID = hashutils.RandBytesB64URL(config.KeyLengthStrong)
	b.Etag = hashutils.RandBytesB64URL(config.KeyLengthMin)

	if b.Date == 0 {
		return fmt.Errorf("Date Cannot Be Empty")
	}

	if b.Symbol == "" {
		return fmt.Errorf("Symbol Cannot Be Empty")
	}

	if b.Price == 0.0 {
		return fmt.Errorf("Price Cannot Be Empty")
	}

	if b.Share == 0 {
		return fmt.Errorf("Share Cannot Be Empty")
	}

	if b.BuyPoint == "" {
		return fmt.Errorf("BuyPoint Cannot Be Empty")
	}

	if b.Cost == 0.0 {
		return fmt.Errorf("Cost Cannot Be Empty")
	}

	if b.Stage == 0.0 {
		return fmt.Errorf("Stage Cannot Be Empty")
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
