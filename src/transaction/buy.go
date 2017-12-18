package transaction

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"encoding/json"
	"fmt"
	"hashutils"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Buy struct {
	ID   string `datastore:",noindex"`
	Etag string

	ChartD string `json:"-" datastore:",noindex"`
	ChartW string `json:"-" datastore:",noindex"`

	ChartNDQCD string `json:"-" datastore:",noindex"`
	ChartNDQCW string `json:"-" datastore:",noindex"`

	ChartSP5D string `json:"-" datastore:",noindex"`
	ChartSP5W string `json:"-" datastore:",noindex"`

	ChartNYCD string `datastore:",noindex"`
	ChartNYCW string `datastore:",noindex"`

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

	JsonChartNYCD string `datastore:"-"`
	JsonChartNYCW string `datastore:"-"`

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

	Capital float32 `datastore:",noindex"`

	Position float32 `datastore:",noindex"`

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

	if b.JsonChartD == "" {
		return fmt.Errorf("Dialy Chart Cannot Be Empty")
	}

	if b.JsonChartW == "" {
		return fmt.Errorf("Weekly Chart Cannot Be Empty")
	}

	b.ChartD = hashutils.RandBytesGenerateB64URL(config.KeyLengthMax)
	b.ChartW = hashutils.RandBytesGenerateB64URL(config.KeyLengthMax)

	if b.JsonChartNDQCD == "" {
		return fmt.Errorf("NDQC Dialy Chart Cannot Be Empty")
	}

	if b.JsonChartNDQCW == "" {
		return fmt.Errorf("NDQC Weekly Chart Cannot Be Empty")
	}

	b.ChartNDQCD = hashutils.RandBytesGenerateB64URL(config.KeyLengthMax)
	b.ChartNDQCW = hashutils.RandBytesGenerateB64URL(config.KeyLengthMax)

	if b.JsonChartSP5D == "" {
		return fmt.Errorf("S&P5 Dialy Chart Cannot Be Empty")
	}

	if b.JsonChartSP5W == "" {
		return fmt.Errorf("S&P5 Weekly Chart Cannot Be Empty")
	}

	b.ChartSP5D = hashutils.RandBytesGenerateB64URL(config.KeyLengthMax)
	b.ChartSP5W = hashutils.RandBytesGenerateB64URL(config.KeyLengthMax)

	if b.JsonChartNYCD == "" {
		return fmt.Errorf("NYC Dialy Chart Cannot Be Empty")
	}

	if b.JsonChartNYCW == "" {
		return fmt.Errorf("NYC Weekly Chart Cannot Be Empty")
	}

	b.ChartNYCD = hashutils.RandBytesGenerateB64URL(config.KeyLengthMax)
	b.ChartNYCW = hashutils.RandBytesGenerateB64URL(config.KeyLengthMax)

	if b.JsonChartDJIAD == "" {
		return fmt.Errorf("DJIA Dialy Chart Cannot Be Empty")
	}

	if b.JsonChartDJIAW == "" {
		return fmt.Errorf("DJIA Weekly Chart Cannot Be Empty")
	}

	b.ChartDJIAD = hashutils.RandBytesGenerateB64URL(config.KeyLengthMax)
	b.ChartDJIAW = hashutils.RandBytesGenerateB64URL(config.KeyLengthMax)

	if b.JsonChartRUSD == "" {
		return fmt.Errorf("RUS Dialy Chart Cannot Be Empty")
	}

	if b.JsonChartRUSW == "" {
		return fmt.Errorf("RUS Weekly Chart Cannot Be Empty")
	}

	b.ChartRUSD = hashutils.RandBytesGenerateB64URL(config.KeyLengthMax)
	b.ChartRUSW = hashutils.RandBytesGenerateB64URL(config.KeyLengthMax)

	if b.JsonIBDCheckup == "" {
		return fmt.Errorf("IBD Checkup Cannot Be Empty")
	}

	b.IBDCheckup = hashutils.RandBytesGenerateB64URL(config.KeyLengthMax)

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
