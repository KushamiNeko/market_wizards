package charts

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"config"
	"datautils"
	"encoding/json"
	"fmt"
	"ibd"
	"math"
	"sort"
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
	UpDownVolumeRatio    string
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

	err = c.getMarketCapitalization()
	if err != nil {
		return nil, err
	}

	err = c.getUpDownVolumeRatio()
	if err != nil {
		return nil, err
	}

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

func (c *ChartIBD) getUpDownVolumeRatio() error {
	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		"Up/Down Volume Ratio",
		"Winner",
		map[string]string{
			"role": "style",
		},
		"Loser",
		map[string]string{
			"role": "style",
		},
	})

	var interval float64 = 0.5

	intervalDictW := make(map[float64]int)
	intervalDictL := make(map[float64]int)

	for _, o := range c.ibdCheckupsW {
		for _, f := range o.Contents {
			if f.Label == "Up/Down Volume" {
				vf, err := strconv.ParseFloat(f.Value, 64)
				if err != nil {
					return err
				}

				grp := math.Floor(vf / interval)
				grps := float64(grp * interval)

				if val, ok := intervalDictW[grps]; ok {
					intervalDictW[grps] = val + 1
				} else {
					intervalDictW[grps] = 1
				}

				break
			}
		}
	}

	for _, o := range c.ibdCheckupsL {
		for _, f := range o.Contents {
			if f.Label == "Up/Down Volume" {
				vf, err := strconv.ParseFloat(f.Value, 64)
				if err != nil {
					return err
				}

				grp := math.Floor(vf / interval)
				grps := float64(grp * interval)

				if val, ok := intervalDictL[grps]; ok {
					intervalDictL[grps] = val + 1
				} else {
					intervalDictL[grps] = 1
				}

				break
			}
		}
	}

	ck := make([]float64, 0)

	for k, _ := range intervalDictW {
		ck = append(ck, k)
	}

outer:
	for k, _ := range intervalDictL {
		for _, c := range ck {
			if c == k {
				continue outer
			}
		}

		ck = append(ck, k)
	}

	sort.Float64s(ck)

	for _, k := range ck {

		var vw int
		var vl int

		grp := math.Floor(float64(k) / interval)
		grpk := fmt.Sprintf(config.PriceIntervalFormat, grp*interval, (grp+1)*interval)

		if v, ok := intervalDictW[k]; ok {
			vw = v
		} else {
			vw = 0
		}

		if v, ok := intervalDictL[k]; ok {
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

	c.UpDownVolumeRatio = jg

	return nil
}

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
