package transaction

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"encoding/json"
	"fmt"
	"hashutils"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type BuyOrder struct {
	//JsonIBDCheckup string `bson:"-"`

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

//func (b *BuyOrder) GetJsonIBDCheckup() string {
//return b.JsonIBDCheckup
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (b *BuyOrder) GetDate() int {
	return b.Date
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (b *BuyOrder) GetSymbol() string {
	return b.Symbol
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (b *BuyOrder) JsonDecode(buffer []byte) error {

	err := json.Unmarshal(buffer, b)
	if err != nil {
		return err
	}

	if b.Order != "buy" && b.Order != "sell" {
		return fmt.Errorf("Invalid Order")
	}

	b.ID = hashutils.RandBytesB64URL(config.KeyLengthStrong)
	b.Etag = hashutils.RandBytesB64URL(config.KeyLengthMin)

	//if b.JsonIBDCheckup == "" {
	//return fmt.Errorf("IBD Checkup Cannot Be Empty")
	//}

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

func (b *BuyOrder) GetIBDCheckupID() string {
	return fmt.Sprintf("%d_%v", b.Date, b.Symbol)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
