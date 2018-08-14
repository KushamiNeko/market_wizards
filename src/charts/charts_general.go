package charts

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"config"
	"datautils"
	"encoding/base64"
	"fmt"
	"html/template"
	"math"
	"sort"
	"strconv"
	"strings"
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
	BuyPoints     template.URL
	PriceInterval string
	Stage         string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ChartGeneralNew(filterOrders, winners, losers []*transaction.Transaction) (*ChartGeneral, error) {

	c := new(ChartGeneral)

	c.filterOrders = filterOrders

	c.winners = winners
	c.losers = losers

	var err error

	err = c.getGainVsDaysHeld()
	if err != nil {
		return nil, err
	}

	err = c.getBuyPoints()
	if err != nil {
		return nil, err
	}

	err = c.getPriceInterval()
	if err != nil {
		return nil, err
	}

	err = c.getStage()
	if err != nil {
		return nil, err
	}

	return c, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getGainVsDaysHeld() error {

	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = "Days Held vs Gain(%)"
	p.X.Label.Text = "Days Held"
	p.Y.Label.Text = "Gain(%)"

	p.Title.Font.Size = vg.Points(config.ChartFontSizeL)
	p.X.Label.Font.Size = vg.Points(config.ChartFontSizeM)
	p.Y.Label.Font.Size = vg.Points(config.ChartFontSizeM)

	p.X.Tick.Label.Font.Size = vg.Points(config.ChartFontSizeS)
	p.Y.Tick.Label.Font.Size = vg.Points(config.ChartFontSizeS)

	err = p.Title.Font.SetName(config.ChartFont)
	if err != nil {
		return err
	}

	err = p.X.Label.Font.SetName(config.ChartFont)
	if err != nil {
		return err
	}

	err = p.Y.Label.Font.SetName(config.ChartFont)
	if err != nil {
		return err
	}

	err = p.X.Tick.Label.Font.SetName(config.ChartFont)
	if err != nil {
		return err
	}

	err = p.Y.Tick.Label.Font.SetName(config.ChartFont)
	if err != nil {
		return err
	}

	p.Add(plotter.NewGrid())

	pts := make(plotter.XYs, 0)
	for _, o := range c.winners {
		pts = append(pts, struct{ X, Y float64 }{
			float64(o.Sell.DaysHeld),
			float64(o.Sell.GainP),
		})
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

	//p.Legend.Add("winners", ws)
	//p.Legend.Add("losers", ls)
	//p.Legend.Font.Size = vg.Points(config.ChartPointRadius * 3)

	writer, err := p.WriterTo(vg.Points(config.ChartWidth), vg.Points(config.ChartHeight), "png")
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)

	_, err = writer.WriteTo(buffer)
	if err != nil {
		return err
	}

	encode := base64.StdEncoding.EncodeToString(buffer.Bytes())

	c.GainVsDaysHeld = template.URL(fmt.Sprintf("data:image/png;base64,%s", encode))

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

	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = "Buy Points"
	//p.X.Label.Text = "Days Held"
	//p.Y.Label.Text = "Gain(%)"

	p.Title.Font.Size = vg.Points(config.ChartFontSizeL)
	p.X.Label.Font.Size = vg.Points(config.ChartFontSizeM)
	p.Y.Label.Font.Size = vg.Points(config.ChartFontSizeM)

	p.X.Tick.Label.Font.Size = vg.Points(config.ChartFontSizeS)
	p.Y.Tick.Label.Font.Size = vg.Points(config.ChartFontSizeS)
	p.X.Tick.Label.Rotation = config.ChartLabelRotation

	p.X.Tick.Label.XAlign = draw.XLeft
	p.Y.Tick.Label.XAlign = draw.XLeft
	p.X.Tick.Label.YAlign = draw.YCenter
	p.Y.Tick.Label.YAlign = draw.YCenter

	p.X.Padding = vg.Points(config.ChartXLabelPadding)

	err = p.Title.Font.SetName(config.ChartFont)
	if err != nil {
		return err
	}

	err = p.X.Label.Font.SetName(config.ChartFont)
	if err != nil {
		return err
	}

	err = p.Y.Label.Font.SetName(config.ChartFont)
	if err != nil {
		return err
	}

	err = p.X.Tick.Label.Font.SetName(config.ChartFont)
	if err != nil {
		return err
	}

	err = p.Y.Tick.Label.Font.SetName(config.ChartFont)
	if err != nil {
		return err
	}

	p.Add(plotter.NewGrid())

	width := vg.Points(config.ChartBarWidth)

	winners := make(map[string]int)
	losers := make(map[string]int)

	for _, o := range c.winners {
		buyPoint := strings.TrimSpace(o.Buy.BuyPoint)

		if val, ok := winners[buyPoint]; ok {
			winners[buyPoint] = val + 1
		} else {
			winners[buyPoint] = 1
		}
	}

	for _, o := range c.losers {
		buyPoint := strings.TrimSpace(o.Buy.BuyPoint)

		if val, ok := losers[buyPoint]; ok {
			losers[buyPoint] = val + 1
		} else {
			losers[buyPoint] = 1
		}
	}

	ck := make([]string, 0)

	for k, _ := range winners {
		ck = append(ck, k)
	}

outer:
	for k, _ := range losers {

		for _, kk := range ck {
			if kk == k {
				continue outer
			}
		}

		ck = append(ck, k)
	}

	sort.Strings(ck)

	winnersG := make([]float64, len(ck))
	losersG := make([]float64, len(ck))

	for i, c := range ck {
		if vw, ok := winners[c]; ok {
			winnersG[i] = float64(vw)
		} else {
			winnersG[i] = 0.0
		}
		if vl, ok := losers[c]; ok {
			losersG[i] = float64(vl)
		} else {
			losersG[i] = 0.0
		}
	}

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

	//p.Legend.Add("winners", wb)
	//p.Legend.Add("losers", lb)
	//p.Legend.Top = true

	p.NominalX(ck...)

	writer, err := p.WriterTo(vg.Points(config.ChartWidth), vg.Points(config.ChartHeight), "png")
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)

	_, err = writer.WriteTo(buffer)
	if err != nil {
		return err
	}

	encode := base64.StdEncoding.EncodeToString(buffer.Bytes())

	c.BuyPoints = template.URL(fmt.Sprintf("data:image/png;base64,%s", encode))

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
	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		"Price Interval",
		"Winner",
		map[string]string{
			"role": "style",
		},
		"Loser",
		map[string]string{
			"role": "style",
		},
	})

	dictPriceW := make(map[int]int)
	dictPriceL := make(map[int]int)

	for _, o := range c.winners {

		grp := math.Floor(o.Buy.Price / config.PriceInterval)
		grps := int(grp * config.PriceInterval)

		if val, ok := dictPriceW[grps]; ok {
			dictPriceW[grps] = val + 1
		} else {
			dictPriceW[grps] = 1
		}
	}

	for _, o := range c.losers {

		grp := math.Floor(o.Buy.Price / config.PriceInterval)
		grps := int(grp * config.PriceInterval)

		if val, ok := dictPriceL[grps]; ok {
			dictPriceL[grps] = val + 1
		} else {
			dictPriceL[grps] = 1
		}
	}

	ck := make([]int, 0)

	for k, _ := range dictPriceW {
		ck = append(ck, k)
	}

outer:
	for k, _ := range dictPriceL {
		for _, c := range ck {
			if c == k {
				continue outer
			}
		}

		ck = append(ck, k)
	}

	sort.Ints(ck)

	p := message.NewPrinter(message.MatchLanguage("en"))

	for _, k := range ck {

		var vw int
		var vl int

		grp := math.Floor(float64(k) / config.PriceInterval)
		//grpk := fmt.Sprintf(config.PriceIntervalFormat, int(grp*config.PriceInterval), int((grp+1)*config.PriceInterval))
		grpk := p.Sprintf(config.PriceIntervalFormat, int(grp*config.PriceInterval), int((grp+1)*config.PriceInterval))

		if v, ok := dictPriceW[k]; ok {
			vw = v
		} else {
			vw = 0
		}

		if v, ok := dictPriceL[k]; ok {
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
		return err
	}

	c.PriceInterval = jg

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getStage() error {

	g := make([][]interface{}, 0)

	g = append(g, []interface{}{
		"Stage",
		"Winner",
		map[string]string{
			"role": "style",
		},
		"Loser",
		map[string]string{
			"role": "style",
		},
	})

	dictStageW := make(map[string]int)
	dictStageL := make(map[string]int)

	for _, o := range c.winners {
		stages := strconv.FormatFloat(math.Floor(o.Buy.Stage), 'f', -1, 64)

		if val, ok := dictStageW[stages]; ok {
			dictStageW[stages] = val + 1
		} else {
			dictStageW[stages] = 1
		}
	}

	for _, o := range c.losers {
		stages := strconv.FormatFloat(math.Floor(o.Buy.Stage), 'f', -1, 64)

		if val, ok := dictStageL[stages]; ok {
			dictStageL[stages] = val + 1
		} else {
			dictStageL[stages] = 1
		}
	}

	ck := make([]string, 0)

	for k, _ := range dictStageW {
		ck = append(ck, k)
	}

outer:
	for k, _ := range dictStageL {
		for _, kk := range ck {
			if kk == k {
				continue outer
			}
		}

		ck = append(ck, k)
	}

	sort.Strings(ck)

	for _, c := range ck {

		var vw int
		var vl int

		if v, ok := dictStageW[c]; ok {
			vw = v
		} else {
			vw = 0
		}

		if v, ok := dictStageL[c]; ok {
			vl = v
		} else {
			vl = 0
		}

		g = append(g, []interface{}{
			c,
			vw,
			fmt.Sprintf(config.StyleFormat, config.WinnerColor, config.WinnerOpacity),
			vl,
			fmt.Sprintf(config.StyleFormat, config.LoserColor, config.LoserOpacity),
		})
	}

	jg, err := datautils.JsonB64Encrypt(g)
	if err != nil {
		return err
	}

	c.Stage = jg

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
