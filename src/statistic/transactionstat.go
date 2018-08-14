package statistic

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"math"
	"sort"
	"strconv"
	"strings"
	"transaction"

	//"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type TransactionStat struct {
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

func NewTransactionStat(orders []*transaction.Transaction) (*TransactionStat, error) {

	dictPrice := make(map[int]int)

	dictBuyPoint := make(map[string]int)
	dictStage := make(map[string]int)

	sliceGainP := make([]float64, 0)
	sliceGainD := make([]float64, 0)
	sliceDaysHeld := make([]float64, 0)

	for _, o := range orders {

		grp := math.Floor(o.Buy.Price / config.PriceInterval)
		grps := int(grp * config.PriceInterval)

		if val, ok := dictPrice[grps]; ok {
			dictPrice[grps] = val + 1
		} else {
			dictPrice[grps] = 1
		}

		buyPoint := strings.TrimSpace(o.Buy.BuyPoint)

		if val, ok := dictBuyPoint[buyPoint]; ok {
			dictBuyPoint[buyPoint] = val + 1
		} else {
			dictBuyPoint[buyPoint] = 1
		}

		sliceGainP = append(sliceGainP, o.Sell.GainP)
		sliceGainD = append(sliceGainD, o.Sell.GainD)
		sliceDaysHeld = append(sliceDaysHeld, float64(o.Sell.DaysHeld))

		stages := strconv.FormatFloat(math.Floor(o.Buy.Stage), 'f', -1, 64)

		if val, ok := dictStage[stages]; ok {
			dictStage[stages] = val + 1
		} else {
			dictStage[stages] = 1
		}

	}

	t := new(TransactionStat)
	t.TotalTrade = len(orders)
	t.Price = dictPrice
	t.BuyPoint = dictBuyPoint
	t.Stage = dictStage

	//var err error

	t.GainPMean = stat.Mean(sliceGainP, nil)
	//t.GainPMean, err = stats.Mean(sliceGainP)
	//if err != nil {
	//return nil, err
	//}

	sort.Float64s(sliceGainP)

	t.GainPMax = sliceGainP[len(sliceGainP)-1]
	//t.GainPMax, err = stats.Max(sliceGainP)
	//if err != nil {
	//return nil, err
	//}

	t.GainPMin = sliceGainP[0]
	//t.GainPMin, err = stats.Min(sliceGainP)
	//if err != nil {
	//return nil, err
	//}

	t.GainDMean = stat.Mean(sliceGainD, nil)
	//t.GainDMean, err = stats.Mean(sliceGainD)
	//if err != nil {
	//return nil, err
	//}

	sort.Float64s(sliceGainD)

	t.GainDMax = sliceGainD[len(sliceGainD)-1]
	//t.GainDMax, err = stats.Max(sliceGainD)
	//if err != nil {
	//return nil, err
	//}

	t.GainDMin = sliceGainD[0]
	//t.GainDMin, err = stats.Min(sliceGainD)
	//if err != nil {
	//return nil, err
	//}

	t.DaysHeldMean = stat.Mean(sliceDaysHeld, nil)
	//t.DaysHeldMean, err = stats.Mean(sliceDaysHeld)
	//if err != nil {
	//return nil, err
	//}

	sort.Float64s(sliceDaysHeld)

	t.DaysHeldMax = sliceDaysHeld[len(sliceDaysHeld)-1]
	//t.DaysHeldMax, err = stats.Max(sliceDaysHeld)
	//if err != nil {
	//return nil, err
	//}

	t.DaysHeldMin = sliceDaysHeld[0]
	//t.DaysHeldMin, err = stats.Min(sliceDaysHeld)
	//if err != nil {
	//return nil, err
	//}

	return t, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
