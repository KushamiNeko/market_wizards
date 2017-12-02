package transaction

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"encoding/json"
	"hashutils"
	"time"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Buy struct {
	ID   string `datastore:",noindex"`
	Etag string

	ChartDaily  string `datastore:",noindex"`
	ChartWeekly string `datastore:",noindex"`

	ChartNasdaq   string `datastore:",noindex"`
	ChartSP500    string `datastore:",noindex"`
	ChartNYSE     string `datastore:",noindex"`
	ChartDowJones string `datastore:",noindex"`

	IBDCheckup string `datastore:",noindex"`

	//DateHTTP string `datastore:"_"`

	Date time.Time

	Symbol string

	Price float32

	Share uint

	Position float32

	Stage float32

	Note string `datastore:",noindex"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (b *Buy) JsonDecode(buffer []byte) error {

	err := json.Unmarshal(buffer, b)
	if err != nil {
		return err
	}

	b.ID = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	b.Etag = hashutils.RandBytesGenerateB64(config.KeyLengthMin)

	b.ChartDaily = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	b.ChartWeekly = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	b.ChartNasdaq = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	b.ChartSP500 = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	b.ChartNYSE = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	b.ChartDowJones = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	b.IBDCheckup = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	//t, err := time.Parse(http.TimeFormat, b.DateHTTP)
	//if err != nil {
	//return err
	//}

	//b.Date = t

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
