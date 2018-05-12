package charts

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"datautils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"transaction"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ChartGeneral struct {
	filterOrders []*transaction.Order

	winners []*transaction.Order
	losers  []*transaction.Order

	GainVsDaysHeld string
	BuyPoints      string
	//BuyPointsL     string
	PriceInterval string
	//PriceIntervalL string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ChartGeneralNew(filterOrders, winners, losers []*transaction.Order) (*ChartGeneral, error) {

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

	//err = c.getBuyPointsL()
	//if err != nil {
	//return nil, err
	//}

	err = c.getPriceInterval()
	if err != nil {
		return nil, err
	}

	//err = c.getPriceIntervalL()
	//if err != nil {
	//return nil, err
	//}

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
			o.DaysHeld,
			o.GainP,
			fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
		})
	}

	for _, o := range c.losers {
		g = append(g, []interface{}{
			o.DaysHeld,
			o.GainP,
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
		"Loser",
		//map[string]string{
		//"role": "style",
		//},
	})

	dictBuyPointW := make(map[string]int)
	dictBuyPointL := make(map[string]int)

	for _, o := range c.winners {
		buyPoint := strings.TrimSpace(o.BuyPoint)

		if val, ok := dictBuyPointW[buyPoint]; ok {
			dictBuyPointW[buyPoint] = val + 1
		} else {
			dictBuyPointW[buyPoint] = 1
		}
	}

	for _, o := range c.losers {
		buyPoint := strings.TrimSpace(o.BuyPoint)

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
			vl,
			//fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
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

//func (c *ChartGeneral) getBuyPointsL() error {
//g := make([][]interface{}, 0)

//g = append(g, []interface{}{
//"Element",
//"Density",
////map[string]string{
////"role": "style",
////},
//})

//dictBuyPoint := make(map[string]int)

//for _, o := range c.losers {
//buyPoint := strings.TrimSpace(o.BuyPoint)

//if val, ok := dictBuyPoint[buyPoint]; ok {
//dictBuyPoint[buyPoint] = val + 1
//} else {
//dictBuyPoint[buyPoint] = 1
//}
//}

//for k, v := range dictBuyPoint {
//g = append(g, []interface{}{
//k,
//v,
//})
//}

//jg, err := datautils.JsonB64Encrypt(g)
//if err != nil {
//return err
//}

//c.BuyPointsL = jg

//return nil
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getPriceInterval() error {
	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		"Price Interval",
		"Winner",
		"Loser",
		//map[string]string{
		//"role": "style",
		//},
	})

	dictPriceW := make(map[int]int)
	dictPriceL := make(map[int]int)

	for _, o := range c.winners {

		grp := math.Floor(o.Price / config.PriceInterval)
		grps := int(grp * config.PriceInterval)

		if val, ok := dictPriceW[grps]; ok {
			dictPriceW[grps] = val + 1
		} else {
			dictPriceW[grps] = 1
		}
	}

	for _, o := range c.losers {

		grp := math.Floor(o.Price / config.PriceInterval)
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

	for k, _ := range dictPriceL {
		for _, c := range ck {
			if c == k {
				continue
			}
		}

		ck = append(ck, k)
	}

	for _, k := range ck {

		var vw int
		var vl int

		grps := strconv.FormatFloat(float64(k)*config.PriceInterval, 'f', -1, 64)
		grpe := strconv.FormatFloat((float64(k+1))*config.PriceInterval, 'f', -1, 64)

		grpk := fmt.Sprintf(config.PriceIntervalFormat, grps, grpe)

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
			vl,
			//fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
		})
	}

	//for k, v := range dictPrice {
	//g = append(g, []interface{}{
	//k,
	//v,
	////fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
	//})
	//}

	//dictPrice = make(map[int]int)

	//for _, o := range c.losers {

	//grp := math.Floor(o.Price / config.PriceInterval)
	//grps := int(grp * config.PriceInterval)

	//if val, ok := dictPrice[grps]; ok {
	//dictPrice[grps] = val + 1
	//} else {
	//dictPrice[grps] = 1
	//}
	//}

	//for k, v := range dictPrice {
	//g = append(g, []interface{}{
	//k,
	//v,
	//fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
	//})
	//}

	jg, err := datautils.JsonB64Encrypt(g)
	if err != nil {
		return err
	}

	c.PriceInterval = jg

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func (c *ChartGeneral) getPriceIntervalL() error {
//g := make([][]interface{}, 0)

//dictPrice := make(map[int]int)

//for _, o := range c.losers {

//grp := math.Floor(o.Price / config.PriceInterval)
//grps := int(grp * config.PriceInterval)

//if val, ok := dictPrice[grps]; ok {
//dictPrice[grps] = val + 1
//} else {
//dictPrice[grps] = 1
//}
//}

//for k, v := range dictPrice {
//g = append(g, []interface{}{
//k,
//v,
//})
//}

//jg, err := datautils.JsonB64Encrypt(g)
//if err != nil {
//return err
//}

//c.PriceIntervalL = jg

//return nil
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
