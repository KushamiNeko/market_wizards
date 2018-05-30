package charts

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"config"
	"datautils"
	"encoding/json"
	"fmt"
	"ibd"
	"strconv"
	"strings"
	"transaction"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ChartIBD struct {
	filterOrders []*transaction.Order

	winnersIBD []*bytes.Buffer
	losersIBD  []*bytes.Buffer

	ibdCheckupsW []*ibd.IBDCheckup
	ibdCheckupsL []*ibd.IBDCheckup

	MarketCapitalization string
	//BuyPoints      string
	//PriceInterval  string
	//Stage          string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ChartIBDNew(filterOrders []*transaction.Order, winnersIBD, losersIBD []*bytes.Buffer) (*ChartIBD, error) {

	c := new(ChartIBD)

	c.filterOrders = filterOrders

	c.winnersIBD = winnersIBD
	c.losersIBD = losersIBD

	c.ibdCheckupsW = make([]*ibd.IBDCheckup, len(c.winnersIBD))
	c.ibdCheckupsL = make([]*ibd.IBDCheckup, len(c.losersIBD))

	var err error

	for i, w := range c.winnersIBD {

		checkup := ibd.IBDCheckupNew()
		err = json.Unmarshal(w.Bytes(), checkup)
		if err != nil {
			return nil, err
		}

		c.ibdCheckupsW[i] = checkup
	}

	for i, l := range c.losersIBD {

		checkup := ibd.IBDCheckupNew()
		err = json.Unmarshal(l.Bytes(), checkup)
		if err != nil {
			return nil, err
		}

		c.ibdCheckupsL[i] = checkup
	}

	//fmt.Println(c.ibdCheckupsW[0])
	//fmt.Println(c.ibdCheckupsL[0])

	err = c.getMarketCapitalization()
	if err != nil {
		return nil, err
	}

	//err = c.getBuyPoints()
	//if err != nil {
	//return nil, err
	//}

	//err = c.getPriceInterval()
	//if err != nil {
	//return nil, err
	//}

	//err = c.getStage()
	//if err != nil {
	//return nil, err
	//}

	return c, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func (c *ChartIBD) getMarketCapitalization() error {

//g := make([][]interface{}, 0)

//g = append(g, []interface{}{
//"DaysHeld",
//"Gain(%)",
//map[string]string{
//"role": "style",
//},
//})

//for _, o := range c.winners {
//g = append(g, []interface{}{
//o.DaysHeld,
//o.GainP,
//fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
//})
//}

//for _, o := range c.losers {
//g = append(g, []interface{}{
//o.DaysHeld,
//o.GainP,
//fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
//})
//}

//jg, err := datautils.JsonB64Encrypt(g)
//if err != nil {
//return err
//}

//c.GainVsDaysHeld = jg

//return nil
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getMarketCapitalization() error {
	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		"Market Capitalization",
		"Winner",
		map[string]string{
			"role": "style",
		},
		"Loser",
		map[string]string{
			"role": "style",
		},
	})

	var smallCapThreshold int64 = 1000000000
	var largeCapThreshold int64 = 10000000000

	smallCapW := 0
	midCapW := 0
	largeCapW := 0

	smallCapL := 0
	midCapL := 0
	largeCapL := 0

	for _, o := range c.ibdCheckupsW {
		for _, f := range o.Contents {
			if f.Label == "Market Capitalization" {

				//fmt.Println(f.Value)

				v := strings.Replace(f.Value, "$", "", -1)
				vi, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}

				if vi <= smallCapThreshold {
					smallCapW += 1
				} else if vi <= largeCapThreshold {
					midCapW += 1
				} else if vi > largeCapThreshold {
					largeCapW += 1
				}

				break
			}
		}
	}

	for _, o := range c.ibdCheckupsL {
		for _, f := range o.Contents {
			if f.Label == "Market Capitalization" {

				//fmt.Println(f.Value)

				v := strings.Replace(f.Value, "$", "", -1)
				vi, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}

				if vi <= smallCapThreshold {
					smallCapL += 1
				} else if vi <= largeCapThreshold {
					midCapL += 1
				} else if vi > largeCapThreshold {
					largeCapL += 1
				}

				break
			}
		}
	}

	g = append(g, []interface{}{
		"Small Cap",
		smallCapW,
		fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
		smallCapL,
		fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
	})

	g = append(g, []interface{}{
		"Mid Cap",
		midCapW,
		fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
		midCapL,
		fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
	})

	g = append(g, []interface{}{
		"Large Cap",
		largeCapW,
		fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
		largeCapL,
		fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
	})

	jg, err := datautils.JsonB64Encrypt(g)
	if err != nil {
		return err
	}

	c.MarketCapitalization = jg

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func (c *ChartGeneral) getPriceInterval() error {
//g := make([][]interface{}, 0)

//g = append(g, []interface{}{
//"Price Interval",
//"Winner",
//map[string]string{
//"role": "style",
//},
//"Loser",
//map[string]string{
//"role": "style",
//},
//})

//dictPriceW := make(map[int]int)
//dictPriceL := make(map[int]int)

//for _, o := range c.winners {

//grp := math.Floor(o.Price / config.PriceInterval)
//grps := int(grp * config.PriceInterval)

//if val, ok := dictPriceW[grps]; ok {
//dictPriceW[grps] = val + 1
//} else {
//dictPriceW[grps] = 1
//}
//}

//for _, o := range c.losers {

//grp := math.Floor(o.Price / config.PriceInterval)
//grps := int(grp * config.PriceInterval)

//if val, ok := dictPriceL[grps]; ok {
//dictPriceL[grps] = val + 1
//} else {
//dictPriceL[grps] = 1
//}
//}

//ck := make([]int, 0)

//for k, _ := range dictPriceW {
//ck = append(ck, k)
//}

//outer:
//for k, _ := range dictPriceL {
//for _, c := range ck {
//if c == k {
//continue outer
//}
//}

//ck = append(ck, k)
//}

//for _, k := range ck {

//var vw int
//var vl int

//grp := math.Floor(float64(k) / config.PriceInterval)
//grpk := fmt.Sprintf(config.PriceIntervalFormat, int(grp*config.PriceInterval), int((grp+1)*config.PriceInterval))

//if v, ok := dictPriceW[k]; ok {
//vw = v
//} else {
//vw = 0
//}

//if v, ok := dictPriceL[k]; ok {
//vl = v
//} else {
//vl = 0
//}

//g = append(g, []interface{}{
//grpk,
//vw,
//fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
//vl,
//fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
//})
//}

//jg, err := datautils.JsonB64Encrypt(g)
//if err != nil {
//return err
//}

//c.PriceInterval = jg

//return nil
//}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func (c *ChartGeneral) getStage() error {

//g := make([][]interface{}, 0)

//g = append(g, []interface{}{
//"Stage",
//"Winner",
//map[string]string{
//"role": "style",
//},
//"Loser",
//map[string]string{
//"role": "style",
//},
//})

//dictStageW := make(map[string]int)
//dictStageL := make(map[string]int)

//for _, o := range c.winners {
////g = append(g, []interface{}{
////o.DaysHeld,
////o.GainP,
////fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
////})

//stages := strconv.FormatFloat(math.Floor(o.Stage), 'f', -1, 64)

//if val, ok := dictStageW[stages]; ok {
//dictStageW[stages] = val + 1
//} else {
//dictStageW[stages] = 1
//}
//}

//for _, o := range c.losers {
////g = append(g, []interface{}{
////o.DaysHeld,
////o.GainP,
////fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
////})

//stages := strconv.FormatFloat(math.Floor(o.Stage), 'f', -1, 64)

//if val, ok := dictStageL[stages]; ok {
//dictStageL[stages] = val + 1
//} else {
//dictStageL[stages] = 1
//}
//}

//ck := make([]string, 0)

//for k, _ := range dictStageW {
//ck = append(ck, k)
//}

//outer:
//for k, _ := range dictStageL {
//for _, kk := range ck {
//if kk == k {
//continue outer
//}
//}

//ck = append(ck, k)
//}

//for _, c := range ck {

//var vw int
//var vl int

//if v, ok := dictStageW[c]; ok {
//vw = v
//} else {
//vw = 0
//}

//if v, ok := dictStageL[c]; ok {
//vl = v
//} else {
//vl = 0
//}

//g = append(g, []interface{}{
//c,
//vw,
//fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
//vl,
//fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
//})
//}

//jg, err := datautils.JsonB64Encrypt(g)
//if err != nil {
//return err
//}

//c.Stage = jg

//return nil
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
