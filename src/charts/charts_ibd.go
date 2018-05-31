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
	RSRating             string
	IndustryGroupRank    string
	CompositeRating      string
	EPSRating            string
	SMRRating            string
	AccDisRating         string

	EPSChgLastQtr         string
	Last3QtrsAvgEPSGrowth string
	QtrsofEPSAcceleration string

	LastQtrEarningsSurprise string
	ThreeYrEPSGrowthRate    string
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

	err = c.getRSRating()
	if err != nil {
		return nil, err
	}

	err = c.getIndustryGroupRank()
	if err != nil {
		return nil, err
	}

	err = c.getCompositeRating()
	if err != nil {
		return nil, err
	}

	err = c.getEPSRating()
	if err != nil {
		return nil, err
	}

	err = c.getSMRRating()
	if err != nil {
		return nil, err
	}

	err = c.getAccDisRating()
	if err != nil {
		return nil, err
	}

	err = c.getEPSChgLastQtr()
	if err != nil {
		return nil, err
	}

	err = c.getLast3QtrsAvgEPSGrowth()
	if err != nil {
		return nil, err
	}

	err = c.getQtrsofEPSAcceleration()
	if err != nil {
		return nil, err
	}

	err = c.getLastQtrEarningsSurprise()
	if err != nil {
		return nil, err
	}

	err = c.getThreeYrEPSGrowthRate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

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

	var err error
	var interval float64 = 0.5

	c.UpDownVolumeRatio, err = columnChartFloatInterval("Up/Down Volume", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getRSRating() error {

	var err error
	var interval float64 = 10.0

	c.RSRating, err = columnChartIntInterval("RS Rating", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getIndustryGroupRank() error {

	var err error
	var interval float64 = 20.0

	c.IndustryGroupRank, err = columnChartIntInterval("Industry Group Rank (1 to 197)", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getCompositeRating() error {

	var err error
	var interval float64 = 10.0

	c.CompositeRating, err = columnChartIntInterval("Composite Rating", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getEPSRating() error {

	var err error
	var interval float64 = 10.0

	c.EPSRating, err = columnChartIntInterval("EPS Rating", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getSMRRating() error {

	var err error

	c.SMRRating, err = columnChartStringRank("SMR Rating", c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getAccDisRating() error {

	var err error

	c.AccDisRating, err = columnChartStringRank("Accumulation/Distribution Rating", c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getEPSChgLastQtr() error {

	var err error

	c.EPSChgLastQtr, err = columnChartPercent("EPS % Chg (Last Qtr)", c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getLast3QtrsAvgEPSGrowth() error {

	var err error

	c.Last3QtrsAvgEPSGrowth, err = columnChartPercent("Last 3 Qtrs Avg EPS Growth", c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getQtrsofEPSAcceleration() error {

	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		"# Qtrs of EPS Acceleration",
		"Winner",
		map[string]string{
			"role": "style",
		},
		"Loser",
		map[string]string{
			"role": "style",
		},
	})

	intervalDictW := make(map[int]int)
	intervalDictL := make(map[int]int)

	for _, o := range c.ibdCheckupsW {
		for _, f := range o.Contents {
			if f.Label == "# Qtrs of EPS Acceleration" {

				if f.Value == config.NullValue {
					grps := math.MaxInt32

					if val, ok := intervalDictW[grps]; ok {
						intervalDictW[grps] = val + 1
					} else {
						intervalDictW[grps] = 1
					}

					break
				} else {
					vf, err := strconv.ParseFloat(f.Value, 64)
					if err != nil {
						return err
					}

					grps := int(vf)

					if val, ok := intervalDictW[grps]; ok {
						intervalDictW[grps] = val + 1
					} else {
						intervalDictW[grps] = 1
					}

					break

				}
			}
		}
	}

	for _, o := range c.ibdCheckupsL {
		for _, f := range o.Contents {
			if f.Label == "# Qtrs of EPS Acceleration" {

				if f.Value == config.NullValue {
					grps := math.MaxInt32

					if val, ok := intervalDictL[grps]; ok {
						intervalDictL[grps] = val + 1
					} else {
						intervalDictL[grps] = 1
					}

					break
				} else {

					vf, err := strconv.ParseFloat(f.Value, 64)
					if err != nil {
						return err
					}

					grps := int(vf)

					if val, ok := intervalDictL[grps]; ok {
						intervalDictL[grps] = val + 1
					} else {
						intervalDictL[grps] = 1
					}

					break

				}
			}
		}
	}

	ck := make([]int, 0)

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

	sort.Ints(ck)

	for _, k := range ck {

		var vw int
		var vl int

		var grpk string

		if k == math.MaxInt32 {
			grpk = config.NullValue
		} else {
			grpk = fmt.Sprintf("%v", k)
		}

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

	c.QtrsofEPSAcceleration = jg

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getLastQtrEarningsSurprise() error {

	var err error

	c.LastQtrEarningsSurprise, err = columnChartPercent("Last Quarter % Earnings Surprise", c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getThreeYrEPSGrowthRate() error {

	var err error

	c.ThreeYrEPSGrowthRate, err = columnChartPercent("3 Yr EPS Growth Rate", c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
