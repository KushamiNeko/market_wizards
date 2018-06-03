package marketsmith

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	regexInfoCell = `\s*<div\s*class="cell">\s*<div\s*class="key">([^<]+)<\/div>\s*<div\s*class="value">([^<]+)<\/div>\s*<\/div>\s*`

	regexIndustryGroup = `\s*<div\s*class="companySymbol">\s*<span\s*class="companyInfCoName">[^<]+<\/span>\(\w+\)\s*\S+\s*([^<]+)<\/div>\s*`

	regexOptions = `\s*<div\s*class="Options">Options\s*([^<]+)<\/div>\s*`

	regexFloatShare = `\s*<div\s*class="companyInfoCenter"\s*[^>]+>\s*<div\s*class="companyInfoLable">\s*<div\s*class="cell">([^<]+)<\/div>\s*<div\s*class="cell">([^<]+)<\/div>\s*<div\s*class="cell">([^<]+)<\/div>\s*<\/div>\s*<div\s*class="companyInfoVlaue">\s*<div\s*class="cell">([^<]+)<\/div>\s*<div\s*class="cell">([^<]+)<\/div>\s*<div\s*class="cell">([^<]+)<\/div>\s*<\/div>\s*<\/div>\s*`

	regexQuarterlyResults = `\s*<div\s*class="cell"\s*style="display:\s*block;">\s*<div\s*class="quarterly\s*\w*">[^<]+<span\s*[^>]+>\s*[^<]*\s*<\/span>\s*<\/div>\s*<div\s*class="eps\s*\w*">\s*([^<]+)\s*<\/div>\s*<div\s*class="epsChg\s*\w*">([^<]+)<\/div>\s*<div\s*class="sales\s*\w*">([^<]+)<\/div>\s*<div\s*class="salesChg\s*\w*">([^<]+)<\/div>\s*<\/div>\s*`

	regexFunds = `\s*<div\s*class="cell">\s*<span>[^<]+<\/span>\s*<span>(\d+)<\/span>\s*<\/div>\s*`
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	reInfoCell *regexp.Regexp

	reIndustryGroup *regexp.Regexp
	reOptions       *regexp.Regexp

	reFloatShare       *regexp.Regexp
	reQuarterlyResults *regexp.Regexp

	reFunds *regexp.Regexp
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	reInfoCell = regexp.MustCompile(regexInfoCell)

	reIndustryGroup = regexp.MustCompile(regexIndustryGroup)
	reOptions = regexp.MustCompile(regexOptions)

	reFloatShare = regexp.MustCompile(regexFloatShare)
	reQuarterlyResults = regexp.MustCompile(regexQuarterlyResults)

	reFunds = regexp.MustCompile(regexFunds)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type MarketSmithDatastore struct {
	ID   string
	Data []byte `datastore:",noindex"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func MarketSmithDatastoreNew(date int, symbol, chartType string, data []byte) *MarketSmithDatastore {
	m := new(MarketSmithDatastore)

	m.ID = MarketSmithDatastoreGetID(date, symbol, chartType)
	m.Data = data

	return m
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func MarketSmithDatastoreGetID(date int, symbol, chartType string) string {
	return fmt.Sprintf("%d_%v_%v", date, symbol, chartType)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type MarketSmith struct {
	Contents []field
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type field struct {
	Label string
	Value string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func MarketSmithNew() *MarketSmith {
	m := new(MarketSmith)
	m.Contents = make([]field, 0)

	return m
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Parse(buffer *bytes.Buffer) (*MarketSmith, error) {

	m := MarketSmithNew()

	var err error

	err = m.getInfoCell(buffer)
	if err != nil {
		return nil, err
	}

	err = m.getIndustryGroup(buffer)
	if err != nil {
		return nil, err
	}

	err = m.getOptions(buffer)
	if err != nil {
		return nil, err
	}

	err = m.getFloatShare(buffer)
	if err != nil {
		return nil, err
	}

	err = m.getQuarterlyResults(buffer)
	if err != nil {
		return nil, err
	}

	err = m.getFunds(buffer)
	if err != nil {
		return nil, err
	}

	return m, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketSmith) getInfoCell(buffer *bytes.Buffer) error {

	var results [][]string

	results = reInfoCell.FindAllStringSubmatch(buffer.String(), -1)
	if results == nil {
		return fmt.Errorf("no matching found in the Info Cell\n")
	}

	for _, r := range results {
		if len(r) < 3 {
			return fmt.Errorf("problems occur while parsing Info Cell \n")
		}

		m.Contents = append(m.Contents, field{
			strings.TrimSpace(r[1]),
			strings.TrimSpace(r[2]),
		})
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketSmith) getIndustryGroup(buffer *bytes.Buffer) error {

	var results [][]string

	results = reIndustryGroup.FindAllStringSubmatch(buffer.String(), -1)
	if results == nil {
		return fmt.Errorf("no matching found in the Industry Group\n")
	}

	for _, r := range results {
		if len(r) < 2 {
			return fmt.Errorf("problems occur while parsing Industry Group \n")
		}

		m.Contents = append(m.Contents, field{
			"Industry Group",
			strings.TrimSpace(r[1]),
		})
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketSmith) getOptions(buffer *bytes.Buffer) error {

	var results [][]string

	results = reOptions.FindAllStringSubmatch(buffer.String(), -1)
	if results == nil {
		return fmt.Errorf("no matching found in the Options\n")
	}

	for _, r := range results {
		if len(r) < 2 {
			return fmt.Errorf("problems occur while parsing Options \n")
		}

		m.Contents = append(m.Contents, field{
			"Options",
			strings.TrimSpace(r[1]),
		})
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketSmith) getFloatShare(buffer *bytes.Buffer) error {

	var results [][]string

	results = reFloatShare.FindAllStringSubmatch(buffer.String(), -1)
	if results == nil {
		return fmt.Errorf("no matching found in the Float Share\n")
	}

	for _, r := range results {
		if len(r) < 7 {
			return fmt.Errorf("problems occur while parsing Float Share \n")
		}

		m.Contents = append(m.Contents, field{
			strings.TrimSpace(r[1]),
			strings.TrimSpace(r[4]),
		})

		m.Contents = append(m.Contents, field{
			strings.TrimSpace(r[2]),
			strings.TrimSpace(r[5]),
		})

		m.Contents = append(m.Contents, field{
			strings.TrimSpace(r[3]),
			strings.TrimSpace(r[6]),
		})
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketSmith) getQuarterlyResults(buffer *bytes.Buffer) error {

	//var results [][]string

	//results = reQuarterlyResults.FindAllStringSubmatch(buffer.String(), -1)
	//if results == nil {
	//return fmt.Errorf("no matching found in the Quarterly Results\n")
	//}

	//for _, r := range results {
	//fmt.Println(len(r))
	//if len(r) < 5 {
	//return fmt.Errorf("problems occur while parsing Quarterly Results \n")
	//}

	//fmt.Println(r)

	////m.Contents = append(m.Contents, field{
	////"Industry Group",
	////r[1],
	////})
	//}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketSmith) getFunds(buffer *bytes.Buffer) error {

	var results [][]string

	results = reFunds.FindAllStringSubmatch(buffer.String(), -1)
	if results == nil {
		return fmt.Errorf("no matching found in the Funds\n")
	}

	funds := make([]float64, 4)

	for i, r := range results {
		if len(r) < 2 {
			return fmt.Errorf("problems occur while parsing Funds \n")
		}

		num, err := strconv.ParseFloat(strings.TrimSpace(r[1]), 64)
		if err != nil {
			return err
		}

		funds[i] = num
	}

	mean, err := stats.Mean(funds)
	if err != nil {
		return err
	}

	m.Contents = append(m.Contents, field{
		"# Funds Holding in Average for 4 Quarters",
		fmt.Sprintf("%v", int(mean)),
	})

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
