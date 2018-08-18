package charts

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

	winnersI []interface{}
	losersI  []interface{}

	GainVsDaysHeld template.URL
	BuyPoints      template.URL
	PriceInterval  template.URL
	Stage          template.URL
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ChartGeneralNew(filterOrders, winners, losers []*transaction.Trade) (*ChartGeneral, error) {

	var wg sync.WaitGroup

	c := new(ChartGeneral)

	c.filterOrders = filterOrders

	c.winners = winners
	c.losers = losers

	c.winnersI = make([]interface{}, len(c.winners))
	c.losersI = make([]interface{}, len(c.losers))

	for i, v := range c.winners {
		c.winnersI[i] = v
	}

	for i, v := range c.losers {
		c.losersI[i] = v
	}

	fs := 4

	wg.Add(fs)
	errs := make([]error, fs)

	go func() {
		err := c.getGainVsDaysHeld()
		if err != nil {
			errs[0] = err
		}
		wg.Done()
	}()

	go func() {
		err := c.getBuyPoints()
		if err != nil {
			errs[1] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getPriceInterval()
		if err != nil {
			errs[2] = err
		}

		wg.Done()
	}()

	go func() {
		err := c.getStage()
		if err != nil {
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

	p.X.Min = -(p.X.Max * (config.ChartBarXPaddingRatio - 1.0) / 2.0)
	p.X.Max = p.X.Max * (((config.ChartBarXPaddingRatio - 1.0) / 2.0) + 1.0)

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
