package transaction

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"encoding/json"
	"fmt"
	"hashutils"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Close struct {
	ID             string
	Etag           string
	Order          string
	Date           int
	Symbol         string
	Price          float64
	Share          int
	DateOfPurchase int
	Revenue        float64
	Cost           float64
	GainD          float64
	GainP          float64
	DaysHeld       int
	Note           string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *Close) JsonDecode(buffer []byte) error {

	err := json.Unmarshal(buffer, s)
	if err != nil {
		return err
	}

	if s.Order != "buy" && s.Order != "sell" {
		return fmt.Errorf("Invalid Order")
	}

	s.ID = hashutils.RandBytesB64URL(config.KeyLengthStrong)
	s.Etag = hashutils.RandBytesB64URL(config.KeyLengthMin)

	if s.Date == 0 {
		return fmt.Errorf("Date Cannot Be Empty")
	}

	if s.Symbol == "" {
		return fmt.Errorf("Symbol Cannot Be Empty")
	}

	if s.Price == 0.0 {
		return fmt.Errorf("Price Cannot Be Empty")
	}

	if s.Share == 0 {
		return fmt.Errorf("Share Cannot Be Empty")
	}

	if s.Cost == 0.0 {
		return fmt.Errorf("Cost Cannot Be Empty")
	}

	if s.DateOfPurchase == 0 {
		return fmt.Errorf("DateOfPurchase Cannot Be Empty")
	}

	if s.Revenue == 0.0 {
		return fmt.Errorf("Revenue Cannot Be Empty")
	}

	if s.GainD == 0.0 {
		return fmt.Errorf("GainD Cannot Be Empty")
	}

	if s.GainP == 0.0 {
		return fmt.Errorf("GainP Cannot Be Empty")
	}

	if s.DaysHeld == 0 {
		return fmt.Errorf("DaysHeld Cannot Be Empty")
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
