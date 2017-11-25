package ibd

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"regexp"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	regexSymbol      = `<div class="companyTitle">\s*<span class="companyName">\s*[\s\S]+?\s*<a\s*[^>]+>\((\w+)\)<\/a>\s*<\/div>\s*<\/div>`
	regexRankInGroup = `<span[^>]+>\s*(\d+)\s*<\/span>\s*<div class=\"listCompany_right\">\s*<div class=\"listCompanyName\">\s*<a[^>]+>(\w+)<\/a>\s*([^<]+)\s*<\/div>`

	regexQuote = `<tr[^>]*>\s*<td class=\"first\">\s*<a[^>]+>([^<]+)<\/a>\s*<\/td>\s*<td class=\"second\">\s*(<span\s*\w+=\"(\w+)\"\s*\/*>\s*([^<]*)\s*(?:<\/span>)*(?:\s*<input[^>]+>)*|[^<]+)?\s*<\/td>\s*<td class=\"third\">\s*(?:<a[^>]+>\s*[^<]+<img src=\".+\/(\w+)\.gif\"[^>]+>[^<]+<\/a>|[^<]+)*\s*<\/td>\s*<\/tr>`

	regexPercent = `([0-9.-]+)%`
	regexPrice   = `<span[^>]+>\s*\$([0-9.]+)\s*<\/span>`

	regexMktCap = `\$\s*([0-9.]+)\s*(\w+)`
	regexVolume = `([0-9.,]+)\s*(\w*)`
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
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type IBDCheckup struct {
	Symbol string

	RankInGroup           int
	CompositeRating       int
	MarketUptrend         string
	IndustryGroupRank     int
	EPSRating             int
	EPSChgLastQtr         float32
	Last3QtrsAvgEPSGrowth float32
	NQtrsOfEPSAccel       int

	EPSEstChgCurrentQtr    float32
	EstimateRevisions      string
	LastQtrEarningsSuprise float32

	ThrYrEpsGrowthRate    float32
	NYrsOfAnnualEPSGrowth int
	EPSEstChgCurrentYr    float32

	SMRRating            string
	SalesChgLastQtr      float32
	ThrYrSalesGrowthRate float32
	AnnualPreTaxMargin   float32
	AnnualROE            float32
	DebtEquityRatio      float32

	Price          float32
	RSRating       int
	Off52WeekHigh  float32
	PriceVS50DayMA float32
	AvgVolume50Day int64

	MarketCapital int64
	AccDisRating  string
	UpDownVolume  float32
	ChgInFunds    float32
	QtrsOfIncFund int
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Parse(buffer *bytes.Buffer) (*IBDCheckup, error) {

	info := new(IBDCheckup)

	var results [][]string

	results = reSymbol.FindAllStringSubmatch(buffer.String(), -1)

	info.Symbol = cleanup(results, 0, 1)

	results = reRankInGroup.FindAllStringSubmatch(buffer.String(), -1)

	for _, v := range results {
		r := v[1]
		s := v[2]

		if s == info.Symbol {
			var err error

			info.RankInGroup, _ = parseInt(r)
			if err != nil {
				return nil, err
			}

			break
		}
	}

	results = reQuote.FindAllStringSubmatch(buffer.String(), -1)

	var err error

	info.CompositeRating, err = parseInt(cleanup(results, 0, 2))
	if err != nil {
		return nil, err
	}

	info.MarketUptrend = cleanup(results, 1, 5)

	info.IndustryGroupRank, err = parseInt(cleanup(results, 2, 2))
	if err != nil {
		return nil, err
	}

	info.EPSRating, err = parseInt(cleanup(results, 4, 2))
	if err != nil {
		return nil, err
	}

	info.EPSChgLastQtr, err = parsePercent(cleanup(results, 5, 2))
	if err != nil {
		return nil, err
	}

	info.Last3QtrsAvgEPSGrowth, err = parsePercent(cleanup(results, 6, 2))
	if err != nil {
		return nil, err
	}

	info.NQtrsOfEPSAccel, err = parseInt(cleanup(results, 7, 2))
	if err != nil {
		return nil, err
	}

	info.EPSEstChgCurrentQtr, err = parsePercent(cleanup(results, 8, 2))
	if err != nil {
		return nil, err
	}

	info.EstimateRevisions = cleanup(results, 9, 3)

	info.LastQtrEarningsSuprise, err = parsePercent(cleanup(results, 10, 2))
	if err != nil {
		return nil, err
	}

	info.ThrYrEpsGrowthRate, err = parsePercent(cleanup(results, 11, 2))
	if err != nil {
		return nil, err
	}

	info.NYrsOfAnnualEPSGrowth, err = parseInt(cleanup(results, 12, 2))
	if err != nil {
		return nil, err
	}

	info.EPSEstChgCurrentYr, err = parsePercent(cleanup(results, 13, 2))
	if err != nil {
		return nil, err
	}

	info.SMRRating = cleanup(results, 14, 2)

	info.SalesChgLastQtr, err = parsePercent(cleanup(results, 15, 2))
	if err != nil {
		return nil, err
	}

	info.ThrYrSalesGrowthRate, err = parsePercent(cleanup(results, 16, 2))
	if err != nil {
		return nil, err
	}

	info.AnnualPreTaxMargin, err = parsePercent(cleanup(results, 17, 2))
	if err != nil {
		return nil, err
	}

	info.AnnualROE, err = parsePercent(cleanup(results, 18, 2))
	if err != nil {
		return nil, err
	}

	info.DebtEquityRatio, err = parsePercent(cleanup(results, 19, 2))
	if err != nil {
		return nil, err
	}

	info.Price, err = parsePrice(cleanup(results, 20, 2))
	if err != nil {
		return nil, err
	}

	info.RSRating, err = parseInt(cleanup(results, 21, 2))
	if err != nil {
		return nil, err
	}

	info.Off52WeekHigh, err = parsePercent(cleanup(results, 22, 2))
	if err != nil {
		return nil, err
	}

	info.PriceVS50DayMA, err = parsePercent(cleanup(results, 23, 2))
	if err != nil {
		return nil, err
	}

	info.AvgVolume50Day, err = parseVolume(cleanup(results, 24, 2))
	if err != nil {
		return nil, err
	}

	info.MarketCapital, err = parseMktCap(cleanup(results, 25, 2))
	if err != nil {
		return nil, err
	}

	info.AccDisRating = cleanup(results, 26, 2)

	info.UpDownVolume, err = parseFloat(cleanup(results, 27, 2))
	if err != nil {
		return nil, err
	}

	info.ChgInFunds, err = parsePercent(cleanup(results, 28, 2))
	if err != nil {
		return nil, err
	}

	info.QtrsOfIncFund, err = parseInt(cleanup(results, 29, 2))
	if err != nil {
		return nil, err
	}

	return info, nil

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
