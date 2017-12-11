package transaction

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"encoding/json"
	"hashutils"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Buy struct {
	ID   string `datastore:",noindex"`
	Etag string

	ChartD string `datastore:",noindex"`
	ChartW string `datastore:",noindex"`

	ChartNDQCD string `datastore:",noindex"`
	ChartNDQCW string `datastore:",noindex"`

	ChartSP5D string `datastore:",noindex"`
	ChartSP5W string `datastore:",noindex"`

	ChartNYCD string `datastore:",noindex"`
	ChartNYCW string `datastore:",noindex"`

	ChartDJIAD string `datastore:",noindex"`
	ChartDJIAW string `datastore:",noindex"`

	ChartRUSD string `datastore:",noindex"`
	ChartRUSW string `datastore:",noindex"`

	IBDCheckup string `datastore:",noindex"`

	Date int

	Symbol string

	Price float32

	Share uint

	Total float32

	Capital float32

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

	b.ChartD = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	b.ChartW = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	b.ChartNDQCD = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	b.ChartNDQCW = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	b.ChartSP5D = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	b.ChartSP5W = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	b.ChartNYCD = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	b.ChartNYCW = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	b.ChartDJIAD = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	b.ChartDJIAW = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	b.ChartRUSD = hashutils.RandBytesGenerateB64(config.KeyLengthMax)
	b.ChartRUSW = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	b.IBDCheckup = hashutils.RandBytesGenerateB64(config.KeyLengthMax)

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
