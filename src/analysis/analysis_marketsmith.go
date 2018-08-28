package analysis

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"transaction"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ChartMarketSmith struct {
	filterOrders []*transaction.Trade

	winnersMS []*bytes.Buffer
	losersMS  []*bytes.Buffer

	winnersI []interface{}
	losersI  []interface{}

	//msW []*marketsmith.MarketSmith
	//msL []*marketsmith.MarketSmith

	//msW []datautils.Contents
	//msL []datautils.Contents

	//Alpha string
	//Beta  string

	//PERatio           string
	//BookValue         string
	//InventoryTO       string
	//EarningsStability string

	//RSRating          string
	//GroupRSRating     string
	//AccDisRating      string
	//EPSRating         string
	//SMRRating         string
	//CompositeRating   string
	//TimelinessRating  string
	//SponsorshipRating string
	//UDVolRatio        string

	//ROE           string
	//EPSGrowthRate string
	//IndustryGroup string

	//Options string

	//Debt string
	//RnD  string

	//Mgmt  string
	//Banks string
	//Funds string

	//MarketCapitalization string
	//SharesInFloat        string
	//SharesOutstanding    string

	//AvgFundsHolding4Q string

	//Yield       string
	//ExDiv       string
	//NewCEO  string
	//FiveYearPERange string
	//CashFlow string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ChartMarketSmithNew(filterOrders []*transaction.Trade, winnersMS, losersMS []*bytes.Buffer) (*ChartMarketSmith, error) {

	c := new(ChartMarketSmith)

	c.filterOrders = filterOrders

	c.winnersMS = winnersMS
	c.losersMS = losersMS

	//c.msW = make([]*marketsmith.MarketSmith, len(c.winnersMS))
	//c.msL = make([]*marketsmith.MarketSmith, len(c.losersMS))

	//c.msW = make([]datautils.Contents, len(c.winnersMS))
	//c.msL = make([]datautils.Contents, len(c.losersMS))

	var err error

	//for i, w := range c.winnersMS {
	//m := marketsmith.MarketSmithNew()
	//err = json.Unmarshal(w.Bytes(), m)
	//if err != nil {
	//return nil, err
	//}

	////c.msW[i] = m
	//c.winnersI[i] = m
	//}

	//for i, l := range c.losersMS {
	//m := marketsmith.MarketSmithNew()
	//err = json.Unmarshal(l.Bytes(), m)
	//if err != nil {
	//return nil, err
	//}

	////c.msL[i] = m
	//c.losersI[i] = m
	//}

	err = c.getAlpha()
	if err != nil {
		return nil, err
	}

	err = c.getBeta()
	if err != nil {
		return nil, err
	}

	err = c.getPERatio()
	if err != nil {
		return nil, err
	}

	err = c.getBookValue()
	if err != nil {
		return nil, err
	}

	err = c.getInventoryTO()
	if err != nil {
		return nil, err
	}

	err = c.getEarningsStability()
	if err != nil {
		return nil, err
	}

	err = c.getRSRating()
	if err != nil {
		return nil, err
	}

	err = c.getGroupRSRating()
	if err != nil {
		return nil, err
	}

	err = c.getAccDisRating()
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

	err = c.getCompositeRating()
	if err != nil {
		return nil, err
	}

	err = c.getTimelinessRating()
	if err != nil {
		return nil, err
	}

	err = c.getSponsorshipRating()
	if err != nil {
		return nil, err
	}

	err = c.getUDVolRatio()
	if err != nil {
		return nil, err
	}

	err = c.getROE()
	if err != nil {
		return nil, err
	}

	err = c.getEPSGrowthRate()
	if err != nil {
		return nil, err
	}

	err = c.getIndustryGroup()
	if err != nil {
		return nil, err
	}

	err = c.getOptions()
	if err != nil {
		return nil, err
	}

	err = c.getDebt()
	if err != nil {
		return nil, err
	}

	err = c.getRnD()
	if err != nil {
		return nil, err
	}

	err = c.getMgmt()
	if err != nil {
		return nil, err
	}

	err = c.getBanks()
	if err != nil {
		return nil, err
	}

	err = c.getFunds()
	if err != nil {
		return nil, err
	}

	err = c.getMarketCapitalization()
	if err != nil {
		return nil, err
	}

	err = c.getSharesInFloat()
	if err != nil {
		return nil, err
	}

	err = c.getAvgFundsHolding4Q()
	if err != nil {
		return nil, err
	}

	return c, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getAlpha() error {

	//var err error
	//var interval float64 = 0.25

	//c.Alpha, err = columnChartFloatInterval("Alpha", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getBeta() error {

	//var err error
	//var interval float64 = 0.25

	//c.Beta, err = columnChartFloatInterval("Beta", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getEarningsStability() error {

	//var err error
	//var interval float64 = 5

	//c.EarningsStability, err = columnChartIntInterval("Earnings Stability", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getPERatio() error {

	//label := "P/E Ratio"

	//for _, w := range c.msW {
	//for _, f := range w.GetContents() {
	//if f.GetLabel() == label {
	//if f.GetValue() == config.NullValue {
	//break
	//}

	//ns := strings.Split(f.GetValue(), "(")
	//if len(ns) != 2 {
	//f.SetValue(config.NullValue)
	//break
	//}

	//newValue := strings.TrimSpace(ns[0])

	//f.SetValue(newValue)
	//break
	//}
	//}
	//}

	//for _, w := range c.msL {
	//for _, f := range w.GetContents() {
	//if f.GetLabel() == label {
	//if f.GetValue() == config.NullValue {
	//break
	//}

	//ns := strings.Split(f.GetValue(), "(")
	//if len(ns) != 2 {
	//f.SetValue(config.NullValue)
	//break
	//}

	//newValue := strings.TrimSpace(ns[0])

	//f.SetValue(newValue)
	//break
	//}
	//}
	//}

	//var err error
	//var interval float64 = 5

	//c.PERatio, err = columnChartIntInterval(label, interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getBookValue() error {

	//label := "Book Value"

	//for _, w := range c.msW {
	//for _, f := range w.GetContents() {
	//if f.GetLabel() == label {
	//if f.GetValue() == config.NullValue {
	//break
	//}

	//newValue := strings.Replace(strings.ToLower(f.GetValue()), "x", "", -1)
	//f.SetValue(newValue)
	//break
	//}
	//}
	//}

	//for _, w := range c.msL {
	//for _, f := range w.GetContents() {
	//if f.GetLabel() == label {
	//if f.GetValue() == config.NullValue {
	//break
	//}

	//newValue := strings.Replace(strings.ToLower(f.GetValue()), "x", "", -1)
	//f.SetValue(newValue)
	//break
	//}
	//}
	//}

	//var err error
	//var interval float64 = 0.25

	//c.BookValue, err = columnChartFloatInterval(label, interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getInventoryTO() error {

	//label := "Inventory T/O"

	//for _, w := range c.msW {
	//for _, f := range w.GetContents() {
	//if f.GetLabel() == label {
	//if f.GetValue() == config.NullValue {
	//break
	//}

	//newValue := strings.Replace(strings.ToLower(f.GetValue()), "x", "", -1)
	//f.SetValue(newValue)
	//break
	//}
	//}
	//}

	//for _, w := range c.msL {
	//for _, f := range w.GetContents() {
	//if f.GetLabel() == label {
	//if f.GetValue() == config.NullValue {
	//break
	//}

	//newValue := strings.Replace(strings.ToLower(f.GetValue()), "x", "", -1)
	//f.SetValue(newValue)
	//break
	//}
	//}
	//}

	//var err error
	//var interval float64 = 0.25

	//c.InventoryTO, err = columnChartFloatInterval(label, interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getRSRating() error {

	//var err error
	//var interval float64 = 10.0

	//c.RSRating, err = columnChartIntInterval("RS Rating", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getGroupRSRating() error {

	//var err error
	//var interval float64 = 10.0

	//c.GroupRSRating, err = columnChartIntInterval("Group RS Rating", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getAccDisRating() error {

	//var err error

	//c.AccDisRating, err = columnChartStringRank("Acc/Dis Rating", c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getEPSRating() error {

	//var err error
	//var interval float64 = 10.0

	//c.EPSRating, err = columnChartIntInterval("EPS Rating", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getSMRRating() error {

	//var err error

	//c.SMRRating, err = columnChartStringRank("SMR Rating", c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getCompositeRating() error {

	//var err error
	//var interval float64 = 10.0

	//c.CompositeRating, err = columnChartIntInterval("Composite Rating", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getTimelinessRating() error {

	//var err error

	//c.TimelinessRating, err = columnChartStringRank("Timeliness Rating", c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getSponsorshipRating() error {

	//var err error

	//c.SponsorshipRating, err = columnChartStringRank("Sponsorship Rating", c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getUDVolRatio() error {

	//var err error
	//var interval float64 = 0.5

	//c.UDVolRatio, err = columnChartFloatInterval("U/D Vol Ratio", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getROE() error {

	//label := "Return on Equity"

	//for _, w := range c.msW {
	//for _, f := range w.GetContents() {
	//if f.GetLabel() == label {
	//if strings.Contains(f.GetValue(), config.NullValue) {
	//newValue := strings.Replace(f.GetValue(), "[", "", -1)
	//newValue = strings.Replace(newValue, "]", "", -1)
	//f.SetValue(newValue)
	//break
	//}
	//}
	//}
	//}

	//for _, w := range c.msL {
	//for _, f := range w.GetContents() {
	//if f.GetLabel() == label {
	//if strings.Contains(f.GetValue(), config.NullValue) {
	//newValue := strings.Replace(f.GetValue(), "x", "", -1)
	//newValue = strings.Replace(newValue, "]", "", -1)
	//f.SetValue(newValue)
	//break
	//}
	//}
	//}
	//}

	//var err error
	//var interval float64 = 5.0

	//c.ROE, err = columnChartPercent("Return on Equity", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getEPSGrowthRate() error {

	//var err error
	//var interval float64 = 20.0

	//c.EPSGrowthRate, err = columnChartPercent("EPS Growth Rate", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getIndustryGroup() error {

	//var err error

	//c.IndustryGroup, err = columnChartString("Industry Group", c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getOptions() error {

	//var err error

	//c.Options, err = columnChartString("Options", c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getDebt() error {

	//var err error
	//var interval float64 = 5.0

	//c.Debt, err = columnChartPercent("Debt", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getRnD() error {

	//var err error
	//var interval float64 = 5.0

	//c.RnD, err = columnChartPercent("R&amp;D", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getMgmt() error {

	//var err error
	//var interval float64 = 5.0

	//c.Mgmt, err = columnChartPercent("Mgmt", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getBanks() error {

	//var err error
	//var interval float64 = 5.0

	//c.Banks, err = columnChartPercent("Banks", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getFunds() error {

	//var err error
	//var interval float64 = 5.0

	//c.Funds, err = columnChartPercent("Funds", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getMarketCapitalization() error {

	//label := "Market Capitalization"

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

	//var smallCapThreshold float64 = 1000000000.0
	//var largeCapThreshold float64 = 10000000000.0

	//smallCapW := 0
	//midCapW := 0
	//largeCapW := 0

	//smallCapL := 0
	//midCapL := 0
	//largeCapL := 0

	//for _, o := range c.msW {
	//for _, f := range o.GetContents() {
	//if f.GetLabel() == label {
	//if f.GetValue() == config.NullValue {
	//return fmt.Errorf("Market Smith Market Capitalization Chart Parsing Error\n")
	//}

	//ns := strings.Split(strings.TrimSpace(f.GetValue()), " ")
	//if len(ns) != 2 {
	//return fmt.Errorf("Market Smith Market Capitalization Chart Parsing Error\n")
	//}

	//amountS := strings.Replace(ns[0], "$", "", -1)
	//unit := ns[1]

	//amount, err := strconv.ParseFloat(amountS, 64)
	//if err != nil {
	//return fmt.Errorf("Market Smith Market Capitalization Chart Parsing Error\n")
	//}

	//var newValue float64

	//switch unit {
	//case "Mil":
	//newValue = amount * 1000000.0
	//case "Million":
	//newValue = amount * 1000000.0
	//case "Bil":
	//newValue = amount * 1000000000.0
	//case "Billion":
	//newValue = amount * 1000000000.0
	//default:
	//return fmt.Errorf("Market Smith Market Capitalization Chart Parsing Error\n")
	//}

	//if newValue <= smallCapThreshold {
	//smallCapW += 1
	//} else if newValue <= largeCapThreshold {
	//midCapW += 1
	//} else if newValue > largeCapThreshold {
	//largeCapW += 1
	//}

	//break
	//}
	//}
	//}

	//for _, o := range c.msL {
	//for _, f := range o.GetContents() {
	//if f.GetLabel() == label {
	//if f.GetValue() == config.NullValue {
	//return fmt.Errorf("Market Smith Market Capitalization Chart Parsing Error\n")
	//}

	//ns := strings.Split(strings.TrimSpace(f.GetValue()), " ")
	//if len(ns) != 2 {
	//return fmt.Errorf("Market Smith Market Capitalization Chart Parsing Error\n")
	//}

	//amountS := strings.Replace(ns[0], "$", "", -1)
	//unit := ns[1]

	//amount, err := strconv.ParseFloat(amountS, 64)
	//if err != nil {
	//return fmt.Errorf("Market Smith Market Capitalization Chart Parsing Error\n")
	//}

	//var newValue float64

	//switch unit {
	//case "Mil":
	//newValue = amount * 1000000.0
	//case "Million":
	//newValue = amount * 1000000.0
	//case "Bil":
	//newValue = amount * 1000000000.0
	//case "Billion":
	//newValue = amount * 1000000000.0
	//default:
	//return fmt.Errorf("Market Smith Market Capitalization Chart Parsing Error\n")
	//}

	//if newValue <= smallCapThreshold {
	//smallCapL += 1
	//} else if newValue <= largeCapThreshold {
	//midCapL += 1
	//} else if newValue > largeCapThreshold {
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

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getSharesInFloat() error {

	//label := "Shares in Float"

	//for _, w := range c.msW {
	//for _, f := range w.GetContents() {
	//if f.GetLabel() == label {
	//if f.GetValue() == config.NullValue {
	//break
	//}

	//ns := strings.Split(strings.TrimSpace(f.GetValue()), " ")
	//if len(ns) != 2 {
	//f.SetValue(config.NullValue)
	//break
	//}

	//amount, err := strconv.ParseFloat(ns[0], 64)
	//if err != nil {
	//f.SetValue(config.NullValue)
	//break
	//}

	//unit := ns[1]

	//var newValue string

	//switch unit {
	//case "Mil":
	//newValue = fmt.Sprintf("%v", int64(amount*1000000.0))
	//case "Million":
	//newValue = fmt.Sprintf("%v", int64(amount*1000000.0))
	//case "Bil":
	//newValue = fmt.Sprintf("%v", int64(amount*1000000000.0))
	//case "Billion":
	//newValue = fmt.Sprintf("%v", int64(amount*1000000000.0))
	//default:
	//f.SetValue(config.NullValue)
	//break
	//}

	//f.SetValue(newValue)
	//break
	//}
	//}
	//}

	//for _, w := range c.msL {
	//for _, f := range w.GetContents() {
	//if f.GetLabel() == label {
	//if f.GetValue() == config.NullValue {
	//break
	//}

	//ns := strings.Split(strings.TrimSpace(f.GetValue()), " ")
	//if len(ns) != 2 {
	//f.SetValue(config.NullValue)
	//break
	//}

	//amount, err := strconv.ParseFloat(ns[0], 64)
	//if err != nil {
	//f.SetValue(config.NullValue)
	//break
	//}

	//unit := ns[1]

	//var newValue string

	//switch unit {
	//case "Mil":
	//newValue = fmt.Sprintf("%v", int64(amount*1000000.0))
	//case "Million":
	//newValue = fmt.Sprintf("%v", int64(amount*1000000.0))
	//case "Bil":
	//newValue = fmt.Sprintf("%v", int64(amount*1000000000.0))
	//case "Billion":
	//newValue = fmt.Sprintf("%v", int64(amount*1000000000.0))
	//default:
	//f.SetValue(config.NullValue)
	//break
	//}

	//f.SetValue(newValue)
	//break
	//}
	//}
	//}

	//var err error
	//var interval float64 = 5000000.0

	//c.SharesInFloat, err = columnChartIntInterval("Shares in Float", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartMarketSmith) getAvgFundsHolding4Q() error {

	//var err error
	//var interval float64 = 50.0

	//c.AvgFundsHolding4Q, err = columnChartIntInterval("Avg Funds Holding 4Q", interval, c.msW, c.msL)
	//if err != nil {
	//return err
	//}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
