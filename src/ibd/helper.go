package ibd

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	none = "N/A"

	regexL = `\s*<a\s*class=\"glossDef\"[^>]+>\s*([^<]+)\s*<\/a>\s*`
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	reL *regexp.Regexp
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	reL = regexp.MustCompile(regexL)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func cleanup(results [][]string, row, col int) string {
	return strings.TrimSpace(results[row][col])
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func cleanupL(results [][]string, label string, col int) string {
	for _, r := range results {
		match := reL.FindStringSubmatch(r[0])
		if len(match) == 0 {
			return ""
		}

		if strings.Compare(strings.TrimSpace(match[1]), strings.TrimSpace(label)) == 0 {
			return strings.TrimSpace(r[col])
		}
	}

	return ""
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func parsePercent(str string) (float64, error) {
	if strings.Compare(str, none) == 0 {
		return math.MaxFloat64, nil
	}

	ss := rePercent.FindStringSubmatch(str)
	if ss == nil {
		return 0.0, fmt.Errorf("No Match for Percent\n")
	}

	f, err := strconv.ParseFloat(ss[1], 64)
	if err != nil {
		return 0.0, err
	}

	return f, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func parsePrice(str string) (float64, error) {
	if strings.Compare(str, none) == 0 {
		return math.MaxFloat64, nil
	}

	ss := rePrice.FindStringSubmatch(str)
	if ss == nil {
		return 0.0, fmt.Errorf("No Match for Price\n")
	}

	f, err := strconv.ParseFloat(ss[1], 64)
	if err != nil {
		return 0.0, err
	}

	return f, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func parseVolume(str string) (int64, error) {

	if strings.Compare(str, none) == 0 {
		return math.MaxInt64, nil
	}

	results := reVolume.FindStringSubmatch(strings.Replace(str, ",", "", -1))
	if results == nil {
		return 0, fmt.Errorf("No Match for Volume")
	}

	f, err := strconv.ParseFloat(results[1], 64)
	if err != nil {
		return 0, err
	}

	switch results[2] {
	case "Mil":
		return int64(f * 1000000.0), nil
	case "Bil":
		return int64(f * 1000000000.0), nil
	default:
		return int64(f), nil
	}

	return 0, fmt.Errorf("No Match for Volume")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func parseMktCap(str string) (int64, error) {
	if strings.Compare(str, none) == 0 {
		return math.MaxInt64, nil
	}

	ss := reMktCap.FindStringSubmatch(str)
	if ss == nil {
		return 0.0, fmt.Errorf("No Match for Market Capital\n")
	}

	f, err := strconv.ParseFloat(ss[1], 64)
	if err != nil {
		return 0.0, err
	}

	var cap int64

	switch ss[2] {
	case "M":
		cap = int64(f * 1000000.0)

	case "B":
		cap = int64(f * 1000000000.0)
	}

	return cap, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func parseInt(str string) (int, error) {
	if strings.Compare(str, none) == 0 {
		return math.MaxInt32, nil
	}

	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}

	return int(i), nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func parseFloat(str string) (float64, error) {
	if strings.Compare(str, none) == 0 {
		return math.MaxFloat64, nil
	}

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0, err
	}

	return f, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
