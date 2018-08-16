package charts

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"config"
	"encoding/json"
	"fmt"
	"html/template"
	"ibd"
	"sort"
	"strconv"
	"strings"
	"sync"
	"transaction"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ChartIBD struct {
	filterOrders []*transaction.Transaction

	winnersIBD []*bytes.Buffer
	losersIBD  []*bytes.Buffer

	winnersI []interface{}
	losersI  []interface{}

	//ibdCheckupsW []*ibd.IBDCheckup
	//ibdCheckupsL []*ibd.IBDCheckup

	//ibdCheckupsW []datautils.Contents
	//ibdCheckupsL []datautils.Contents

	MarketCapitalization template.URL
	UpDownVolumeRatio    template.URL

	RSRating          template.URL
	IndustryGroupRank template.URL

	CompositeRating template.URL
	EPSRating       template.URL
	SMRRating       template.URL
	AccDisRating    template.URL

	EPSChgLastQtr           template.URL
	Last3QtrsAvgEPSGrowth   template.URL
	QtrsofEPSAcceleration   template.URL
	EPSEstChgCurrentQtr     template.URL
	EstimateRevisions       template.URL
	LastQtrEarningsSurprise template.URL

	ThreeYrEPSGrowthRate            template.URL
	ConsecutiveYrsofAnnualEPSGrowth template.URL
	EPSEstChgforCurrentYear         template.URL

	SalesChgLastQtr        template.URL
	ThreeYrSalesGrowthRate template.URL

	AnnualPreTaxMargin template.URL
	AnnualROE          template.URL
	DebtEquityRatio    template.URL

	Off52WeekHigh             template.URL
	Pricevs50DayMovingAverage template.URL
	FiftyDayAverageVolume     template.URL

	ChangeInFundsOwningStock      template.URL
	QtrsOfIncreasingFundOwnership template.URL
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ChartIBDNew(filterOrders []*transaction.Transaction, winnersIBD, losersIBD []*bytes.Buffer) (*ChartIBD, error) {

	var wg sync.WaitGroup

	c := new(ChartIBD)

	c.filterOrders = filterOrders

	c.winnersIBD = winnersIBD
	c.losersIBD = losersIBD

	//c.ibdCheckupsW = make([]*ibd.IBDCheckup, len(c.winnersIBD))
	//c.ibdCheckupsL = make([]*ibd.IBDCheckup, len(c.losersIBD))

	c.winnersI = make([]interface{}, len(c.winnersIBD))
	c.losersI = make([]interface{}, len(c.losersIBD))

	//c.ibdCheckupsW = make([]datautils.Contents, len(c.winnersIBD))
	//c.ibdCheckupsL = make([]datautils.Contents, len(c.losersIBD))

	var err error

	for i, w := range c.winnersIBD {

		checkup := ibd.IBDCheckupNew()
		err = json.Unmarshal(w.Bytes(), checkup)
		if err != nil {
			return nil, err
		}

		//c.ibdCheckupsW[i] = checkup

		c.winnersI[i] = checkup
	}

	for i, l := range c.losersIBD {

		checkup := ibd.IBDCheckupNew()
		err = json.Unmarshal(l.Bytes(), checkup)
		if err != nil {
			return nil, err
		}

		//c.ibdCheckupsL[i] = checkup

		c.losersI[i] = checkup
	}

	fs := 27

	errs := make([]error, fs)
	wg.Add(fs)

	go func() {
		err := c.getMarketCapitalization()
		if err != nil {
			//return nil, err
			errs[0] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getUpDownVolumeRatio()
		if err != nil {
			//return nil, err
			errs[1] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getRSRating()
		if err != nil {
			//return nil, err
			errs[2] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getIndustryGroupRank()
		if err != nil {
			//return nil, err
			errs[3] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getCompositeRating()
		if err != nil {
			//return nil, err
			errs[4] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getEPSRating()
		if err != nil {
			//return nil, err
			errs[5] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getSMRRating()
		if err != nil {
			//return nil, err
			errs[6] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getAccDisRating()
		if err != nil {
			//return nil, err
			errs[7] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getEPSChgLastQtr()
		if err != nil {
			//return nil, err
			errs[8] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getLast3QtrsAvgEPSGrowth()
		if err != nil {
			//return nil, err
			errs[9] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getQtrsofEPSAcceleration()
		if err != nil {
			//return nil, err
			errs[10] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getEPSEstChgCurrentQtr()
		if err != nil {
			//return nil, err
			errs[11] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getEstimateRevisions()
		if err != nil {
			//return nil, err
			errs[12] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getLastQtrEarningsSurprise()
		if err != nil {
			//return nil, err
			errs[13] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getThreeYrEPSGrowthRate()
		if err != nil {
			//return nil, err
			errs[14] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getConsecutiveYrsofAnnualEPSGrowth()
		if err != nil {
			//return nil, err
			errs[15] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getEPSEstChgforCurrentYear()
		if err != nil {
			//return nil, err
			errs[16] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getSalesChgLastQtr()
		if err != nil {
			//return nil, err
			errs[17] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getThreeYrSalesGrowthRate()
		if err != nil {
			//return nil, err
			errs[18] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getAnnualPreTaxMargin()
		if err != nil {
			//return nil, err
			errs[19] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getAnnualROE()
		if err != nil {
			//return nil, err
			errs[20] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getDebtEquityRatio()
		if err != nil {
			//return nil, err
			errs[21] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getOff52WeekHigh()
		if err != nil {
			//return nil, err
			errs[22] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getPricevs50DayMovingAverage()
		if err != nil {
			//return nil, err
			errs[23] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getFiftyDayAverageVolume()
		if err != nil {
			//return nil, err
			errs[24] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getChangeInFundsOwningStock()
		if err != nil {
			//return nil, err
			errs[25] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getQtrsOfIncreasingFundOwnership()
		if err != nil {
			//return nil, err
			errs[26] = err
		}

		wg.Done()
	}()

	wg.Wait()

	for _, e := range errs {
		if e != nil {
			return nil, e
		}
	}

	return c, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getMarketCapitalization() error {

	label := "Market Capitalization"

	p, err := makePlot(
		label,
		"",
		"",
		true,
		func(p *plot.Plot) {
			p.X.Padding = vg.Points(config.ChartXLabelPadding)
			//p.X.Tick.Label.Rotation = config.ChartLabelRotation

			p.X.Tick.Label.XAlign = draw.XLeft
			p.X.Tick.Label.YAlign = draw.YCenter

			p.Y.Tick.Label.XAlign = draw.XRight
		},
	)
	if err != nil {
		return err
	}

	keys, winnersG, losersG, wmax, lmax, err :=
		makeValueSlice(
			c.winnersI,
			c.losersI,
			func(o interface{}) (interface{}, error) {

				var smallCapThreshold int64 = 1000000000
				var largeCapThreshold int64 = 10000000000

				t := o.(*ibd.IBDCheckup)

				v := strings.Replace(t.Contents[label], "$", "", -1)
				vi, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return "", err
				}

				if vi <= smallCapThreshold {
					return "Small Cap", nil
				} else if vi <= largeCapThreshold {
					return "Mid Cap", nil
				} else if vi > largeCapThreshold {
					return "Large Cap", nil
				}

				return "", fmt.Errorf("Uncaptured Error\n")
			},
			func(keys []interface{}) {
				sort.Slice(keys, func(i, j int) bool {
					is := keys[i].(string)
					js := keys[j].(string)

					return is > js
				})
			},
			func(key interface{}) string {
				return key.(string)
			},
		)
	if err != nil {
		return err
	}

	err = makeBarCharts(p, keys, winnersG, losersG, wmax, lmax)
	if err != nil {
		return err
	}

	c.MarketCapitalization, err = plotToDataUrl(p)
	if err != nil {
		return err
	}

	return nil

	//g := make([][]interface{}, 0)

	//g = append(g, []interface{}{
	//"Market Capitalization",
	//"Winner",
	//map[string]string{
	//"role": "style",
	//},
	//"Loser",
	//map[string]string{
	//"role": "style",
	//},
	//})

	//var smallCapThreshold int64 = 1000000000
	//var largeCapThreshold int64 = 10000000000

	//smallCapW := 0
	//midCapW := 0
	//largeCapW := 0

	//smallCapL := 0
	//midCapL := 0
	//largeCapL := 0

	//for _, o := range c.ibdCheckupsW {
	//for _, f := range o.GetContents() {
	//if f.GetLabel() == "Market Capitalization" {

	//v := strings.Replace(f.GetValue(), "$", "", -1)
	//vi, err := strconv.ParseInt(v, 10, 64)
	//if err != nil {
	//return err
	//}

	//if vi <= smallCapThreshold {
	//smallCapW += 1
	//} else if vi <= largeCapThreshold {
	//midCapW += 1
	//} else if vi > largeCapThreshold {
	//largeCapW += 1
	//}

	//break
	//}
	//}
	//}

	//for _, o := range c.ibdCheckupsL {
	//for _, f := range o.GetContents() {
	//if f.GetLabel() == "Market Capitalization" {

	//v := strings.Replace(f.GetValue(), "$", "", -1)
	//vi, err := strconv.ParseInt(v, 10, 64)
	//if err != nil {
	//return err
	//}

	//if vi <= smallCapThreshold {
	//smallCapL += 1
	//} else if vi <= largeCapThreshold {
	//midCapL += 1
	//} else if vi > largeCapThreshold {
	//largeCapL += 1
	//}

	//break
	//}
	//}
	//}

	//g = append(g, []interface{}{
	//"Small Cap",
	//smallCapW,
	//fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
	//smallCapL,
	//fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
	//})

	//g = append(g, []interface{}{
	//"Mid Cap",
	//midCapW,
	//fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
	//midCapL,
	//fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
	//})

	//g = append(g, []interface{}{
	//"Large Cap",
	//largeCapW,
	//fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
	//largeCapL,
	//fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
	//})

	//jg, err := datautils.JsonB64Encrypt(g)
	//if err != nil {
	//return err
	//}

	//c.MarketCapitalization = jg

	//return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getUpDownVolumeRatio() error {

	var err error
	var interval float64 = 0.5

	c.UpDownVolumeRatio, err = barChartFloatInterval("Up/Down Volume", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	//p, err := makePlot(
	//"Up/Down Volume",
	//"",
	//"",
	//true,
	//func(p *plot.Plot) {
	//p.X.Padding = vg.Points(config.ChartXLabelPadding)
	//p.X.Tick.Label.Rotation = config.ChartLabelRotation

	//p.X.Tick.Label.XAlign = draw.XLeft
	//p.X.Tick.Label.YAlign = draw.YCenter

	//p.Y.Tick.Label.XAlign = draw.XRight
	//},
	//)
	//if err != nil {
	//return err
	//}

	//var interval float64 = 0.5

	//keys, winnersG, losersG, wmax, lmax, err :=
	//makeValueSlice(
	//c.winnersI,
	//c.losersI,
	//func(o interface{}) (interface{}, error) {

	//t := o.(*ibd.IBDCheckup)
	//value := t.Contents["Up/Down Volume"]

	//var grps float64

	//if value == config.NullValue {
	//grps = math.MaxFloat64
	//} else {

	//vf, err := strconv.ParseFloat(value, 64)
	//if err != nil {
	//return "", err
	//}

	//grp := math.Floor(vf / interval)
	//grps = float64(grp * interval)

	//}

	//return grps, nil
	//},
	//func(keys []interface{}) {
	//sort.Slice(keys, func(i, j int) bool {
	//is := keys[i].(float64)
	//js := keys[j].(float64)

	//return is < js
	//})
	//},
	//func(key interface{}) string {

	//k := key.(float64)

	//p := message.NewPrinter(message.MatchLanguage("en"))
	//var grpk string

	//if k == math.MaxFloat64 {
	//grpk = config.NullValue
	//} else {
	//grp := math.Floor(k / interval)
	//grpk = p.Sprintf(config.IntervalFormat, grp*interval, (grp+1)*interval)
	//}

	//return grpk
	//},
	//)
	//if err != nil {
	//return err
	//}

	//err = makeBarCharts(p, keys, winnersG, losersG, wmax, lmax)
	//if err != nil {
	//return err
	//}

	//c.UpDownVolumeRatio, err = plotToDataUrl(p)
	//if err != nil {
	//return err
	//}

	//return nil

	//g := make([][]interface{}, 0)

	//g = append(g, []interface{}{
	//label,
	//"Winner",
	//map[string]string{
	//"role": "style",
	//},
	//"Loser",
	//map[string]string{
	//"role": "style",
	//},
	//})

	//intervalDictW := make(map[float64]int)
	//intervalDictL := make(map[float64]int)

	//for _, o := range winners {
	//for _, f := range o.GetContents() {
	//if f.GetLabel() == label {
	//var grps float64

	//if f.GetValue() == config.NullValue {
	//grps = math.MaxFloat64
	//} else {

	//vf, err := strconv.ParseFloat(f.GetValue(), 64)
	//if err != nil {
	//return "", err
	//}

	//grp := math.Floor(vf / interval)
	//grps = float64(grp * interval)
	//}

	//if val, ok := intervalDictW[grps]; ok {
	//intervalDictW[grps] = val + 1
	//} else {
	//intervalDictW[grps] = 1
	//}

	//break
	//}
	//}
	//}

	//for _, o := range losers {
	//for _, f := range o.GetContents() {
	//if f.GetLabel() == label {
	//var grps float64

	//if f.GetValue() == config.NullValue {
	//grps = math.MaxFloat64
	//} else {

	//vf, err := strconv.ParseFloat(f.GetValue(), 64)
	//if err != nil {
	//return "", err
	//}

	//grp := math.Floor(vf / interval)
	//grps = float64(grp * interval)
	//}

	//if val, ok := intervalDictL[grps]; ok {
	//intervalDictL[grps] = val + 1
	//} else {
	//intervalDictL[grps] = 1
	//}

	//break
	//}
	//}
	//}

	//ck := make([]float64, 0)

	//for k, _ := range intervalDictW {
	//ck = append(ck, k)
	//}

	//outer:
	//for k, _ := range intervalDictL {
	//for _, c := range ck {
	//if c == k {
	//continue outer
	//}
	//}

	//ck = append(ck, k)
	//}

	//sort.Float64s(ck)

	//p := message.NewPrinter(message.MatchLanguage("en"))

	//for _, k := range ck {

	//var vw int
	//var vl int

	//var grpk string

	//if k == math.MaxFloat64 {
	//grpk = config.NullValue
	//} else {
	//grp := math.Floor(float64(k) / interval)
	////grpk = fmt.Sprintf(config.IntervalFormat, grp*interval, (grp+1)*interval)
	//grpk = p.Sprintf(config.IntervalFormat, grp*interval, (grp+1)*interval)
	//}

	//if v, ok := intervalDictW[k]; ok {
	//vw = v
	//} else {
	//vw = 0
	//}

	//if v, ok := intervalDictL[k]; ok {
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
	//return "", err
	//}

	//return jg, nil

	//var err error
	//var interval float64 = 0.5

	//c.UpDownVolumeRatio, err = columnChartFloatInterval("Up/Down Volume", interval, c.ibdCheckupsW, c.ibdCheckupsL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getRSRating() error {

	var err error
	var interval float64 = 10.0

	c.RSRating, err = barChartIntInterval("RS Rating", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getIndustryGroupRank() error {

	var err error
	var interval float64 = 20.0

	c.IndustryGroupRank, err = barChartIntInterval(
		"Industry Group Rank (1 to 197)",
		interval,
		c.winnersI,
		c.losersI,
	)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getCompositeRating() error {

	var err error
	var interval float64 = 10.0

	c.CompositeRating, err = barChartIntInterval("Composite Rating", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getEPSRating() error {

	var err error
	var interval float64 = 10.0

	c.EPSRating, err = barChartIntInterval("EPS Rating", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getSMRRating() error {

	var err error

	c.SMRRating, err = barChartStringRank("SMR Rating", c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getAccDisRating() error {

	var err error

	c.AccDisRating, err = barChartStringRank("Accumulation/Distribution Rating", c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getEPSChgLastQtr() error {

	var err error
	var interval float64 = 20.0

	c.EPSChgLastQtr, err = barChartPercent("EPS % Chg (Last Qtr)", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getLast3QtrsAvgEPSGrowth() error {

	var err error
	var interval float64 = 20.0

	c.Last3QtrsAvgEPSGrowth, err = barChartPercent("Last 3 Qtrs Avg EPS Growth", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getQtrsofEPSAcceleration() error {

	var err error

	c.QtrsofEPSAcceleration, err = barChartString("# Qtrs of EPS Acceleration", c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getEPSEstChgCurrentQtr() error {

	var err error
	var interval float64 = 20.0

	c.EPSEstChgCurrentQtr, err = barChartPercent("EPS Est % Chg (Current Qtr)", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getEstimateRevisions() error {

	var err error

	c.EstimateRevisions, err = barChartString("Estimate Revisions", c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getLastQtrEarningsSurprise() error {

	var err error
	var interval float64 = 20.0

	c.LastQtrEarningsSurprise, err = barChartPercent("Last Quarter % Earnings Surprise", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getThreeYrEPSGrowthRate() error {

	var err error
	var interval float64 = 20.0

	c.ThreeYrEPSGrowthRate, err = barChartPercent("3 Yr EPS Growth Rate", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getConsecutiveYrsofAnnualEPSGrowth() error {

	var err error

	c.ConsecutiveYrsofAnnualEPSGrowth, err = barChartString("Consecutive Yrs of Annual EPS Growth", c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getEPSEstChgforCurrentYear() error {

	var err error
	var interval float64 = 20.0

	c.EPSEstChgforCurrentYear, err = barChartPercent("EPS Est % Chg for Current Year", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getSalesChgLastQtr() error {

	var err error
	var interval float64 = 20.0

	c.SalesChgLastQtr, err = barChartPercent("Sales % Chg (Last Qtr)", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getThreeYrSalesGrowthRate() error {

	var err error
	var interval float64 = 20.0

	c.ThreeYrSalesGrowthRate, err = barChartPercent("3 Yr Sales Growth Rate", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getAnnualPreTaxMargin() error {

	var err error
	var interval float64 = 5.0

	c.AnnualPreTaxMargin, err = barChartPercent("Annual Pre-Tax Margin", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getAnnualROE() error {

	var err error
	var interval float64 = 5.0

	c.AnnualROE, err = barChartPercent("Annual ROE", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getDebtEquityRatio() error {

	var err error
	var interval float64 = 5.0

	c.DebtEquityRatio, err = barChartPercent("Debt/Equity Ratio", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getOff52WeekHigh() error {

	var err error
	var interval float64 = 5.0

	c.Off52WeekHigh, err = barChartPercent("% Off 52 Week High", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getPricevs50DayMovingAverage() error {

	var err error
	var interval float64 = 5.0

	c.Pricevs50DayMovingAverage, err = barChartPercent("Price vs. 50-Day Moving Average", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getFiftyDayAverageVolume() error {

	var err error
	var interval float64 = 200000.0

	c.FiftyDayAverageVolume, err = barChartIntInterval("50-Day Average Volume", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getChangeInFundsOwningStock() error {

	var err error
	var interval float64 = 5.0

	c.ChangeInFundsOwningStock, err = barChartPercent("% Change In Funds Owning Stock", interval, c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartIBD) getQtrsOfIncreasingFundOwnership() error {

	var err error

	c.QtrsOfIncreasingFundOwnership, err = barChartString("Qtrs Of Increasing Fund Ownership", c.winnersI, c.losersI)
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
