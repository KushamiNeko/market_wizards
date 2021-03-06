package analysis

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"html/template"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"transaction"

	"golang.org/x/text/message"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ChartGeneral struct {
	filterOrders []*transaction.Trade

	winners []*transaction.Trade
	losers  []*transaction.Trade

	SortedTrades []*transaction.Trade

	winnersI []interface{}
	losersI  []interface{}

	Threshold float64

	GainVsDaysHeld template.URL
	BuyPoints      template.URL
	PriceInterval  template.URL
	Stage          template.URL

	BattingAverage  template.URL
	AverageGainLoss template.URL
	//MaxGL          template.URL
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ChartGeneralNew(filterOrders, winners, losers []*transaction.Trade, threshold float64) (*ChartGeneral, error) {

	var wg sync.WaitGroup

	c := new(ChartGeneral)

	c.filterOrders = filterOrders

	c.winners = winners
	c.losers = losers

	c.winnersI = make([]interface{}, len(c.winners))
	c.losersI = make([]interface{}, len(c.losers))

	c.Threshold = threshold

	trades := make([]*transaction.Trade, 0)

	for _, w := range c.winners {
		trades = append(trades, w)
	}

	for _, l := range c.losers {
		trades = append(trades, l)
	}

	sort.Slice(trades, func(i, j int) bool {
		return trades[i].Close.Date < trades[j].Close.Date
	})

	c.SortedTrades = trades

	for i, v := range c.winners {
		c.winnersI[i] = v
	}

	for i, v := range c.losers {
		c.losersI[i] = v
	}

	fs := 6
	ei := 0

	wg.Add(fs)
	errs := make([]error, fs)

	go func(ei int) {
		err := c.getGainVsDaysHeld()
		if err != nil {
			errs[ei] = err
		}
		wg.Done()
	}(ei)
	ei += 1

	go func(ei int) {
		err := c.getBuyPoints()
		if err != nil {
			errs[ei] = err
		}

		wg.Done()
	}(ei)
	ei += 1

	go func(ei int) {
		err := c.getPriceInterval()
		if err != nil {
			errs[ei] = err
		}

		wg.Done()
	}(ei)
	ei += 1

	go func(ei int) {
		err := c.getStage()
		if err != nil {
			errs[ei] = err
		}

		wg.Done()
	}(ei)
	ei += 1

	go func(ei int) {
		err := c.getBattingAverage()
		if err != nil {
			errs[ei] = err
		}

		wg.Done()
	}(ei)
	ei += 1

	go func(ei int) {
		err := c.getAverageGainLoss()
		if err != nil {
			errs[ei] = err
		}

		wg.Done()
	}(ei)
	ei += 1

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

	p, err := makePlot(
		"Days Held vs Gain(%)",
		"Days Held",
		"Gain(%)",
		true,
		nil,
	)
	if err != nil {
		return err
	}

	max := 0.0

	pts := make(plotter.XYs, 0)

	for _, o := range c.winners {
		pts = append(pts, struct{ X, Y float64 }{
			float64(o.Close.DaysHeld),
			float64(o.Close.GainP),
		})

		if float64(o.Close.GainP) > max {
			max = float64(o.Close.GainP)
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
			float64(o.Close.DaysHeld),
			float64(o.Close.GainP),
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

	p.X.Min = -(p.X.Max * (config.ChartBarPaddingXRatio - 1.0) / 2.0)
	p.X.Max = p.X.Max * (((config.ChartBarPaddingXRatio - 1.0) / 2.0) + 1.0)

	p.Legend.Add("winners", ws)
	p.Legend.Add("losers", ls)

	c.GainVsDaysHeld, err = plotToDataUrl(p)
	if err != nil {
		return err
	}

	return nil

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getBuyPoints() error {

	p, err := makePlot(
		"Buy Points",
		"",
		"",
		true,
		func(p *plot.Plot) {
			barChartSetup(p)
			p.X.Tick.Label.Rotation = config.ChartLabelRotation
		},
	)
	if err != nil {
		return err
	}

	keys, winnersG, losersG, wmax, lmax, err :=
		makeValueSlice(
			c.winnersI,
			c.losersI,
			func(o interface{}) (interface{}, error) {
				t := o.(*transaction.Trade)
				return strings.TrimSpace(t.Open.BuyPoint), nil
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
	if err != nil {
		return err
	}

	err = makeBarCharts(p, keys, winnersG, losersG, wmax, lmax)
	if err != nil {
		return err
	}

	c.BuyPoints, err = plotToDataUrl(p)
	if err != nil {
		return err
	}

	return nil

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getPriceInterval() error {

	p, err := makePlot(
		"Price Interval",
		"",
		"",
		true,
		func(p *plot.Plot) {
			barChartSetup(p)
			p.X.Tick.Label.Rotation = config.ChartLabelRotation
		},
	)
	if err != nil {
		return err
	}

	keys, winnersG, losersG, wmax, lmax, err :=
		makeValueSlice(
			c.winnersI,
			c.losersI,
			func(o interface{}) (interface{}, error) {
				t := o.(*transaction.Trade)
				grp := math.Floor(t.Open.Price / config.PriceInterval)
				grps := int(grp * config.PriceInterval)

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
				p := message.NewPrinter(message.MatchLanguage("en"))

				grp := math.Floor(float64(key.(int)) / config.PriceInterval)
				grpk := p.Sprintf(config.PriceIntervalFormat, int(grp*config.PriceInterval), int((grp+1)*config.PriceInterval))

				return grpk
			},
		)
	if err != nil {
		return err
	}

	err = makeBarCharts(p, keys, winnersG, losersG, wmax, lmax)
	if err != nil {
		return err
	}

	c.PriceInterval, err = plotToDataUrl(p)
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getStage() error {

	p, err := makePlot(
		"Stage",
		"",
		"",
		true,
		func(p *plot.Plot) {
			barChartSetup(p)
		},
	)
	if err != nil {
		return err
	}

	keys, winnersG, losersG, wmax, lmax, err :=
		makeValueSlice(
			c.winnersI,
			c.losersI,
			func(o interface{}) (interface{}, error) {
				t := o.(*transaction.Trade)
				stage := strconv.FormatFloat(math.Floor(t.Open.Stage), 'f', -1, 64)
				return stage, nil
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
	if err != nil {
		return err
	}

	err = makeBarCharts(p, keys, winnersG, losersG, wmax, lmax)
	if err != nil {
		return err
	}

	c.Stage, err = plotToDataUrl(p)
	if err != nil {
		return err
	}

	return nil

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getBattingAverage() error {

	max := 0.0
	pts := make(plotter.XYs, 0)

	round := int(float64(len(c.SortedTrades)) / config.StatisticBase)

	for i := 0; i < round; i++ {
		slice := c.SortedTrades[i*config.StatisticBase : (i+1)*config.StatisticBase]

		win := 0.0

		for _, s := range slice {
			if s.Close.GainP >= c.Threshold {
				win += 1.0
			}
		}

		batting := float64(win) / float64(config.StatisticBase)

		pts = append(pts, struct{ X, Y float64 }{
			float64(config.StatisticBase * (i + 1)),
			batting,
		})

		if batting > max {
			max = batting
		}
	}

	p, err := makePlot(
		"Batting Average",
		"Trades",
		"",
		true,
		nil,
	)
	if err != nil {
		return err
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
	p.Y.Max = max * config.ChartLegendPaddingYRatio
	p.Y.Min = -(p.Y.Max * (config.ChartBarPaddingXRatio - 1.0) / 2.0)

	//p.X.Min = -(p.X.Max * (config.ChartBarXPaddingRatio - 1.0) / 2.0)
	//p.X.Max = p.X.Max * (((config.ChartBarXPaddingRatio - 1.0) / 2.0) + 1.0)

	c.BattingAverage, err = plotToDataUrl(p)
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *ChartGeneral) getAverageGainLoss() error {

	max := 0.0
	min := 0.0

	ptsw := make(plotter.XYs, 0)
	ptsl := make(plotter.XYs, 0)

	round := int(float64(len(c.SortedTrades)) / config.StatisticBase)

	for i := 0; i < round; i++ {
		slice := c.SortedTrades[i*config.StatisticBase : (i+1)*config.StatisticBase]

		mws := make([]float64, 0)
		mls := make([]float64, 0)

		for _, s := range slice {
			if s.Close.GainP >= c.Threshold {
				mws = append(mws, s.Close.GainP)

				if s.Close.GainP > max {
					max = s.Close.GainP
				}

			} else {
				mls = append(mls, s.Close.GainP)

				if s.Close.GainP < min {
					min = s.Close.GainP
				}
			}
		}

		mw := stat.Mean(mws, nil)
		ml := stat.Mean(mls, nil)

		ptsw = append(ptsw, struct{ X, Y float64 }{
			float64(config.StatisticBase * (i + 1)),
			mw,
		})

		ptsl = append(ptsl, struct{ X, Y float64 }{
			float64(config.StatisticBase * (i + 1)),
			ml,
		})
	}

	p, err := makePlot(
		"Gain/Loss Average",
		"",
		"Gain(%)",
		true,
		nil,
	)
	if err != nil {
		return err
	}

	linew, pointsw, err := plotter.NewLinePoints(ptsw)
	if err != nil {
		return err
	}

	linew.Color = config.WinnerRGBA
	linew.LineStyle.Width = vg.Points(config.ChartLineWidth)

	pointsw.Shape = draw.CircleGlyph{}
	pointsw.Radius = vg.Points(config.ChartPointRadius)
	pointsw.Color = config.WinnerRGBA

	p.Add(linew, pointsw)

	linel, pointsl, err := plotter.NewLinePoints(ptsl)
	if err != nil {
		return err
	}

	linel.Color = config.LoserRGBA
	linel.LineStyle.Width = vg.Points(config.ChartLineWidth)

	pointsl.Shape = draw.CircleGlyph{}
	pointsl.Radius = vg.Points(config.ChartPointRadius)
	pointsl.Color = config.LoserRGBA

	p.Add(linel, pointsl)

	p.Legend.Add("winners", linew, pointsw)
	p.Legend.Add("losers", linel, pointsl)

	p.Y.Max = max * config.ChartLegendPaddingYRatio
	p.Y.Min = min * config.ChartLegendPaddingYRatio

	//p.X.Min = -(p.X.Max * (config.ChartBarXPaddingRatio - 1.0) / 2.0)
	//p.X.Max = p.X.Max * (((config.ChartBarXPaddingRatio - 1.0) / 2.0) + 1.0)

	c.AverageGainLoss, err = plotToDataUrl(p)
	if err != nil {
		return err
	}

	return nil

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
