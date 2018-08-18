package statistic

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"math"
	"sort"
	"strconv"
	"strings"
	"transaction"

	"gonum.org/v1/gonum/stat"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Trades struct {
	TotalTrade int

	Price map[int]int

	BuyPoint map[string]int

	GainPMean float64
	GainPMax  float64
	GainPMin  float64

	GainDMean float64
	GainDMax  float64
	GainDMin  float64

	DaysHeldMean float64
	DaysHeldMax  float64
	DaysHeldMin  float64

	Stage map[string]int
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func newTrades(orders []*transaction.Trade) (*Trades, error) {

	dictPrice := make(map[int]int)

	dictBuyPoint := make(map[string]int)
	dictStage := make(map[string]int)

	sliceGainP := make([]float64, 0)
	sliceGainD := make([]float64, 0)
	sliceDaysHeld := make([]float64, 0)

	for _, o := range orders {

		grp := math.Floor(o.Open.Price / config.PriceInterval)
		grps := int(grp * config.PriceInterval)

		if val, ok := dictPrice[grps]; ok {
			dictPrice[grps] = val + 1
		} else {
			dictPrice[grps] = 1
		}

		buyPoint := strings.TrimSpace(o.Open.BuyPoint)

		if val, ok := dictBuyPoint[buyPoint]; ok {
			dictBuyPoint[buyPoint] = val + 1
		} else {
			dictBuyPoint[buyPoint] = 1
		}

		sliceGainP = append(sliceGainP, o.Close.GainP)
		sliceGainD = append(sliceGainD, o.Close.GainD)
		sliceDaysHeld = append(sliceDaysHeld, float64(o.Close.DaysHeld))

		stages := strconv.FormatFloat(math.Floor(o.Open.Stage), 'f', -1, 64)

		if val, ok := dictStage[stages]; ok {
			dictStage[stages] = val + 1
		} else {
			dictStage[stages] = 1
		}

	}

	t := new(Trades)
	t.TotalTrade = len(orders)
	t.Price = dictPrice
	t.BuyPoint = dictBuyPoint
	t.Stage = dictStage

	sort.Float64s(sliceGainP)
	t.GainPMax = sliceGainP[len(sliceGainP)-1]
	t.GainPMin = sliceGainP[0]
	t.GainPMean = stat.Mean(sliceGainP, nil)

	sort.Float64s(sliceGainD)
	t.GainDMax = sliceGainD[len(sliceGainD)-1]
	t.GainDMin = sliceGainD[0]
	t.GainDMean = stat.Mean(sliceGainD, nil)

	sort.Float64s(sliceDaysHeld)
	t.DaysHeldMax = sliceDaysHeld[len(sliceDaysHeld)-1]
	t.DaysHeldMin = sliceDaysHeld[0]
	t.DaysHeldMean = stat.Mean(sliceDaysHeld, nil)

	return t, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
