package ibd

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"strconv"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func parseArrow(str string) (string, error) {
	if strings.Compare(str, none) == 0 {
		return "", fmt.Errorf("The value is N/A\n")
	}

	ss := reArrow.FindStringSubmatch(str)
	if ss == nil {
		return "", fmt.Errorf("No Match for Arrow\n")
	}

	return ss[1], nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func parsePercent(str string) (float64, error) {
	if strings.Compare(str, none) == 0 {
		//return math.MaxFloat64, nil
		return 0.0, fmt.Errorf("The value is N/A\n")
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
		//return math.MaxFloat64, nil
		return 0.0, fmt.Errorf("The value is N/A\n")
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
		//return math.MaxInt64, nil
		return 0, fmt.Errorf("The value is N/A\n")
	}

	results := reVolume.FindStringSubmatch(strings.Replace(str, ",", "", -1))
	if results == nil {
		return 0, fmt.Errorf("No Match for Volume\n")
	}

	if reFloat.MatchString(results[1]) && results[2] == "" {
		return 0, fmt.Errorf("No Match for Volume\n")
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

	return 0, fmt.Errorf("No Match for Volume\n")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func parseMktCap(str string) (int64, error) {
	if strings.Compare(str, none) == 0 {
		//return math.MaxInt64, nil
		return 0, fmt.Errorf("The value is N/A\n")
	}

	ss := reMktCap.FindStringSubmatch(str)
	if ss == nil {
		return 0, fmt.Errorf("No Match for Market Capital\n")
	}

	f, err := strconv.ParseFloat(ss[1], 64)
	if err != nil {
		return 0, err
	}

	var cap int64

	switch ss[2] {
	case "M":
		cap = int64(f * 1000000.0)
	case "B":
		cap = int64(f * 1000000000.0)
	default:
		return 0, fmt.Errorf("No Match for Market Capital\n")
	}

	return cap, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
