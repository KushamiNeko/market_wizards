package study

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hashutils"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ChartsAnalysis struct {
	Id string

	From int
	To   int

	Symbol string
	Daily  string
	Weekly string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ChartsAnalysisNewBytes(from, to int, symbol string, daily, weekly []byte) *ChartsAnalysis {
	c := new(ChartsAnalysis)

	c.Id = hashutils.RandBytesB64(config.KeyLengthDefault)

	c.From = from
	c.To = to
	c.Symbol = symbol

	c.Daily = base64.StdEncoding.EncodeToString(daily)
	c.Weekly = base64.StdEncoding.EncodeToString(weekly)

	return c
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ChartsAnalysisNewString(from, to int, symbol string, daily, weekly string) *ChartsAnalysis {
	c := new(ChartsAnalysis)

	c.Id = hashutils.RandBytesB64(config.KeyLengthDefault)

	c.From = from
	c.To = to
	c.Symbol = symbol

	c.Daily = daily
	c.Weekly = weekly

	return c
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartsAnalysis) JsonDecode(buffer []byte) error {

	err := json.Unmarshal(buffer, c)
	if err != nil {
		return err
	}

	if c.From == 0 {
		return fmt.Errorf("From cannot be empty")
	}

	if c.To == 0 {
		return fmt.Errorf("To cannot be empty")
	}

	if c.Symbol == "" {
		return fmt.Errorf("Symbol cannot be empty")
	}

	if c.Daily == "" {
		return fmt.Errorf("Daily Data cannot be empty")
	}

	if c.Weekly == "" {
		return fmt.Errorf("Weekly Data cannot be empty")
	}

	c.Id = hashutils.RandBytesB64(config.KeyLengthDefault)

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
