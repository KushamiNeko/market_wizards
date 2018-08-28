package analysis

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"config"
	"datautils"
	"encoding/base64"
	"fmt"
	"html/template"
	"ibd"
	"math"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/message"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func makePlot(
	title,
	xLabel,
	yLabel string,
	grid bool,
	setupCb func(p *plot.Plot),
) (*plot.Plot, error) {

	p, err := plot.New()
	if err != nil {
		return nil, err
	}

	p.Title.Text = title

	if xLabel != "" {
		p.X.Label.Text = xLabel
	}

	if yLabel != "" {
		p.Y.Label.Text = yLabel
	}

	p.Title.Font.Size = vg.Points(config.ChartFontSizeL)

	p.X.Label.Font.Size = vg.Points(config.ChartFontSizeM)
	p.Y.Label.Font.Size = vg.Points(config.ChartFontSizeM)

	p.X.Tick.Label.Font.Size = vg.Points(config.ChartFontSizeS)
	p.Y.Tick.Label.Font.Size = vg.Points(config.ChartFontSizeS)

	if setupCb != nil {
		setupCb(p)
	}

	err = p.Title.Font.SetName(config.ChartFont)
	if err != nil {
		return nil, err
	}

	err = p.X.Label.Font.SetName(config.ChartFont)
	if err != nil {
		return nil, err
	}

	err = p.Y.Label.Font.SetName(config.ChartFont)
	if err != nil {
		return nil, err
	}

	err = p.X.Tick.Label.Font.SetName(config.ChartFont)
	if err != nil {
		return nil, err
	}

	err = p.Y.Tick.Label.Font.SetName(config.ChartFont)
	if err != nil {
		return nil, err
	}

	err = p.Legend.Font.SetName(config.ChartFont)
	if err != nil {
		return nil, err
	}

	if grid {
		p.Add(plotter.NewGrid())
	}

	p.Legend.Font.Size = vg.Points(config.ChartFontSizeS)
	p.Legend.YAlign = draw.YBottom
	p.Legend.TextStyle.YAlign = draw.YBottom
	p.Legend.Top = true

	return p, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func plotToDataUrl(p *plot.Plot) (template.URL, error) {

	writer, err := p.WriterTo(vg.Points(config.ChartWidth), vg.Points(config.ChartHeight), "png")
	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)

	_, err = writer.WriteTo(buffer)
	if err != nil {
		return "", err
	}

	encode := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return template.URL(fmt.Sprintf(config.ChartDataUrlFormat, encode)), nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func makeValueSlice(
	winners []interface{},
	losers []interface{},
	labelCb func(o interface{}) (interface{}, error),
	keysSortCb func(keys []interface{}),
	keyFormatCb func(key interface{}) string,
) ([]string, []float64, []float64, float64, float64, error) {

	dictW := make(map[interface{}]int)
	dictL := make(map[interface{}]int)

	for _, o := range winners {

		grps, err := labelCb(o)
		if err != nil {
			return nil, nil, nil, 0.0, 0.0, err
		}

		if val, ok := dictW[grps]; ok {
			dictW[grps] = val + 1
		} else {
			dictW[grps] = 1
		}
	}

	for _, o := range losers {

		grps, err := labelCb(o)
		if err != nil {
			return nil, nil, nil, 0.0, 0.0, err
		}

		if val, ok := dictL[grps]; ok {
			dictL[grps] = val + 1
		} else {
			dictL[grps] = 1
		}
	}

	ck := make([]interface{}, 0)

	for k, _ := range dictW {
		ck = append(ck, k)
	}

outer:
	for k, _ := range dictL {
		for _, c := range ck {
			if c == k {
				continue outer
			}
		}

		ck = append(ck, k)
	}

	keysSortCb(ck)

	keys := make([]string, len(ck))

	winnersG := make([]float64, len(ck))
	losersG := make([]float64, len(ck))

	wmax := 0.0
	lmax := 0.0

	for i, c := range ck {
		if vw, ok := dictW[c]; ok {
			winnersG[i] = float64(vw)
			if float64(vw) > wmax {
				wmax = float64(vw)
			}
		} else {
			winnersG[i] = 0.0
		}
		if vl, ok := dictL[c]; ok {
			losersG[i] = float64(vl)
			if float64(vl) > lmax {
				lmax = float64(vl)
			}
		} else {
			losersG[i] = 0.0
		}

		keys[i] = keyFormatCb(c)
	}

	return keys, winnersG, losersG, wmax, lmax, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func makeBarCharts(
	p *plot.Plot,
	keys []string,
	winnersG,
	losersG []float64,
	wmax,
	lmax float64,
) error {

	width := vg.Points(config.ChartBarWidth)

	wb, err := plotter.NewBarChart(plotter.Values(winnersG), width)
	if err != nil {
		return err
	}

	wb.LineStyle.Width = vg.Length(0)
	wb.Color = config.WinnerRGBA
	wb.Offset = -width / 2

	lb, err := plotter.NewBarChart(plotter.Values(losersG), width)
	if err != nil {
		return err
	}

	lb.LineStyle.Width = vg.Length(0)
	lb.Color = config.LoserRGBA
	lb.Offset = width / 2

	p.Add(wb, lb)

	max := 0.0

	if wmax > lmax {
		max = wmax
	} else {
		max = lmax
	}

	pts := make(plotter.XYs, len(winnersG))
	//sma := make(plotter.XYs, len(winnersG))

	for i, _ := range winnersG {

		batting := winnersG[i] / (winnersG[i] + losersG[i])
		batting = batting * max

		pts[i] = struct{ X, Y float64 }{
			float64(i),
			batting,
		}

		//// Calculate Simple Moving Average
		//smaf := make([]float64, len(winnersG))
		//smab := make([]float64, len(winnersG))

		//if i < config.StatisticSMA-1.0 {
		////sma[i] = struct{ X, Y float64 }{
		////float64(i),
		////batting,
		////}

		//smaf[i] = batting

		//} else {

		//bs := make([]float64, config.StatisticSMA)

		//for j := 0; j < config.StatisticSMA; j++ {
		//b := winnersG[i-j] / (winnersG[i-j] + losersG[i-j])
		//b = b * max
		//bs[j] = b
		//}

		//s := stat.Mean(bs, nil)
		//smaf[i] = s
		//}

		//if i > len(winnersG)-config.StatisticSMA-1.0 {
		//smab[i] = batting
		//} else {
		//bs := make([]float64, config.StatisticSMA)

		//for j := 0; j < config.StatisticSMA; j++ {
		//b := winnersG[i+j] / (winnersG[i+j] + losersG[i+j])
		//b = b * max
		//bs[j] = b
		//}

		//s := stat.Mean(bs, nil)
		//smab[i] = s
		//}

		//sma[i] = struct{ X, Y float64 }{
		//float64(i),
		//stat.Mean([]float64{
		//smaf[i],
		//smab[i],
		//}, nil),
		//}
		//// Calculate Simple Moving Average

	}

	line, points, err := plotter.NewLinePoints(pts)
	if err != nil {
		return err
	}

	line.Color = config.InfoRGBA
	line.LineStyle.Width = vg.Points(config.ChartLineWidth)

	points.Shape = draw.CircleGlyph{}
	points.Radius = vg.Points(config.ChartPointRadius)
	points.Color = config.InfoRGBA

	p.Add(line, points)

	//smaLine, smaPoints, err := plotter.NewLinePoints(sma)
	//if err != nil {
	//return err
	//}

	//smaLine.Color = config.SMARGBA
	//smaPoints.Shape = draw.CircleGlyph{}
	//smaPoints.Radius = vg.Points(config.ChartPointRadius)
	//smaPoints.Color = config.SMARGBA

	//p.Add(smaLine, smaPoints)

	p.Y.Max = math.Max(wmax, lmax) * config.ChartLegendPaddingYRatio

	//p.X.Min = -(p.X.Max * float64(config.ChartBarPaddingXRatio-1.0))
	//p.X.Max = p.X.Max * float64(config.ChartBarPaddingXRatio)
	p.X.Min = -(p.X.Max * (float64(config.ChartBarPaddingXRatio) - 1.0) / 2.0)
	p.X.Max = p.X.Max * (((float64(config.ChartBarPaddingXRatio) - 1.0) / 2.0) + 1.0)

	p.Legend.Add("winners", wb)
	p.Legend.Add("losers", lb)
	p.Legend.Add("batting average", line, points)

	p.NominalX(keys...)

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func barChartSetup(p *plot.Plot) {
	p.X.Padding = vg.Points(config.ChartLabelPaddingX)

	p.X.Tick.Label.XAlign = draw.XLeft
	p.X.Tick.Label.YAlign = draw.YCenter

	p.Y.Tick.Label.XAlign = draw.XRight
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func barChartPercent(
	label string,
	interval float64,
	winners,
	losers []interface{},
) (template.URL, error) {

	p, err := makePlot(
		label,
		"",
		"",
		true,
		func(p *plot.Plot) {
			barChartSetup(p)
			p.X.Tick.Label.Rotation = config.ChartLabelRotation
		},
	)
	if err != nil {
		return "", err
	}

	keys, winnersG, losersG, wmax, lmax, err :=
		makeValueSlice(
			winners,
			losers,
			func(o interface{}) (interface{}, error) {

				t := o.(*ibd.IBDCheckup)
				value := t.Contents[label]

				var grps int

				if value == config.NullValue {
					grps = math.MaxInt32
				} else {
					vf, err := strconv.ParseFloat(strings.Replace(value, "%", "", -1), 64)
					if err != nil {
						return "", err
					}

					grp := math.Floor(vf / interval)
					grps = int(grp * interval)
				}

				return grps, nil
			},
			func(keys []interface{}) {
				sort.Slice(keys, func(i, j int) bool {
					is := keys[i].(int)
					js := keys[j].(int)

					return is < js
				})
			},
			func(key interface{}) string {

				k := key.(int)

				p := message.NewPrinter(message.MatchLanguage("en"))
				var grpk string

				if k == math.MaxInt32 {
					grpk = config.NullValue
				} else {
					grp := math.Floor(float64(k) / interval)
					grpk = p.Sprintf(config.PercentIntervalFormat, int(grp*interval), int((grp+1)*interval))
				}

				return grpk
			},
		)
	if err != nil {
		return "", err
	}

	err = makeBarCharts(p, keys, winnersG, losersG, wmax, lmax)
	if err != nil {
		return "", err
	}

	str, err := plotToDataUrl(p)

	if err != nil {
		return "", err
	}

	return str, nil

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func barChartStringRank(
	label string,
	winners,
	losers []interface{},
) (template.URL, error) {

	p, err := makePlot(
		label,
		"",
		"",
		true,
		func(p *plot.Plot) {
			barChartSetup(p)
		},
	)
	if err != nil {
		return "", err
	}

	keys, winnersG, losersG, wmax, lmax, err :=
		makeValueSlice(
			winners,
			losers,
			func(o interface{}) (interface{}, error) {

				t := o.(*ibd.IBDCheckup)

				value := t.Contents[label]

				var grps string
				grps = value
				grps = strings.Replace(grps, "+", "", -1)
				grps = strings.Replace(grps, "-", "", -1)

				return grps, nil
			},
			func(keys []interface{}) {
				sort.Slice(keys, func(i, j int) bool {
					is := keys[i].(string)
					js := keys[j].(string)

					return is < js
				})
			},
			func(key interface{}) string {
				k := key.(string)
				return k
			},
		)
	if err != nil {
		return "", err
	}

	err = makeBarCharts(p, keys, winnersG, losersG, wmax, lmax)
	if err != nil {
		return "", err
	}

	str, err := plotToDataUrl(p)

	if err != nil {
		return "", err
	}

	return str, nil

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func barChartIntInterval(
	label string,
	interval float64,
	winners,
	losers []interface{},
) (template.URL, error) {

	p, err := makePlot(
		label,
		"",
		"",
		true,
		func(p *plot.Plot) {
			barChartSetup(p)
			p.X.Tick.Label.Rotation = config.ChartLabelRotation
		},
	)
	if err != nil {
		return "", err
	}

	keys, winnersG, losersG, wmax, lmax, err :=
		makeValueSlice(
			winners,
			losers,
			func(o interface{}) (interface{}, error) {

				t := o.(*ibd.IBDCheckup)
				value := t.Contents[label]

				var grps int

				if value == config.NullValue {
					grps = math.MaxInt32
				} else {

					vf, err := strconv.ParseFloat(value, 64)
					if err != nil {
						return "", err
					}

					grp := math.Floor(vf / interval)
					grps = int(grp * interval)

				}

				return grps, nil
			},
			func(keys []interface{}) {
				sort.Slice(keys, func(i, j int) bool {
					is := keys[i].(int)
					js := keys[j].(int)

					return is < js
				})
			},
			func(key interface{}) string {

				k := key.(int)

				p := message.NewPrinter(message.MatchLanguage("en"))

				var grpk string

				if k == math.MaxInt32 {
					grpk = config.NullValue
				} else {
					grp := math.Floor(float64(k) / interval)
					grpk = p.Sprintf(config.IntervalFormat, int64(grp*interval), int64((grp+1)*interval))
				}

				return grpk
			},
		)
	if err != nil {
		return "", err
	}

	err = makeBarCharts(p, keys, winnersG, losersG, wmax, lmax)
	if err != nil {
		return "", err
	}

	str, err := plotToDataUrl(p)

	if err != nil {
		return "", err
	}

	return str, nil

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func barChartFloatInterval(
	label string,
	interval float64,
	winners,
	losers []interface{},
) (template.URL, error) {

	p, err := makePlot(
		label,
		"",
		"",
		true,
		func(p *plot.Plot) {
			barChartSetup(p)
			p.X.Tick.Label.Rotation = config.ChartLabelRotation
		},
	)
	if err != nil {
		return "", err
	}

	keys, winnersG, losersG, wmax, lmax, err :=
		makeValueSlice(
			winners,
			losers,
			func(o interface{}) (interface{}, error) {

				t := o.(*ibd.IBDCheckup)
				value := t.Contents[label]

				var grps float64

				if value == config.NullValue {
					grps = math.MaxFloat64
				} else {

					vf, err := strconv.ParseFloat(value, 64)
					if err != nil {
						return "", err
					}

					grp := math.Floor(vf / interval)
					grps = float64(grp * interval)

				}

				return grps, nil
			},
			func(keys []interface{}) {
				sort.Slice(keys, func(i, j int) bool {
					is := keys[i].(float64)
					js := keys[j].(float64)

					return is < js
				})
			},
			func(key interface{}) string {

				k := key.(float64)

				p := message.NewPrinter(message.MatchLanguage("en"))
				var grpk string

				if k == math.MaxFloat64 {
					grpk = config.NullValue
				} else {
					grp := math.Floor(k / interval)
					grpk = p.Sprintf(config.IntervalFormat, grp*interval, (grp+1)*interval)
				}

				return grpk
			},
		)
	if err != nil {
		return "", err
	}

	err = makeBarCharts(p, keys, winnersG, losersG, wmax, lmax)
	if err != nil {
		return "", err
	}

	str, err := plotToDataUrl(p)

	if err != nil {
		return "", err
	}

	return str, nil

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func barChartInt(
	label string,
	winners,
	losers []interface{},
) (template.URL, error) {

	p, err := makePlot(
		label,
		"",
		"",
		true,
		func(p *plot.Plot) {
			barChartSetup(p)
		},
	)
	if err != nil {
		return "", err
	}

	keys, winnersG, losersG, wmax, lmax, err :=
		makeValueSlice(
			winners,
			losers,
			func(o interface{}) (interface{}, error) {
				t := o.(*ibd.IBDCheckup)
				value := t.Contents[label]

				var grps int

				if value == config.NullValue {
					grps = math.MaxInt32

				} else {
					vf, err := strconv.ParseFloat(value, 64)
					if err != nil {
						return "", err
					}

					grps = int(vf)
				}

				return grps, nil
			},
			func(keys []interface{}) {
				sort.Slice(keys, func(i, j int) bool {
					is := keys[i].(int)
					js := keys[j].(int)

					return is < js
				})
			},
			func(key interface{}) string {
				k := key.(int)

				p := message.NewPrinter(message.MatchLanguage("en"))

				var grpk string

				if k == math.MaxInt32 {
					grpk = config.NullValue
				} else {
					grpk = p.Sprintf("%v", k)
				}

				return grpk
			},
		)
	if err != nil {
		return "", err
	}

	err = makeBarCharts(p, keys, winnersG, losersG, wmax, lmax)
	if err != nil {
		return "", err
	}

	str, err := plotToDataUrl(p)

	if err != nil {
		return "", err
	}

	return str, nil

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func barChartString(
	label string,
	winners,
	losers []interface{},
) (template.URL, error) {

	p, err := makePlot(
		label,
		"",
		"",
		true,
		func(p *plot.Plot) {
			barChartSetup(p)
		},
	)
	if err != nil {
		return "", err
	}

	keys, winnersG, losersG, wmax, lmax, err :=
		makeValueSlice(
			winners,
			losers,
			func(o interface{}) (interface{}, error) {
				t := o.(*ibd.IBDCheckup)
				value := t.Contents[label]

				return value, nil
			},
			func(keys []interface{}) {
				sort.Slice(keys, func(i, j int) bool {
					is := keys[i].(string)
					js := keys[j].(string)

					return is < js
				})
			},
			func(key interface{}) string {
				k := key.(string)
				return k
			},
		)
	if err != nil {
		return "", err
	}

	err = makeBarCharts(p, keys, winnersG, losersG, wmax, lmax)
	if err != nil {
		return "", err
	}

	str, err := plotToDataUrl(p)

	if err != nil {
		return "", err
	}

	return str, nil

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func makeGoogleChartsDataStruct(
	chartLabel string,
	keys []string,
	winners []float64,
	losers []float64,
) (string, error) {

	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		chartLabel,
		"Winner",
		map[string]string{
			"role": "style",
		},
		"Loser",
		map[string]string{
			"role": "style",
		},
	})

	for i, k := range keys {
		g = append(g, []interface{}{
			k,
			winners[i],
			fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
			losers[i],
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
