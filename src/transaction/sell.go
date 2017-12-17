package transaction

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"encoding/json"
	"hashutils"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Sell struct {
	ID   string `datastore:",noindex"`
	Etag string

	ChartD string `json:"-" datastore:",noindex"`
	ChartW string `json:"-" datastore:",noindex"`

	ChartNDQCD string `json:"-" datastore:",noindex"`
	ChartNDQCW string `json:"-" datastore:",noindex"`

	ChartSP5D string `json:"-" datastore:",noindex"`
	ChartSP5W string `json:"-" datastore:",noindex"`

	//ChartNYCD string `datastore:",noindex"`
	//ChartNYCW string `datastore:",noindex"`

	ChartDJIAD string `json:"-" datastore:",noindex"`
	ChartDJIAW string `json:"-" datastore:",noindex"`

	ChartRUSD string `json:"-" datastore:",noindex"`
	ChartRUSW string `json:"-" datastore:",noindex"`

	IBDCheckup string `json:"-" datastore:",noindex"`

	JsonChartD string `datastore:"-"`
	JsonChartW string `datastore:"-"`

	JsonChartNDQCD string `datastore:"-"`
	JsonChartNDQCW string `datastore:"-"`

	JsonChartSP5D string `datastore:"-"`
	JsonChartSP5W string `datastore:"-"`

	//JsonChartNYCD string `datastore:"-"`
	//JsonChartNYCW string `datastore:"-"`

	JsonChartDJIAD string `datastore:"-"`
	JsonChartDJIAW string `datastore:"-"`

	JsonChartRUSD string `datastore:"-"`
	JsonChartRUSW string `datastore:"-"`

	JsonIBDCheckup string `datastore:"-"`

	Date int

	Symbol string

	Price float32

	Share int

	Total float32

	Position float32

	GainLoss float32

	DayHeld int

	Stage float32

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

	//s.ChartDaily = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	//s.ChartWeekly = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	//s.ChartNasdaq = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	//s.ChartSP500 = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	//s.ChartNYSE = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	//s.ChartDowJones = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	//s.IBDCheckup = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
