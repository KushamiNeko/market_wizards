package charts

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"config"
	"encoding/base64"
	"fmt"
	"html/template"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"transaction"

	"golang.org/x/text/message"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ChartGeneral struct {
	filterOrders []*transaction.Transaction

	winners []*transaction.Transaction
	losers  []*transaction.Transaction

	//GainVsDaysHeld string
	GainVsDaysHeld template.URL
	//BuyPoints      string
	BuyPoints template.URL
	//PriceInterval string
	PriceInterval template.URL
	//Stage         string
	Stage template.URL
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var wg sync.WaitGroup

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ChartGeneralNew(filterOrders, winners, losers []*transaction.Transaction) (*ChartGeneral, error) {

	c := new(ChartGeneral)

	c.filterOrders = filterOrders

	c.winners = winners
	c.losers = losers

	//var err error
	errs := make([]error, 4)

	wg.Add(4)

	go func() {
		err := c.getGainVsDaysHeld()
		if err != nil {
			//return nil, err
			errs[0] = err
		}
		wg.Done()
	}()

	go func() {
		err := c.getBuyPoints()
		if err != nil {
			//return nil, err
			errs[1] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getPriceInterval()
		if err != nil {
			//return nil, err
			errs[2] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getStage()
		if err != nil {
			//return nil, err
			errs[3] = err
		}

		wg.Done()
	}()

	wg.Wait()

	for _, e := range errs {
		if e != nil {
			return nil, e
		}
	}

	return c, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getGainVsDaysHeld() error {

	p, err := newPlot(
		"Days Held vs Gain(%)",
		"Days Held",
		"Gain(%)",
		true,
		nil,
	)
	if err != nil {
		return err
	}

	//p, err := plot.New()
	//if err != nil {
	//return err
	//}

	//p.Title.Text = "Days Held vs Gain(%)"
	//p.X.Label.Text = "Days Held"
	//p.Y.Label.Text = "Gain(%)"

	//p.Title.Font.Size = vg.Points(config.ChartFontSizeL)
	//p.X.Label.Font.Size = vg.Points(config.ChartFontSizeM)
	//p.Y.Label.Font.Size = vg.Points(config.ChartFontSizeM)

	//p.X.Tick.Label.Font.Size = vg.Points(config.ChartFontSizeS)
	//p.Y.Tick.Label.Font.Size = vg.Points(config.ChartFontSizeS)

	//err = p.Title.Font.SetName(config.ChartFont)
	//if err != nil {
	//return err
	//}

	//err = p.X.Label.Font.SetName(config.ChartFont)
	//if err != nil {
	//return err
	//}

	//err = p.Y.Label.Font.SetName(config.ChartFont)
	//if err != nil {
	//return err
	//}

	//err = p.X.Tick.Label.Font.SetName(config.ChartFont)
	//if err != nil {
	//return err
	//}

	//err = p.Y.Tick.Label.Font.SetName(config.ChartFont)
	//if err != nil {
	//return err
	//}

	//p.Add(plotter.NewGrid())

	max := 0.0

	pts := make(plotter.XYs, 0)

	for _, o := range c.winners {
		pts = append(pts, struct{ X, Y float64 }{
			float64(o.Sell.DaysHeld),
			float64(o.Sell.GainP),
		})

		if float64(o.Sell.GainP) > max {
			max = float64(o.Sell.GainP)
		}
	}

	ws, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}

	ws.GlyphStyle.Color = config.WinnerRGBA
	ws.GlyphStyle.Radius = vg.Points(config.ChartPointRadius)
	ws.GlyphStyle.Shape = draw.CircleGlyph{}

	pts = make(plotter.XYs, 0)

	for _, o := range c.losers {
		pts = append(pts, struct{ X, Y float64 }{
			float64(o.Sell.DaysHeld),
			float64(o.Sell.GainP),
		})
	}

	ls, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}

	ls.GlyphStyle.Color = config.LoserRGBA
	ls.GlyphStyle.Radius = vg.Points(config.ChartPointRadius)
	ls.GlyphStyle.Shape = draw.CircleGlyph{}

	p.Add(ws, ls)
	p.Y.Max = max * config.ChartLegendPaddingYRatio

	p.Legend.Add("winners", ws)
	p.Legend.Add("losers", ls)
	//p.Legend.Font.Size = vg.Points(config.ChartFontSizeS)
	//p.Legend.YAlign = draw.YBottom
	//p.Legend.TextStyle.YAlign = draw.YBottom

	//writer, err := p.WriterTo(vg.Points(config.ChartWidth), vg.Points(config.ChartHeight), "png")
	//if err != nil {
	//return err
	//}

	//buffer := new(bytes.Buffer)

	//_, err = writer.WriteTo(buffer)
	//if err != nil {
	//return err
	//}

	//encode := base64.StdEncoding.EncodeToString(buffer.Bytes())

	//c.GainVsDaysHeld = template.URL(fmt.Sprintf(config.ChartDataUrlFormat, encode))
	c.GainVsDaysHeld, err = plotToDataUrl(p)
	if err != nil {
		return err
	}

	return nil

	//g := make([][]interface{}, 0)

	//g = append(g, []interface{}{
	//"DaysHeld",
	//"Gain(%)",
	//map[string]string{
	//"role": "style",
	//},
	//})

	//for _, o := range c.winners {
	//g = append(g, []interface{}{
	//o.Sell.DaysHeld,
	//o.Sell.GainP,
	//fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
	//})
	//}

	//for _, o := range c.losers {
	//g = append(g, []interface{}{
	//o.Sell.DaysHeld,
	//o.Sell.GainP,
	//fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
	//})
	//}

	//jg, err := datautils.JsonB64Encrypt(g)
	//if err != nil {
	//return err
	//}

	//c.GainVsDaysHeld = jg

	//return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getBuyPoints() error {

	p, err := newPlot(
		"Buy Points",
		"",
		"",
		true,
		func(p *plot.Plot) {

			p.X.Padding = vg.Points(config.ChartXLabelPadding)

			p.X.Tick.Label.Rotation = config.ChartLabelRotation
			p.X.Tick.Label.XAlign = draw.XLeft
			p.X.Tick.Label.YAlign = draw.YCenter

			p.Y.Tick.Label.XAlign = draw.XRight
		},
	)
	if err != nil {
		return err
	}

	//p, err := plot.New()
	//if err != nil {
	//return err
	//}

	//p.Title.Text = "Buy Points"
	//p.Y.Label.Text = "Trade(s)"

	//p.Title.Font.Size = vg.Points(config.ChartFontSizeL)
	//p.X.Label.Font.Size = vg.Points(config.ChartFontSizeM)
	//p.Y.Label.Font.Size = vg.Points(config.ChartFontSizeM)

	//p.X.Tick.Label.Font.Size = vg.Points(config.ChartFontSizeS)
	//p.Y.Tick.Label.Font.Size = vg.Points(config.ChartFontSizeS)
	//p.X.Tick.Label.Rotation = config.ChartLabelRotation

	//p.X.Tick.Label.XAlign = draw.XLeft
	//p.Y.Tick.Label.XAlign = draw.XLeft
	//p.X.Tick.Label.YAlign = draw.YCenter

	//p.Y.Tick.Label.XAlign = draw.XRight

	////p.X.Padding = vg.Points(config.ChartXLabelPadding)

	//err = p.Title.Font.SetName(config.ChartFont)
	//if err != nil {
	//return err
	//}

	//err = p.X.Label.Font.SetName(config.ChartFont)
	//if err != nil {
	//return err
	//}

	//err = p.Y.Label.Font.SetName(config.ChartFont)
	//if err != nil {
	//return err
	//}

	//err = p.X.Tick.Label.Font.SetName(config.ChartFont)
	//if err != nil {
	//return err
	//}

	//err = p.Y.Tick.Label.Font.SetName(config.ChartFont)
	//if err != nil {
	//return err
	//}

	//p.Add(plotter.NewGrid())

	//winners := make(map[string]int)
	//losers := make(map[string]int)

	//for _, o := range c.winners {
	//buyPoint := strings.TrimSpace(o.Buy.BuyPoint)

	//if val, ok := winners[buyPoint]; ok {
	//winners[buyPoint] = val + 1
	//} else {
	//winners[buyPoint] = 1
	//}
	//}

	//for _, o := range c.losers {
	//buyPoint := strings.TrimSpace(o.Buy.BuyPoint)

	//if val, ok := losers[buyPoint]; ok {
	//losers[buyPoint] = val + 1
	//} else {
	//losers[buyPoint] = 1
	//}
	//}

	//ck := make([]string, 0)

	//for k, _ := range winners {
	//ck = append(ck, k)
	//}

	//outer:
	//for k, _ := range losers {

	//for _, kk := range ck {
	//if kk == k {
	//continue outer
	//}
	//}

	//ck = append(ck, k)
	//}

	//sort.Strings(ck)

	//winnersG := make([]float64, len(ck))
	//losersG := make([]float64, len(ck))

	//wmax := 0.0
	//lmax := 0.0

	//for i, c := range ck {
	//if vw, ok := winners[c]; ok {
	//winnersG[i] = float64(vw)
	//if float64(vw) > wmax {
	//wmax = float64(vw)
	//}
	//} else {
	//winnersG[i] = 0.0
	//}
	//if vl, ok := losers[c]; ok {
	//losersG[i] = float64(vl)
	//if float64(vl) > lmax {
	//lmax = float64(vl)
	//}
	//} else {
	//losersG[i] = 0.0
	//}
	//}

	keys, winnersG, losersG, wmax, lmax :=
		makeWinLoseSlice(
			c.winners,
			c.losers,
			func(o *transaction.Transaction) interface{} {
				return strings.TrimSpace(o.Buy.BuyPoint)
			},
			func(keys []interface{}) {
				sort.Slice(keys, func(i, j int) bool {
					is := keys[i].(string)
					js := keys[j].(string)

					return is < js
				})
			},
			func(key interface{}) string {
				return key.(string)
			},
		)

	//func makeWinLoseFloatSlice(
	//winners []*transaction.Transaction,
	//losers []*transaction.Transaction,
	//labelCb func(o *transaction.Transaction) interface{},
	//keysSortCb func(keys []interface{}),
	//keyFormatCb func(key interface{}) string,
	//) ([]string, []float64, []float64, float64, float64) {

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

	p.Y.Max = math.Max(wmax, lmax) * config.ChartLegendPaddingYRatio

	p.Legend.Add("winners", wb)
	p.Legend.Add("losers", lb)
	//p.Legend.Font.Size = vg.Points(config.ChartFontSizeS)
	//p.Legend.YAlign = draw.YBottom
	//p.Legend.TextStyle.YAlign = draw.YBottom
	//p.Legend.Top = true
	//p.Legend.Left = true

	p.NominalX(keys...)

	//writer, err := p.WriterTo(vg.Points(config.ChartWidth), vg.Points(config.ChartHeight), "png")
	//if err != nil {
	//return err
	//}

	//buffer := new(bytes.Buffer)

	//_, err = writer.WriteTo(buffer)
	//if err != nil {
	//return err
	//}

	//encode := base64.StdEncoding.EncodeToString(buffer.Bytes())

	//c.BuyPoints = template.URL(fmt.Sprintf(config.ChartDataUrlFormat, encode))
	c.BuyPoints, err = plotToDataUrl(p)
	if err != nil {
		return err
	}

	return nil

	//g := make([][]interface{}, 0)

	//g = append(g, []interface{}{
	//"BuyPoint",
	//"Winner",
	//map[string]string{
	//"role": "style",
	//},
	//"Loser",
	//map[string]string{
	//"role": "style",
	//},
	//})

	//dictBuyPointW := make(map[string]int)
	//dictBuyPointL := make(map[string]int)

	//for _, o := range c.winners {
	//buyPoint := strings.TrimSpace(o.Buy.BuyPoint)

	//if val, ok := dictBuyPointW[buyPoint]; ok {
	//dictBuyPointW[buyPoint] = val + 1
	//} else {
	//dictBuyPointW[buyPoint] = 1
	//}
	//}

	//for _, o := range c.losers {
	//buyPoint := strings.TrimSpace(o.Buy.BuyPoint)

	//if val, ok := dictBuyPointL[buyPoint]; ok {
	//dictBuyPointL[buyPoint] = val + 1
	//} else {
	//dictBuyPointL[buyPoint] = 1
	//}
	//}

	//ck := make([]string, 0)

	//for k, _ := range dictBuyPointW {
	//ck = append(ck, k)
	//}

	//outer:
	//for k, _ := range dictBuyPointL {
	//for _, kk := range ck {
	//if kk == k {
	//continue outer
	//}
	//}

	//ck = append(ck, k)
	//}

	//sort.Strings(ck)

	//for _, c := range ck {

	//var vw int
	//var vl int

	//if v, ok := dictBuyPointW[c]; ok {
	//vw = v
	//} else {
	//vw = 0
	//}

	//if v, ok := dictBuyPointL[c]; ok {
	//vl = v
	//} else {
	//vl = 0
	//}

	//g = append(g, []interface{}{
	//c,
	//vw,
	//fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
	//vl,
	//fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
	//})
	//}

	//jg, err := datautils.JsonB64Encrypt(g)
	//if err != nil {
	//return err
	//}

	//c.BuyPoints = jg

	//return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getPriceInterval() error {

	p, err := newPlot(
		"Price Interval",
		"",
		"",
		true,
		func(p *plot.Plot) {

			p.X.Padding = vg.Points(config.ChartXLabelPadding)

			p.X.Tick.Label.Rotation = config.ChartLabelRotation
			p.X.Tick.Label.XAlign = draw.XLeft
			p.X.Tick.Label.YAlign = draw.YCenter

			p.Y.Tick.Label.XAlign = draw.XRight
		},
	)
	if err != nil {
		return err
	}

	keys, winnersG, losersG, wmax, lmax :=
		makeWinLoseSlice(
			c.winners,
			c.losers,
			func(o *transaction.Transaction) interface{} {
				grp := math.Floor(o.Buy.Price / config.PriceInterval)
				grps := int(grp * config.PriceInterval)

				return grps
			},
			func(keys []interface{}) {
				sort.Slice(keys, func(i, j int) bool {
					is := keys[i].(int)
					js := keys[j].(int)

					return is < js
				})
			},
			func(key interface{}) string {
				p := message.NewPrinter(message.MatchLanguage("en"))

				grp := math.Floor(float64(key.(int)) / config.PriceInterval)
				grpk := p.Sprintf(config.PriceIntervalFormat, int(grp*config.PriceInterval), int((grp+1)*config.PriceInterval))

				return grpk
			},
		)

	//func makeWinLoseFloatSlice(
	//winners []*transaction.Transaction,
	//losers []*transaction.Transaction,
	//labelCb func(o *transaction.Transaction) interface{},
	//keysSortCb func(keys []interface{}),
	//keyFormatCb func(key interface{}) string,
	//) ([]string, []float64, []float64, float64, float64) {

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

	p.Y.Max = math.Max(wmax, lmax) * config.ChartLegendPaddingYRatio

	p.Legend.Add("winners", wb)
	p.Legend.Add("losers", lb)

	p.NominalX(keys...)

	c.PriceInterval, err = plotToDataUrl(p)
	if err != nil {
		return err
	}

	return nil

	//g := make([][]interface{}, 0)

	//g = append(g, []interface{}{
	//"Price Interval",
	//"Winner",
	//map[string]string{
	//"role": "style",
	//},
	//"Loser",
	//map[string]string{
	//"role": "style",
	//},
	//})

	//dictPriceW := make(map[int]int)
	//dictPriceL := make(map[int]int)

	//for _, o := range c.winners {

	//grp := math.Floor(o.Buy.Price / config.PriceInterval)
	//grps := int(grp * config.PriceInterval)

	//if val, ok := dictPriceW[grps]; ok {
	//dictPriceW[grps] = val + 1
	//} else {
	//dictPriceW[grps] = 1
	//}
	//}

	//for _, o := range c.losers {

	//grp := math.Floor(o.Buy.Price / config.PriceInterval)
	//grps := int(grp * config.PriceInterval)

	//if val, ok := dictPriceL[grps]; ok {
	//dictPriceL[grps] = val + 1
	//} else {
	//dictPriceL[grps] = 1
	//}
	//}

	//ck := make([]int, 0)

	//for k, _ := range dictPriceW {
	//ck = append(ck, k)
	//}

	//outer:
	//for k, _ := range dictPriceL {
	//for _, c := range ck {
	//if c == k {
	//continue outer
	//}
	//}

	//ck = append(ck, k)
	//}

	//sort.Ints(ck)

	//p := message.NewPrinter(message.MatchLanguage("en"))

	//for _, k := range ck {

	//var vw int
	//var vl int

	//grp := math.Floor(float64(k) / config.PriceInterval)
	////grpk := fmt.Sprintf(config.PriceIntervalFormat, int(grp*config.PriceInterval), int((grp+1)*config.PriceInterval))
	//grpk := p.Sprintf(config.PriceIntervalFormat, int(grp*config.PriceInterval), int((grp+1)*config.PriceInterval))

	//if v, ok := dictPriceW[k]; ok {
	//vw = v
	//} else {
	//vw = 0
	//}

	//if v, ok := dictPriceL[k]; ok {
	//vl = v
	//} else {
	//vl = 0
	//}

	//g = append(g, []interface{}{
	//grpk,
	//vw,
	//fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
	//vl,
	//fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
	//})
	//}

	//jg, err := datautils.JsonB64Encrypt(g)
	//if err != nil {
	//return err
	//}

	//c.PriceInterval = jg

	//return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getStage() error {

	p, err := newPlot(
		"Stage",
		"",
		"",
		true,
		func(p *plot.Plot) {

			p.X.Padding = vg.Points(config.ChartXLabelPadding)
			//p.X.Tick.Label.Rotation = config.ChartLabelRotation

			p.X.Tick.Label.XAlign = draw.XLeft
			p.X.Tick.Label.YAlign = draw.YCenter

			p.Y.Tick.Label.XAlign = draw.XRight
		},
	)
	if err != nil {
		return err
	}

	keys, winnersG, losersG, wmax, lmax :=
		makeWinLoseSlice(
			c.winners,
			c.losers,
			func(o *transaction.Transaction) interface{} {
				stage := strconv.FormatFloat(math.Floor(o.Buy.Stage), 'f', -1, 64)
				return stage
			},
			func(keys []interface{}) {
				sort.Slice(keys, func(i, j int) bool {
					is := keys[i].(string)
					js := keys[j].(string)

					return is < js
				})
			},
			func(key interface{}) string {
				return key.(string)
			},
		)

	//func makeWinLoseFloatSlice(
	//winners []*transaction.Transaction,
	//losers []*transaction.Transaction,
	//labelCb func(o *transaction.Transaction) interface{},
	//keysSortCb func(keys []interface{}),
	//keyFormatCb func(key interface{}) string,
	//) ([]string, []float64, []float64, float64, float64) {

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

	p.Y.Max = math.Max(wmax, lmax) * config.ChartLegendPaddingYRatio

	//p.X.Min = 0.0 - (float64(len(keys)) * 0.05)

	p.Legend.Add("winners", wb)
	p.Legend.Add("losers", lb)

	p.NominalX(keys...)

	c.Stage, err = plotToDataUrl(p)
	if err != nil {
		return err
	}

	return nil

	//g := make([][]interface{}, 0)

	//g = append(g, []interface{}{
	//"Stage",
	//"Winner",
	//map[string]string{
	//"role": "style",
	//},
	//"Loser",
	//map[string]string{
	//"role": "style",
	//},
	//})

	//dictStageW := make(map[string]int)
	//dictStageL := make(map[string]int)

	//for _, o := range c.winners {
	//stages := strconv.FormatFloat(math.Floor(o.Buy.Stage), 'f', -1, 64)

	//if val, ok := dictStageW[stages]; ok {
	//dictStageW[stages] = val + 1
	//} else {
	//dictStageW[stages] = 1
	//}
	//}

	//for _, o := range c.losers {
	//stages := strconv.FormatFloat(math.Floor(o.Buy.Stage), 'f', -1, 64)

	//if val, ok := dictStageL[stages]; ok {
	//dictStageL[stages] = val + 1
	//} else {
	//dictStageL[stages] = 1
	//}
	//}

	//ck := make([]string, 0)

	//for k, _ := range dictStageW {
	//ck = append(ck, k)
	//}

	//outer:
	//for k, _ := range dictStageL {
	//for _, kk := range ck {
	//if kk == k {
	//continue outer
	//}
	//}

	//ck = append(ck, k)
	//}

	//sort.Strings(ck)

	//for _, c := range ck {

	//var vw int
	//var vl int

	//if v, ok := dictStageW[c]; ok {
	//vw = v
	//} else {
	//vw = 0
	//}

	//if v, ok := dictStageL[c]; ok {
	//vl = v
	//} else {
	//vl = 0
	//}

	//g = append(g, []interface{}{
	//c,
	//vw,
	//fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
	//vl,
	//fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
	//})
	//}

	//jg, err := datautils.JsonB64Encrypt(g)
	//if err != nil {
	//return err
	//}

	//c.Stage = jg

	//return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func newPlot(
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

func makeWinLoseSlice(
	winners []*transaction.Transaction,
	losers []*transaction.Transaction,
	labelCb func(o *transaction.Transaction) interface{},
	keysSortCb func(keys []interface{}),
	keyFormatCb func(key interface{}) string,
) ([]string, []float64, []float64, float64, float64) {

	dictW := make(map[interface{}]int)
	dictL := make(map[interface{}]int)

	for _, o := range winners {

		grps := labelCb(o)

		if val, ok := dictW[grps]; ok {
			dictW[grps] = val + 1
		} else {
			dictW[grps] = 1
		}
	}

	for _, o := range losers {

		grps := labelCb(o)

		if val, ok := dictL[grps]; ok {
			dictL[grps] = val + 1
		} else {
			dictL[grps] = 1
		}
	}

	fmt.Println(dictW)
	fmt.Println(dictL)

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
	}

	keys := make([]string, len(ck))

	for i, c := range ck {
		keys[i] = keyFormatCb(c)
	}

	return keys, winnersG, losersG, wmax, lmax
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
