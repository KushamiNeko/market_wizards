package charts

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"datautils"
	"fmt"
	"ibd"
	"math"
	"sort"
	"strconv"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func columnChartPercent(label string, winners, losers []*ibd.IBDCheckup) (string, error) {
	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		label,
		"Winner",
		map[string]string{
			"role": "style",
		},
		"Loser",
		map[string]string{
			"role": "style",
		},
	})

	var interval float64 = 5.0

	intervalDictW := make(map[int]int)
	intervalDictL := make(map[int]int)

	for _, o := range winners {
		for _, f := range o.Contents {
			if f.Label == label {
				var grps int

				if f.Value == config.NullValue {
					grps = math.MaxInt32
				} else {
					vf, err := strconv.ParseFloat(strings.Replace(f.Value, "%", "", -1), 64)
					if err != nil {
						return "", err
					}

					grp := math.Floor(vf / interval)
					grps = int(grp * interval)
				}

				if val, ok := intervalDictW[grps]; ok {
					intervalDictW[grps] = val + 1
				} else {
					intervalDictW[grps] = 1
				}

				break

			}
		}
	}

	for _, o := range losers {
		for _, f := range o.Contents {
			if f.Label == label {
				var grps int

				if f.Value == config.NullValue {
					grps = math.MaxInt32
				} else {

					vf, err := strconv.ParseFloat(strings.Replace(f.Value, "%", "", -1), 64)
					if err != nil {
						return "", err
					}

					grp := math.Floor(vf / interval)
					grps = int(grp * interval)
				}

				if val, ok := intervalDictL[grps]; ok {
					intervalDictL[grps] = val + 1
				} else {
					intervalDictL[grps] = 1
				}

				break
			}
		}
	}

	ck := make([]int, 0)

	for k, _ := range intervalDictW {
		ck = append(ck, k)
	}

outer:
	for k, _ := range intervalDictL {
		for _, c := range ck {
			if c == k {
				continue outer
			}
		}

		ck = append(ck, k)
	}

	sort.Ints(ck)

	for _, k := range ck {

		var vw int
		var vl int

		var grpk string

		if k == math.MaxInt32 {
			grpk = config.NullValue
		} else {
			grp := math.Floor(float64(k) / interval)
			grpk = fmt.Sprintf(config.PercentIntervalFormat, int(grp*interval), int((grp+1)*interval))
		}

		if v, ok := intervalDictW[k]; ok {
			vw = v
		} else {
			vw = 0
		}

		if v, ok := intervalDictL[k]; ok {
			vl = v
		} else {
			vl = 0
		}

		g = append(g, []interface{}{
			grpk,
			vw,
			fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
			vl,
			fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
		})
	}

	jg, err := datautils.JsonB64Encrypt(g)
	if err != nil {
		return "", err
	}

	return jg, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func columnChartStringRank(label string, winners, losers []*ibd.IBDCheckup) (string, error) {
	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		label,
		"Winner",
		map[string]string{
			"role": "style",
		},
		"Loser",
		map[string]string{
			"role": "style",
		},
	})

	intervalDictW := make(map[string]int)
	intervalDictL := make(map[string]int)

	for _, o := range winners {
		for _, f := range o.Contents {
			if f.Label == label {

				var grps string
				grps = f.Value
				grps = strings.Replace(grps, "+", "", -1)
				grps = strings.Replace(grps, "-", "", -1)

				if val, ok := intervalDictW[grps]; ok {
					intervalDictW[grps] = val + 1
				} else {
					intervalDictW[grps] = 1
				}

				break
			}
		}
	}

	for _, o := range losers {
		for _, f := range o.Contents {
			if f.Label == label {

				var grps string
				grps = f.Value
				grps = strings.Replace(grps, "+", "", -1)
				grps = strings.Replace(grps, "-", "", -1)

				if val, ok := intervalDictL[grps]; ok {
					intervalDictL[grps] = val + 1
				} else {
					intervalDictL[grps] = 1
				}

				break
			}
		}
	}

	ck := make([]string, 0)

	for k, _ := range intervalDictW {
		ck = append(ck, k)
	}

outer:
	for k, _ := range intervalDictL {
		for _, c := range ck {
			if c == k {
				continue outer
			}
		}

		ck = append(ck, k)
	}

	sort.Strings(ck)

	for _, k := range ck {

		var vw int
		var vl int

		if v, ok := intervalDictW[k]; ok {
			vw = v
		} else {
			vw = 0
		}

		if v, ok := intervalDictL[k]; ok {
			vl = v
		} else {
			vl = 0
		}

		g = append(g, []interface{}{
			k,
			vw,
			fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
			vl,
			fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
		})
	}

	jg, err := datautils.JsonB64Encrypt(g)
	if err != nil {
		return "", err
	}

	return jg, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func columnChartIntInterval(label string, interval float64, winners, losers []*ibd.IBDCheckup) (string, error) {
	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		label,
		"Winner",
		map[string]string{
			"role": "style",
		},
		"Loser",
		map[string]string{
			"role": "style",
		},
	})

	intervalDictW := make(map[int]int)
	intervalDictL := make(map[int]int)

	for _, o := range winners {
		for _, f := range o.Contents {
			if f.Label == label {
				var grps int

				if f.Value == config.NullValue {
					grps = math.MaxInt32
				} else {

					vf, err := strconv.ParseFloat(f.Value, 64)
					if err != nil {
						return "", err
					}

					grp := math.Floor(vf / interval)
					grps = int(grp * interval)
				}

				if val, ok := intervalDictW[grps]; ok {
					intervalDictW[grps] = val + 1
				} else {
					intervalDictW[grps] = 1
				}

				break
			}
		}
	}

	for _, o := range losers {
		for _, f := range o.Contents {
			if f.Label == label {
				var grps int

				if f.Value == config.NullValue {
					grps = math.MaxInt32
				} else {

					vf, err := strconv.ParseFloat(f.Value, 64)
					if err != nil {
						return "", err
					}

					grp := math.Floor(vf / interval)
					grps = int(grp * interval)
				}

				if val, ok := intervalDictL[grps]; ok {
					intervalDictL[grps] = val + 1
				} else {
					intervalDictL[grps] = 1
				}

				break
			}
		}
	}

	ck := make([]int, 0)

	for k, _ := range intervalDictW {
		ck = append(ck, k)
	}

outer:
	for k, _ := range intervalDictL {
		for _, c := range ck {
			if c == k {
				continue outer
			}
		}

		ck = append(ck, k)
	}

	sort.Ints(ck)

	for _, k := range ck {

		var vw int
		var vl int

		var grpk string

		if k == math.MaxInt32 {
			grpk = config.NullValue
		} else {
			grp := math.Floor(float64(k) / interval)
			grpk = fmt.Sprintf(config.IntervalFormat, int(grp*interval), int((grp+1)*interval))
		}

		if v, ok := intervalDictW[k]; ok {
			vw = v
		} else {
			vw = 0
		}

		if v, ok := intervalDictL[k]; ok {
			vl = v
		} else {
			vl = 0
		}

		g = append(g, []interface{}{
			grpk,
			vw,
			fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
			vl,
			fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
		})
	}

	jg, err := datautils.JsonB64Encrypt(g)
	if err != nil {
		return "", err
	}

	return jg, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func columnChartFloatInterval(label string, interval float64, winners, losers []*ibd.IBDCheckup) (string, error) {
	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		label,
		"Winner",
		map[string]string{
			"role": "style",
		},
		"Loser",
		map[string]string{
			"role": "style",
		},
	})

	intervalDictW := make(map[float64]int)
	intervalDictL := make(map[float64]int)

	for _, o := range winners {
		for _, f := range o.Contents {
			if f.Label == label {
				var grps float64

				if f.Value == config.NullValue {
					grps = math.MaxFloat64
				} else {

					vf, err := strconv.ParseFloat(f.Value, 64)
					if err != nil {
						return "", err
					}

					grp := math.Floor(vf / interval)
					grps = float64(grp * interval)
				}

				if val, ok := intervalDictW[grps]; ok {
					intervalDictW[grps] = val + 1
				} else {
					intervalDictW[grps] = 1
				}

				break
			}
		}
	}

	for _, o := range losers {
		for _, f := range o.Contents {
			if f.Label == label {
				var grps float64

				if f.Value == config.NullValue {
					grps = math.MaxFloat64
				} else {

					vf, err := strconv.ParseFloat(f.Value, 64)
					if err != nil {
						return "", err
					}

					grp := math.Floor(vf / interval)
					grps = float64(grp * interval)
				}

				if val, ok := intervalDictL[grps]; ok {
					intervalDictL[grps] = val + 1
				} else {
					intervalDictL[grps] = 1
				}

				break
			}
		}
	}

	ck := make([]float64, 0)

	for k, _ := range intervalDictW {
		ck = append(ck, k)
	}

outer:
	for k, _ := range intervalDictL {
		for _, c := range ck {
			if c == k {
				continue outer
			}
		}

		ck = append(ck, k)
	}

	sort.Float64s(ck)

	for _, k := range ck {

		var vw int
		var vl int

		var grpk string

		if k == math.MaxFloat64 {
			grpk = config.NullValue
		} else {
			grp := math.Floor(float64(k) / interval)
			grpk = fmt.Sprintf(config.IntervalFormat, grp*interval, (grp+1)*interval)
		}

		if v, ok := intervalDictW[k]; ok {
			vw = v
		} else {
			vw = 0
		}

		if v, ok := intervalDictL[k]; ok {
			vl = v
		} else {
			vl = 0
		}

		g = append(g, []interface{}{
			grpk,
			vw,
			fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
			vl,
			fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
		})
	}

	jg, err := datautils.JsonB64Encrypt(g)
	if err != nil {
		return "", err
	}

	return jg, nil
}
