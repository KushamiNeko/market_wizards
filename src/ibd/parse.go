package ibd

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	regexSymbol      = `<div class="companyTitle">\s*<span class="companyName">\s*[\s\S]+?\s*<a\s*[^>]+>\((\w+)\)<\/a>\s*<\/div>\s*<\/div>`
	regexRankInGroup = `<span[^>]+>\s*(\d+)\s*<\/span>\s*<div class=\"listCompany_right\">\s*<div class=\"listCompanyName\">\s*<a[^>]+>(\w+)<\/a>\s*([^<]+)\s*<\/div>`

	regexQuote = `<tr[^>]*>\s*<td class=\"first\">\s*<a[^>]+>([^<]+)<\/a>\s*<\/td>\s*<td class=\"second\">\s*(<span\s*\w+=\"(\w+)\"\s*\/*>\s*([^<]*)\s*(?:<\/span>)*(?:\s*<input[^>]+>)*|[^<]+)?\s*<\/td>\s*<td class=\"third\">\s*(?:<a[^>]+>\s*[^<]+<img src=\".+\/(\w+)\.gif\"[^>]+>[^<]+<\/a>|[^<]+)*\s*<\/td>\s*<\/tr>`

	regexPercent = `([0-9.-]+)%`
	regexPrice   = `<span[^>]+>\s*\$([0-9.]+)\s*<\/span>`

	//regexMktCap = `\$\s*([0-9.]+)\s*(\w+)`
	regexMktCap = `^\$\s*([0-9.]+)\s*(\w+)$`
	//regexVolume = `([0-9.,]+)\s*(\w*)`
	regexVolume = `^([0-9.,]+)\s*(\w*)$`

	//regexFloat = `^[0-9.]+$`
	regexFloat = `^\d+\.\d+$`
	regexInt   = `^[0-9]+$`

	regexArrow = `^<span\s*class="([a-zA-Z]+)"\s*\/>$`

	regexLabel = `\s*<a\s*class=\"glossDef\"[^>]+>\s*([^<]+)\s*<\/a>\s*`

	none = "N/A"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	reSymbol *regexp.Regexp

	reRankInGroup *regexp.Regexp
	reQuote       *regexp.Regexp

	rePercent *regexp.Regexp
	rePrice   *regexp.Regexp

	reMktCap *regexp.Regexp
	reVolume *regexp.Regexp

	reFloat *regexp.Regexp
	reInt   *regexp.Regexp

	reArrow *regexp.Regexp

	reLabel *regexp.Regexp
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	reSymbol = regexp.MustCompile(regexSymbol)

	reRankInGroup = regexp.MustCompile(regexRankInGroup)
	reQuote = regexp.MustCompile(regexQuote)

	rePercent = regexp.MustCompile(regexPercent)
	rePrice = regexp.MustCompile(regexPrice)

	reMktCap = regexp.MustCompile(regexMktCap)
	reVolume = regexp.MustCompile(regexVolume)

	reFloat = regexp.MustCompile(regexFloat)
	reInt = regexp.MustCompile(regexInt)

	reArrow = regexp.MustCompile(regexArrow)

	reLabel = regexp.MustCompile(regexLabel)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//type IBDCheckup struct {
//ID string

//Symbol string

//RankInGroup           int
//CompositeRating       int
//MarketUptrend         string
//IndustryGroupRank     int
//GroupRSRating         string
//EPSRating             int
//EPSChgLastQtr         float64
//Last3QtrsAvgEPSGrowth float64
//NQtrsOfEPSAccel       int

//EPSEstChgCurrentQtr    float64
//EstimateRevisions      string
//LastQtrEarningsSuprise float64

//ThrYrEpsGrowthRate    float64
//NYrsOfAnnualEPSGrowth int
//EPSEstChgCurrentYr    float64

//SMRRating            string
//SalesChgLastQtr      float64
//ThrYrSalesGrowthRate float64
//AnnualPreTaxMargin   float64
//AnnualROE            float64
//DebtEquityRatio      float64

////Price          float64

//RSRating       int
//Off52WeekHigh  float64
//PriceVS50DayMA float64
//AvgVolume50Day int64

//MarketCapital int64
//AccDisRating  string
//UpDownVolume  float64
//ChgInFunds    float64
//QtrsOfIncFund int
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type IBDCheckupDatastore struct {
	ID   string
	Data []byte `datastore:",noindex"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func IBDCheckupDatastoreNew(date int, symbol string, data []byte) *IBDCheckupDatastore {
	ibd := new(IBDCheckupDatastore)

	//ibd.ID = fmt.Sprintf("%d_%v", date, symbol)
	ibd.ID = IBDCheckupDatastoreGetID(date, symbol)
	ibd.Data = data

	return ibd
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func IBDCheckupDatastoreGetID(date int, symbol string) string {
	return fmt.Sprintf("%d_%v", date, symbol)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type IBDCheckup struct {
	Contents []field
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func IBDCheckupNew() *IBDCheckup {
	checkup := new(IBDCheckup)
	checkup.Contents = make([]field, 0)

	return checkup
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type field struct {
	Label     string
	Value     string
	Condition string `json:",omitempty"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Parse(buffer *bytes.Buffer) (*IBDCheckup, error) {

	checkup := IBDCheckupNew()

	var results [][]string

	results = reSymbol.FindAllStringSubmatch(buffer.String(), -1)

	if results == nil {
		return nil, fmt.Errorf("no matching found in the content\n")
	}

	if len(results) <= 0 {
		return nil, fmt.Errorf("no matching found in the content\n")
	}

	if len(results[0]) < 2 {
		return nil, fmt.Errorf("parsing symbol error\n")
	}

	symbol := strings.TrimSpace(results[0][1])

	checkup.Contents = append(checkup.Contents, field{
		"Symbol",
		symbol,
		"",
	})

	results = reRankInGroup.FindAllStringSubmatch(buffer.String(), -1)

	for _, v := range results {

		if len(v) < 3 {
			return nil, fmt.Errorf("parsing rank in group error\n")
		}

		r := v[1]
		s := v[2]

		if s == symbol {

			checkup.Contents = append(checkup.Contents, field{
				"Rank in Group",
				strings.TrimSpace(r),
				"",
			})

			break
		}
	}

	results = reQuote.FindAllStringSubmatch(buffer.String(), -1)

	if results == nil {
		return nil, fmt.Errorf("parsing contents error\n")
	}

	for _, r := range results {

		if len(r) < 6 {
			return nil, fmt.Errorf("parsing contents error\n")
		}

		match := reLabel.FindStringSubmatch(r[0])
		if len(match) < 2 {
			return nil, fmt.Errorf("parsing label error\n")
		}

		v := strings.TrimSpace(r[2])
		var value string

		var valueI interface{}
		var err error

		valueI, err = parsePrice(v)
		if err == nil {
			value = fmt.Sprintf("$%v", valueI)
			goto found
		}

		valueI, err = parseMktCap(v)
		if err == nil {
			value = fmt.Sprintf("$%v", valueI)
			goto found
		}

		valueI, err = parsePercent(v)
		if err == nil {
			value = fmt.Sprintf("%v%%", valueI)
			goto found
		}

		valueI, err = parseVolume(v)
		if err == nil {
			value = fmt.Sprintf("%v", valueI)
			goto found
		}

		valueI, err = parseArrow(v)
		if err == nil {
			value = fmt.Sprintf("%v", valueI)
			goto found
		}

		value = v

	found:

		condition := strings.TrimSpace(r[5])

		if value == none {
			condition = none
		}

		checkup.Contents = append(checkup.Contents, field{
			strings.TrimSpace(match[1]),
			//strings.TrimSpace(r[2]),
			value,
			//strings.TrimSpace(r[5]),
			condition,
		})
	}

	//for _, c := range checkup.Contents {
	//fmt.Printf("%s: %s %s\n", c.Label, c.Value, c.Condition)
	//}

	return checkup, nil

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func Parse(buffer *bytes.Buffer) (*IBDCheckup, error) {

//info := new(IBDCheckup)

//checkup := IBDCheckupNew()

//var results [][]string

//results = reSymbol.FindAllStringSubmatch(buffer.String(), -1)

//info.Symbol = cleanup(results, 0, 1)

//checkup.Contents = append(checkup.Contents, checkupField{
//"Symbol",
//strings.TrimSpace(results[0][1]),
//"",
//})

//results = reRankInGroup.FindAllStringSubmatch(buffer.String(), -1)

//for _, v := range results {
//r := v[1]
//s := v[2]

//if s == info.Symbol {
////var err error

////checkup["Rank in Group"], _ = parseInt(r)

//checkup.Contents = append(checkup.Contents, checkupField{
//"Rank in Group",
//strings.TrimSpace(r),
//"",
//})

//info.RankInGroup, _ = parseInt(r)
////if err != nil {
////return nil, err
////}

//break
//}
//}

//results = reQuote.FindAllStringSubmatch(buffer.String(), -1)

//for _, r := range results {
//match := reLabel.FindStringSubmatch(r[0])
//if len(match) == 0 {
////continue
//return nil, fmt.Errorf("parsing label error\n")
//}

//checkup.Contents = append(checkup.Contents, checkupField{
//strings.TrimSpace(match[1]),
//strings.TrimSpace(r[2]),
//strings.TrimSpace(r[5]),
//})

////if strings.Compare(strings.TrimSpace(match[1]), strings.TrimSpace(label)) == 0 {
////return strings.TrimSpace(r[col])
////}
//}

//fmt.Println(checkup)

//return nil, nil

//var err error

//info.CompositeRating, err = parseInt(cleanupL(results, "Composite Rating", 2))
//if err != nil {
//return nil, err
//}

//info.MarketUptrend = cleanupL(results, "Market Direction", 5)
//if info.MarketUptrend == "" {
//info.MarketUptrend = cleanupL(results, "Market in confirmed uptrend", 5)
//}

//info.IndustryGroupRank, err = parseInt(cleanupL(results, "Industry Group Rank (1 to 197)", 2))
//if err != nil {
//return nil, err
//}

//info.GroupRSRating = cleanupL(results, "Group RS Rating", 2)

//info.EPSRating, err = parseInt(cleanupL(results, "EPS Rating", 2))
//if err != nil {
//return nil, err
//}

//info.EPSChgLastQtr, err = parsePercent(cleanupL(results, "EPS % Chg (Last Qtr)", 2))
//if err != nil {
//return nil, err
//}

//info.Last3QtrsAvgEPSGrowth, err = parsePercent(cleanupL(results, "Last 3 Qtrs Avg EPS Growth", 2))
//if err != nil {
//return nil, err
//}

//info.NQtrsOfEPSAccel, err = parseInt(cleanupL(results, "# Qtrs of EPS Acceleration", 2))
//if err != nil {
//return nil, err
//}

//info.EPSEstChgCurrentQtr, err = parsePercent(cleanupL(results, "EPS Est % Chg (Current Qtr)", 2))
//if err != nil {
//return nil, err
//}

//info.EstimateRevisions = cleanupL(results, "Estimate Revisions", 5)

//info.LastQtrEarningsSuprise, err = parsePercent(cleanupL(results, `Last Quarter % Earnings Surprise`, 2))
//if err != nil {
//return nil, err
//}

//info.ThrYrEpsGrowthRate, err = parsePercent(cleanupL(results, "3 Yr EPS Growth Rate", 2))
//if err != nil {
//return nil, err
//}

//info.NYrsOfAnnualEPSGrowth, err = parseInt(cleanupL(results, "Consecutive Yrs of Annual EPS Growth", 2))
//if err != nil {
//return nil, err
//}

//info.EPSEstChgCurrentYr, err = parsePercent(cleanupL(results, "EPS Est % Chg for Current Year", 2))
//if err != nil {
//return nil, err
//}

//info.SMRRating = cleanupL(results, "SMR Rating", 2)

//info.SalesChgLastQtr, err = parsePercent(cleanupL(results, "Sales % Chg (Last Qtr)", 2))
//if err != nil {
//return nil, err
//}

//info.ThrYrSalesGrowthRate, err = parsePercent(cleanupL(results, "3 Yr Sales Growth Rate", 2))
//if err != nil {
//return nil, err
//}

//info.AnnualPreTaxMargin, err = parsePercent(cleanupL(results, "Annual Pre-Tax Margin", 2))
//if err != nil {
//return nil, err
//}

//info.AnnualROE, err = parsePercent(cleanupL(results, "Annual ROE", 2))
//if err != nil {
//return nil, err
//}

//info.DebtEquityRatio, err = parsePercent(cleanupL(results, "Debt/Equity Ratio", 2))
//if err != nil {
//return nil, err
//}

////info.Price, err = parsePrice(cleanup(results, 20, 2))
////if err != nil {
////return nil, err
////}

//info.RSRating, err = parseInt(cleanupL(results, "RS Rating", 2))
//if err != nil {
//return nil, err
//}

//info.Off52WeekHigh, err = parsePercent(cleanupL(results, "% Off 52 Week High", 2))
//if err != nil {
//return nil, err
//}

//info.PriceVS50DayMA, err = parsePercent(cleanupL(results, "Price vs. 50-Day Moving Average", 2))
//if err != nil {
//return nil, err
//}

//info.AvgVolume50Day, err = parseVolume(cleanupL(results, "50-Day Average Volume", 2))
//if err != nil {
//return nil, err
//}

//info.MarketCapital, err = parseMktCap(cleanupL(results, "Market Capitalization", 2))
//if err != nil {
//return nil, err
//}

//info.AccDisRating = cleanupL(results, "Accumulation/Distribution Rating", 2)

//info.UpDownVolume, err = parseFloat(cleanupL(results, "Up/Down Volume", 2))
//if err != nil {
//return nil, err
//}

//info.ChgInFunds, err = parsePercent(cleanupL(results, "% Change In Funds Owning Stock", 2))
//if err != nil {
//return nil, err
//}

//info.QtrsOfIncFund, err = parseInt(cleanupL(results, "Qtrs Of Increasing Fund Ownership", 2))
//if err != nil {
//return nil, err
//}

//return info, nil

//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
