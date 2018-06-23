package charts

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"datautils"
	"encoding/json"
	"marketsmith"
	"transaction"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ChartMarketSmith struct {
	filterOrders []*transaction.Transaction

	winnersMS []*bytes.Buffer
	losersMS  []*bytes.Buffer

	//msW []*marketsmith.MarketSmith
	//msL []*marketsmith.MarketSmith

	msW []datautils.Contents
	msL []datautils.Contents

	Alpha string
	Beta  string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ChartMarketSmithNew(filterOrders []*transaction.Transaction, winnersMS, losersMS []*bytes.Buffer) (*ChartMarketSmith, error) {

	c := new(ChartMarketSmith)

	c.filterOrders = filterOrders

	c.winnersMS = winnersMS
	c.losersMS = losersMS

	//c.msW = make([]*marketsmith.MarketSmith, len(c.winnersMS))
	//c.msL = make([]*marketsmith.MarketSmith, len(c.losersMS))

	c.msW = make([]datautils.Contents, len(c.winnersMS))
	c.msL = make([]datautils.Contents, len(c.losersMS))

	var err error

	for i, w := range c.winnersMS {
		m := marketsmith.MarketSmithNew()
		err = json.Unmarshal(w.Bytes(), m)
		if err != nil {
			return nil, err
		}

		c.msW[i] = m
	}

	for i, l := range c.losersMS {
		m := marketsmith.MarketSmithNew()
		err = json.Unmarshal(l.Bytes(), m)
		if err != nil {
			return nil, err
		}

		c.msL[i] = m
	}

	err = c.getAlpha()
	if err != nil {
		return nil, err
	}

	err = c.getBeta()
	if err != nil {
		return nil, err
	}

	return c, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getAlpha() error {

	var err error
	var interval float64 = 0.25

	c.Alpha, err = columnChartFloatInterval("Alpha", interval, c.msW, c.msL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getBeta() error {

	var err error
	var interval float64 = 0.25

	c.Beta, err = columnChartFloatInterval("Beta", interval, c.msW, c.msL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
