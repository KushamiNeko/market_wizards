package marketsmith

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"config"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	//"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	regexInfoCell = `\s*<div\s*class="cell">\s*<div\s*class="key">([^<]+)<\/div>\s*<div\s*class="value">([^<]+)<\/div>\s*<\/div>\s*`

	regexIndustryGroup = `\s*<div\s*class="companySymbol">\s*<span\s*class="companyInfCoName">[^<]+<\/span>\(\w+\)\s*\S+\s*([^<]+)<\/div>\s*`

	regexOptions = `\s*<div\s*class="Options">Options\s*([^<]+)<\/div>\s*`

	regexFloatShare = `\s*<div\s*class="companyInfoCenter"\s*[^>]+>\s*<div\s*class="companyInfoLable">\s*<div\s*class="cell">([^<]+)<\/div>\s*<div\s*class="cell">([^<]+)<\/div>\s*<div\s*class="cell">([^<]+)<\/div>\s*<\/div>\s*<div\s*class="companyInfoVlaue">\s*<div\s*class="cell">([^<]+)<\/div>\s*<div\s*class="cell">([^<]+)<\/div>\s*<div\s*class="cell">([^<]+)<\/div>\s*<\/div>\s*<\/div>\s*`

	regexQuarterlyResults = `\s*<div\s*class="cell"\s*style="display:\s*block;">\s*<div\s*class="quarterly\s*\w*">[^<]+<span\s*[^>]+>\s*[^<]*\s*<\/span>\s*<\/div>\s*<div\s*class="eps\s*\w*">\s*([^<]+)\s*<\/div>\s*<div\s*class="epsChg\s*\w*">([^<]+)<\/div>\s*<div\s*class="sales\s*\w*">([^<]+)<\/div>\s*<div\s*class="salesChg\s*\w*">([^<]+)<\/div>\s*<\/div>\s*`

	regexFunds = `\s*<div\s*class="cell">\s*<span>[^<]+<\/span>\s*<span>(\d+)<\/span>\s*<\/div>\s*`

	regexPercent = `\s*[+]*([-0-9.]+)%\s*`

	regexRSRating = `\s*<g\s*stroke="none"\s*stroke-width="[\d.]+"\s*fill="rgb\(\d+,\d+,\d+\)"\s*font-family="Arial"\s*font-size="11pt"\s*clip-path="url\(#RSView_ClipPath\)">\s*<text\s*x="[\d.]+"\s*y="[\d.]+"\s*text-anchor="start">\s*(\d+)\s*<\/text>\s*<\/g>\s*`
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	reInfoCell *regexp.Regexp

	reIndustryGroup *regexp.Regexp
	reOptions       *regexp.Regexp

	reFloatShare       *regexp.Regexp
	reQuarterlyResults *regexp.Regexp

	reFunds *regexp.Regexp

	rePercent *regexp.Regexp

	reRSRating *regexp.Regexp
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	reInfoCell = regexp.MustCompile(regexInfoCell)

	reIndustryGroup = regexp.MustCompile(regexIndustryGroup)
	reOptions = regexp.MustCompile(regexOptions)

	reFloatShare = regexp.MustCompile(regexFloatShare)
	reQuarterlyResults = regexp.MustCompile(regexQuarterlyResults)

	reFunds = regexp.MustCompile(regexFunds)

	rePercent = regexp.MustCompile(regexPercent)

	reRSRating = regexp.MustCompile(regexRSRating)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type MarketSmith struct {
	//Contents []*field
	Contents map[string]string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func MarketSmithNew() *MarketSmith {
	m := new(MarketSmith)
	//m.Contents = make([]*field, 0)
	m.Contents = make(map[string]string)

	return m
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func (m *MarketSmith) GetContents() []datautils.Fields {

//f := make([]datautils.Fields, len(m.Contents))

//for i, c := range m.Contents {
//f[i] = c
//}

//return f
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//type field struct {
//Label string
//Value string
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func (f *field) GetLabel() string {
//return f.Label
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func (f *field) GetValue() string {
//return f.Value
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func (f *field) SetValue(value string) {
//f.Value = value
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Parse(buffer *bytes.Buffer) (*MarketSmith, error) {

	m := MarketSmithNew()

	var err error

	err = m.getRSRating(buffer)
	if err != nil {
		return nil, err
	}

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

	_ = m.getFunds(buffer)
	//if err != nil {
	//return nil, err
	//}

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

		m.Contents[strings.TrimSpace(r[1])] = strings.TrimSpace(r[2])

		//m.Contents = append(m.Contents, &field{
		//strings.TrimSpace(r[1]),
		//strings.TrimSpace(r[2]),
		//})
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

		m.Contents["Industry Group"] = strings.TrimSpace(r[1])

		//m.Contents = append(m.Contents, &field{
		//"Industry Group",
		//strings.TrimSpace(r[1]),
		//})
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

		m.Contents["Options"] = strings.TrimSpace(r[1])

		//m.Contents = append(m.Contents, &field{
		//"Options",
		//strings.TrimSpace(r[1]),
		//})
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

		m.Contents[strings.TrimSpace(r[1])] = strings.TrimSpace(r[4])
		m.Contents[strings.TrimSpace(r[2])] = strings.TrimSpace(r[5])
		m.Contents[strings.TrimSpace(r[3])] = strings.TrimSpace(r[6])

		//m.Contents = append(m.Contents, &field{
		//strings.TrimSpace(r[1]),
		//strings.TrimSpace(r[4]),
		//})

		//m.Contents = append(m.Contents, &field{
		//strings.TrimSpace(r[2]),
		//strings.TrimSpace(r[5]),
		//})

		//m.Contents = append(m.Contents, &field{
		//strings.TrimSpace(r[3]),
		//strings.TrimSpace(r[6]),
		//})
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketSmith) getQuarterlyResults(buffer *bytes.Buffer) error {

	var results [][]string

	results = reQuarterlyResults.FindAllStringSubmatch(buffer.String(), -1)
	if results == nil {
		return fmt.Errorf("no matching found in the Quarterly Results\n")
	}

	epss := make([]string, len(results))
	saless := make([]string, len(results))

	for i, r := range results {
		if len(r) < 5 {
			return fmt.Errorf("problems occur while parsing Quarterly Results \n")
		}

		epss[i] = strings.TrimSpace(r[2])
		saless[i] = strings.TrimSpace(r[4])
	}

	m.Contents["Avg EPS % Chg 2Q"] = m.getQuarterlyMean(epss[5:])

	//m.Contents = append(m.Contents, &field{
	//"Avg EPS % Chg 2Q",
	//m.getQuarterlyMean(epss[5:]),
	//})

	m.Contents["Avg EPS % Chg 3Q"] = m.getQuarterlyMean(epss[4:])

	//m.Contents = append(m.Contents, &field{
	//"Avg EPS % Chg 3Q",
	//m.getQuarterlyMean(epss[4:]),
	//})

	m.Contents["Avg EPS % Chg 4Q"] = m.getQuarterlyMean(epss[3:])

	//m.Contents = append(m.Contents, &field{
	//"Avg EPS % Chg 4Q",
	//m.getQuarterlyMean(epss[3:]),
	//})

	m.Contents["Avg EPS % Chg 5Q"] = m.getQuarterlyMean(epss[2:])

	//m.Contents = append(m.Contents, &field{
	//"Avg EPS % Chg 5Q",
	//m.getQuarterlyMean(epss[2:]),
	//})

	m.Contents["Avg EPS % Chg 6Q"] = m.getQuarterlyMean(epss[1:])

	//m.Contents = append(m.Contents, &field{
	//"Avg EPS % Chg 6Q",
	//m.getQuarterlyMean(epss[1:]),
	//})

	m.Contents["Avg Sales % Chg 2Q"] = m.getQuarterlyMean(epss[5:])

	//m.Contents = append(m.Contents, &field{
	//"Avg Sales % Chg 2Q",
	//m.getQuarterlyMean(saless[5:]),
	//})

	m.Contents["Avg Sales % Chg 3Q"] = m.getQuarterlyMean(epss[4:])

	//m.Contents = append(m.Contents, &field{
	//"Avg Sales % Chg 3Q",
	//m.getQuarterlyMean(saless[4:]),
	//})

	m.Contents["Avg Sales % Chg 4Q"] = m.getQuarterlyMean(epss[3:])

	//m.Contents = append(m.Contents, &field{
	//"Avg Sales % Chg 4Q",
	//m.getQuarterlyMean(saless[3:]),
	//})

	m.Contents["Avg Sales % Chg 5Q"] = m.getQuarterlyMean(epss[2:])

	//m.Contents = append(m.Contents, &field{
	//"Avg Sales % Chg 5Q",
	//m.getQuarterlyMean(saless[2:]),
	//})

	m.Contents["Avg Sales % Chg 6Q"] = m.getQuarterlyMean(epss[1:])

	//m.Contents = append(m.Contents, &field{
	//"Avg Sales % Chg 6Q",
	//m.getQuarterlyMean(saless[1:]),
	//})

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketSmith) getQuarterlyMean(list []string) string {

	eps := make([]float64, len(list))

	for i, e := range list {
		if strings.Compare(e, config.NullValue) == 0 {
			return config.NullValue
		}

		es := rePercent.FindStringSubmatch(e)
		if es == nil {
			fmt.Println("Problesm in Regex for Parsing Quarterly Earnings")
			return config.NullValue
		}

		f, err := strconv.ParseFloat(es[1], 64)
		if err != nil {
			fmt.Println("Problesm in Regex for Parsing Quarterly Earnings")
			return config.NullValue
		}

		eps[i] = f
	}

	mean := stat.Mean(eps, nil)
	//mean, err := stats.Mean(eps)
	//if err != nil {
	//return err.Error()
	//}

	return fmt.Sprintf("%v", mean)
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

	mean := stat.Mean(funds, nil)
	//mean, err := stats.Mean(funds)
	//if err != nil {
	//return err
	//}

	m.Contents["Avg Funds Holding 4Q"] = fmt.Sprintf("%v", int(mean))

	//m.Contents = append(m.Contents, &field{
	//"Avg Funds Holding 4Q",
	//fmt.Sprintf("%v", int(mean)),
	//})

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketSmith) getRSRating(buffer *bytes.Buffer) error {

	var results [][]string

	results = reRSRating.FindAllStringSubmatch(buffer.String(), -1)
	if results == nil {
		return fmt.Errorf("no matching found in the RS Rating\n")
	}

	if len(results) > 1 {
		return fmt.Errorf("problems occur while parsing RS Rating\n")
	}

	if len(results[0]) < 2 {
		return fmt.Errorf("problems occur while parsing RS Rating\n")
	}

	m.Contents["RS Rating"] = strings.TrimSpace(results[0][1])

	//m.Contents = append(m.Contents, &field{
	//"RS Rating",
	//results[0][1],
	//})

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
