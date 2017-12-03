package transaction

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"encoding/json"
	"hashutils"
	"time"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Sell struct {
	ID   string `datastore:",noindex"`
	Etag string

	ChartDaily  string `datastore:",noindex"`
	ChartWeekly string `datastore:",noindex"`

	ChartNasdaq   string `datastore:",noindex"`
	ChartSP500    string `datastore:",noindex"`
	ChartNYSE     string `datastore:",noindex"`
	ChartDowJones string `datastore:",noindex"`

	IBDCheckup string `datastore:",noindex"`

	Date time.Time

	Symbol string

	Price float32

	Share uint

	Percent float32

	GainLoss float32

	Stage float32

	//BoughtDate time.Time `datastore:"_"`

	DayHeld uint

	Note string `datastore:",noindex"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *Sell) JsonDecode(buffer []byte) error {

	err := json.Unmarshal(buffer, s)
	if err != nil {
		return err
	}

	s.ID = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	s.Etag = hashutils.RandBytesGenerateB64(config.KeyLengthMin)

	s.ChartDaily = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	s.ChartWeekly = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	s.ChartNasdaq = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	s.ChartSP500 = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	s.ChartNYSE = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	s.ChartDowJones = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	s.IBDCheckup = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////