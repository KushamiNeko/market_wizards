package charts

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"datautils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"transaction"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ChartGeneral struct {
	filterOrders []*transaction.Transaction

	winners []*transaction.Transaction
	losers  []*transaction.Transaction

	GainVsDaysHeld string
	BuyPoints      string
	PriceInterval  string
	Stage          string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ChartGeneralNew(filterOrders, winners, losers []*transaction.Transaction) (*ChartGeneral, error) {

	c := new(ChartGeneral)

	c.filterOrders = filterOrders

	c.winners = winners
	c.losers = losers

	var err error

	err = c.getGainVsDaysHeld()
	if err != nil {
		return nil, err
	}

	err = c.getBuyPoints()
	if err != nil {
		return nil, err
	}

	err = c.getPriceInterval()
	if err != nil {
		return nil, err
	}

	err = c.getStage()
	if err != nil {
		return nil, err
	}

	return c, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getGainVsDaysHeld() error {

	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		"DaysHeld",
		"Gain(%)",
		map[string]string{
			"role": "style",
		},
	})

	for _, o := range c.winners {
		g = append(g, []interface{}{
			o.Sell.DaysHeld,
			o.Sell.GainP,
			fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
		})
	}

	for _, o := range c.losers {
		g = append(g, []interface{}{
			o.Sell.DaysHeld,
			o.Sell.GainP,
			fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
		})
	}

	jg, err := datautils.JsonB64Encrypt(g)
	if err != nil {
		return err
	}

	c.GainVsDaysHeld = jg

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getBuyPoints() error {
	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		"BuyPoint",
		"Winner",
		map[string]string{
			"role": "style",
		},
		"Loser",
		map[string]string{
			"role": "style",
		},
	})

	dictBuyPointW := make(map[string]int)
	dictBuyPointL := make(map[string]int)

	for _, o := range c.winners {
		buyPoint := strings.TrimSpace(o.Buy.BuyPoint)

		if val, ok := dictBuyPointW[buyPoint]; ok {
			dictBuyPointW[buyPoint] = val + 1
		} else {
			dictBuyPointW[buyPoint] = 1
		}
	}

	for _, o := range c.losers {
		buyPoint := strings.TrimSpace(o.Buy.BuyPoint)

		if val, ok := dictBuyPointL[buyPoint]; ok {
			dictBuyPointL[buyPoint] = val + 1
		} else {
			dictBuyPointL[buyPoint] = 1
		}
	}

	ck := make([]string, 0)

	for k, _ := range dictBuyPointW {
		ck = append(ck, k)
	}

outer:
	for k, _ := range dictBuyPointL {
		for _, kk := range ck {
			if kk == k {
				continue outer
			}
		}

		ck = append(ck, k)
	}

	sort.Strings(ck)

	for _, c := range ck {

		var vw int
		var vl int

		if v, ok := dictBuyPointW[c]; ok {
			vw = v
		} else {
			vw = 0
		}

		if v, ok := dictBuyPointL[c]; ok {
			vl = v
		} else {
			vl = 0
		}

		g = append(g, []interface{}{
			c,
			vw,
			fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
			vl,
			fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
		})
	}

	jg, err := datautils.JsonB64Encrypt(g)
	if err != nil {
		return err
	}

	c.BuyPoints = jg

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getPriceInterval() error {
	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		"Price Interval",
		"Winner",
		map[string]string{
			"role": "style",
		},
		"Loser",
		map[string]string{
			"role": "style",
		},
	})

	dictPriceW := make(map[int]int)
	dictPriceL := make(map[int]int)

	for _, o := range c.winners {

		grp := math.Floor(o.Buy.Price / config.PriceInterval)
		grps := int(grp * config.PriceInterval)

		if val, ok := dictPriceW[grps]; ok {
			dictPriceW[grps] = val + 1
		} else {
			dictPriceW[grps] = 1
		}
	}

	for _, o := range c.losers {

		grp := math.Floor(o.Buy.Price / config.PriceInterval)
		grps := int(grp * config.PriceInterval)

		if val, ok := dictPriceL[grps]; ok {
			dictPriceL[grps] = val + 1
		} else {
			dictPriceL[grps] = 1
		}
	}

	ck := make([]int, 0)

	for k, _ := range dictPriceW {
		ck = append(ck, k)
	}

outer:
	for k, _ := range dictPriceL {
		for _, c := range ck {
			if c == k {
				continue outer
			}
		}

		ck = append(ck, k)
	}

	sort.Ints(ck)

	for _, k := range ck {

		var vw int
		var vl int

		grp := math.Floor(float64(k) / config.PriceInterval)
		grpk := fmt.Sprintf(config.PriceIntervalFormat, int(grp*config.PriceInterval), int((grp+1)*config.PriceInterval))

		if v, ok := dictPriceW[k]; ok {
			vw = v
		} else {
			vw = 0
		}

		if v, ok := dictPriceL[k]; ok {
			vl = v
		} else {
			vl = 0
		}

		g = append(g, []interface{}{
			grpk,
			vw,
			fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
			vl,
			fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
		})
	}

	jg, err := datautils.JsonB64Encrypt(g)
	if err != nil {
		return err
	}

	c.PriceInterval = jg

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getStage() error {

	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		"Stage",
		"Winner",
		map[string]string{
			"role": "style",
		},
		"Loser",
		map[string]string{
			"role": "style",
		},
	})

	dictStageW := make(map[string]int)
	dictStageL := make(map[string]int)

	for _, o := range c.winners {
		stages := strconv.FormatFloat(math.Floor(o.Buy.Stage), 'f', -1, 64)

		if val, ok := dictStageW[stages]; ok {
			dictStageW[stages] = val + 1
		} else {
			dictStageW[stages] = 1
		}
	}

	for _, o := range c.losers {
		stages := strconv.FormatFloat(math.Floor(o.Buy.Stage), 'f', -1, 64)

		if val, ok := dictStageL[stages]; ok {
			dictStageL[stages] = val + 1
		} else {
			dictStageL[stages] = 1
		}
	}

	ck := make([]string, 0)

	for k, _ := range dictStageW {
		ck = append(ck, k)
	}

outer:
	for k, _ := range dictStageL {
		for _, kk := range ck {
			if kk == k {
				continue outer
			}
		}

		ck = append(ck, k)
	}

	sort.Strings(ck)

	for _, c := range ck {

		var vw int
		var vl int

		if v, ok := dictStageW[c]; ok {
			vw = v
		} else {
			vw = 0
		}

		if v, ok := dictStageL[c]; ok {
			vl = v
		} else {
			vl = 0
		}

		g = append(g, []interface{}{
			c,
			vw,
			fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
			vl,
			fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
		})
	}

	jg, err := datautils.JsonB64Encrypt(g)
	if err != nil {
		return err
	}

	c.Stage = jg

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
