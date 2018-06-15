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
	filterOrders []*transaction.Transaction

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

	EPSChgLastQtr           string
	Last3QtrsAvgEPSGrowth   string
	QtrsofEPSAcceleration   string
	EPSEstChgCurrentQtr     string
	EstimateRevisions       string
	LastQtrEarningsSurprise string

	ThreeYrEPSGrowthRate            string
	ConsecutiveYrsofAnnualEPSGrowth string
	EPSEstChgforCurrentYear         string

	SalesChgLastQtr        string
	ThreeYrSalesGrowthRate string

	AnnualPreTaxMargin string
	AnnualROE          string
	DebtEquityRatio    string

	Off52WeekHigh             string
	Pricevs50DayMovingAverage string
	FiftyDayAverageVolume     string

	ChangeInFundsOwningStock      string
	QtrsOfIncreasingFundOwnership string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ChartIBDNew(filterOrders []*transaction.Transaction, winnersIBD, losersIBD []*bytes.Buffer) (*ChartIBD, error) {

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

	err = c.getEPSEstChgCurrentQtr()
	if err != nil {
		return nil, err
	}

	err = c.getEstimateRevisions()
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

	err = c.getConsecutiveYrsofAnnualEPSGrowth()
	if err != nil {
		return nil, err
	}

	err = c.getEPSEstChgforCurrentYear()
	if err != nil {
		return nil, err
	}

	err = c.getSalesChgLastQtr()
	if err != nil {
		return nil, err
	}

	err = c.getThreeYrSalesGrowthRate()
	if err != nil {
		return nil, err
	}

	err = c.getAnnualPreTaxMargin()
	if err != nil {
		return nil, err
	}

	err = c.getAnnualROE()
	if err != nil {
		return nil, err
	}

	err = c.getDebtEquityRatio()
	if err != nil {
		return nil, err
	}

	err = c.getOff52WeekHigh()
	if err != nil {
		return nil, err
	}

	err = c.getPricevs50DayMovingAverage()
	if err != nil {
		return nil, err
	}

	err = c.getFiftyDayAverageVolume()
	if err != nil {
		return nil, err
	}

	err = c.getChangeInFundsOwningStock()
	if err != nil {
		return nil, err
	}

	err = c.getQtrsOfIncreasingFundOwnership()
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
	var interval float64 = 20.0

	c.EPSChgLastQtr, err = columnChartPercent("EPS % Chg (Last Qtr)", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getLast3QtrsAvgEPSGrowth() error {

	var err error
	var interval float64 = 20.0

	c.Last3QtrsAvgEPSGrowth, err = columnChartPercent("Last 3 Qtrs Avg EPS Growth", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getQtrsofEPSAcceleration() error {

	var err error

	c.QtrsofEPSAcceleration, err = columnChartString("# Qtrs of EPS Acceleration", c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getEPSEstChgCurrentQtr() error {

	var err error
	var interval float64 = 20.0

	c.EPSEstChgCurrentQtr, err = columnChartPercent("EPS Est % Chg (Current Qtr)", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getEstimateRevisions() error {

	var err error

	c.EstimateRevisions, err = columnChartString("Estimate Revisions", c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getLastQtrEarningsSurprise() error {

	var err error
	var interval float64 = 20.0

	c.LastQtrEarningsSurprise, err = columnChartPercent("Last Quarter % Earnings Surprise", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getThreeYrEPSGrowthRate() error {

	var err error
	var interval float64 = 20.0

	c.ThreeYrEPSGrowthRate, err = columnChartPercent("3 Yr EPS Growth Rate", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getConsecutiveYrsofAnnualEPSGrowth() error {

	var err error

	c.ConsecutiveYrsofAnnualEPSGrowth, err = columnChartString("Consecutive Yrs of Annual EPS Growth", c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getEPSEstChgforCurrentYear() error {

	var err error
	var interval float64 = 20.0

	c.EPSEstChgforCurrentYear, err = columnChartPercent("EPS Est % Chg for Current Year", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getSalesChgLastQtr() error {

	var err error
	var interval float64 = 20.0

	c.SalesChgLastQtr, err = columnChartPercent("Sales % Chg (Last Qtr)", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getThreeYrSalesGrowthRate() error {

	var err error
	var interval float64 = 20.0

	c.ThreeYrSalesGrowthRate, err = columnChartPercent("3 Yr Sales Growth Rate", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getAnnualPreTaxMargin() error {

	var err error
	var interval float64 = 5.0

	c.AnnualPreTaxMargin, err = columnChartPercent("Annual Pre-Tax Margin", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getAnnualROE() error {

	var err error
	var interval float64 = 5.0

	c.AnnualROE, err = columnChartPercent("Annual ROE", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getDebtEquityRatio() error {

	var err error
	var interval float64 = 5.0

	c.DebtEquityRatio, err = columnChartPercent("Debt/Equity Ratio", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getOff52WeekHigh() error {

	var err error
	var interval float64 = 5.0

	c.Off52WeekHigh, err = columnChartPercent("% Off 52 Week High", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getPricevs50DayMovingAverage() error {

	var err error
	var interval float64 = 5.0

	c.Pricevs50DayMovingAverage, err = columnChartPercent("Price vs. 50-Day Moving Average", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getFiftyDayAverageVolume() error {

	var err error
	var interval float64 = 200000.0

	c.FiftyDayAverageVolume, err = columnChartIntInterval("50-Day Average Volume", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getChangeInFundsOwningStock() error {

	var err error
	var interval float64 = 5.0

	c.ChangeInFundsOwningStock, err = columnChartPercent("% Change In Funds Owning Stock", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getQtrsOfIncreasingFundOwnership() error {

	var err error

	c.QtrsOfIncreasingFundOwnership, err = columnChartString("Qtrs Of Increasing Fund Ownership", c.ibdCheckupsW, c.ibdCheckupsL)
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
