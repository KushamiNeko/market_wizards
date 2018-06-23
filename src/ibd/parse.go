package ibd

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"datautils"
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

type IBDCheckup struct {
	Contents []*field
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func IBDCheckupNew() *IBDCheckup {
	checkup := new(IBDCheckup)
	checkup.Contents = make([]*field, 0)

	return checkup
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (i *IBDCheckup) GetContents() []datautils.Fields {

	f := make([]datautils.Fields, len(i.Contents))

	for i, c := range i.Contents {
		f[i] = c
	}

	return f
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type field struct {
	Label     string
	Value     string
	Condition string `json:",omitempty"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (f *field) GetLabel() string {
	return f.Label
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (f *field) GetValue() string {
	return f.Value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (f *field) SetValue(value string) {
	f.Value = value
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

	checkup.Contents = append(checkup.Contents, &field{
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

			checkup.Contents = append(checkup.Contents, &field{
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

		checkup.Contents = append(checkup.Contents, &field{
			strings.TrimSpace(match[1]),
			value,
			condition,
		})
	}

	return checkup, nil

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
