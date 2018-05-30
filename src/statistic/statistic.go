package statistic

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"charts"
	"config"
	"fmt"
	"math"
	"strconv"
	"strings"
	"transaction"

	"github.com/montanaflynn/stats"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//const (
//grpPrice  = 50
//grpFormat = "%s ~ %s"

//LoserGainThreshold = 1.0
//)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//type Charts struct {
//General     string
//IBD         string
//MarketSmith string
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Statistic struct {
	StartDate      string
	EndDate        string
	LossThresholdP float64

	//Charts *Charts
	ChartGeneral *charts.ChartGeneral
	ChartIBD     *charts.ChartIBD
	//ChartMarketSmith *charts.ChartGeneral

	TotalTrade int

	BattingAverage float64

	WinLossRatioP float64

	AdjustedWinLossRatioP float64

	ExpectedValueP float64

	WinLossRatioD float64

	AdjustedWinLossRatioD float64

	ExpectedValueD float64

	//Expectancy float64

	Gain *TransactionStat
	Loss *TransactionStat
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewStatistic(winner []*transaction.Order, loser []*transaction.Order) (*Statistic, error) {
	s := new(Statistic)

	var stat *TransactionStat
	var err error

	if winner != nil && len(winner) != 0 {
		stat, err = NewTransactionStat(winner)
		if err != nil {
			return nil, err
		}

		s.Gain = stat
	} else {
		s.Gain = new(TransactionStat)
	}

	if loser != nil && len(loser) != 0 {

		stat, err = NewTransactionStat(loser)
		if err != nil {
			return nil, err
		}

		s.Loss = stat
	} else {
		s.Loss = new(TransactionStat)
	}

	s.TotalTrade = s.Gain.TotalTrade + s.Loss.TotalTrade
	s.BattingAverage = float64(s.Gain.TotalTrade) / float64(s.TotalTrade)

	s.WinLossRatioP = s.Gain.GainPMean / math.Abs(s.Loss.GainPMean)

	s.AdjustedWinLossRatioP = s.Gain.GainPMean * s.BattingAverage / math.Abs(s.Loss.GainPMean) * (1.0 - s.BattingAverage)

	s.ExpectedValueP = s.Gain.GainPMean*s.BattingAverage + s.Loss.GainPMean*(1.0-s.BattingAverage)

	s.WinLossRatioD = s.Gain.GainDMean / math.Abs(s.Loss.GainDMean)

	s.AdjustedWinLossRatioD = s.Gain.GainDMean * s.BattingAverage / math.Abs(s.Loss.GainDMean) * (1.0 - s.BattingAverage)

	s.ExpectedValueD = s.Gain.GainDMean*s.BattingAverage + s.Loss.GainDMean*(1.0-s.BattingAverage)

	return s, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *Statistic) FormatFloat(data float64) string {
	return fmt.Sprintf("%.4f", data)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func FormatGrpPrice(price int) string {

	grp := math.Floor(float64(price) / config.PriceInterval)
	grps := strconv.FormatFloat(grp*config.PriceInterval, 'f', -1, 64)
	grpe := strconv.FormatFloat((grp+1)*config.PriceInterval, 'f', -1, 64)

	grpk := fmt.Sprintf(config.PriceIntervalFormat, grps, grpe)

	return grpk
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type TransactionStat struct {
	//ID   string
	//Etag string

	//Order string

	//Date int

	//Symbol string

	//Price float64

	TotalTrade int

	//Cost map[string]int

	//Price map[string]int
	Price map[int]int

	//Share int

	//BuyPoint string
	BuyPoint map[string]int

	//Revenue float64 `datastore:",omitempty" json:",omitempty"`

	//Cost float64 `datastore:",omitempty" json:",omitempty"`

	//GainD float64 `datastore:",omitempty" json:",omitempty"`

	//GainP float64 `datastore:",omitempty" json:",omitempty"`
	GainPMean float64
	GainPMax  float64
	GainPMin  float64

	GainDMean float64
	GainDMax  float64
	GainDMin  float64

	//DaysHeld int `datastore:",omitempty" json:",omitempty"`
	DaysHeldMean float64
	DaysHeldMax  float64
	DaysHeldMin  float64

	//Stage float64
	Stage map[string]int

	//Note string `datastore:",noindex"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewTransactionStat(orders []*transaction.Order) (*TransactionStat, error) {

	//grpPrice := 10.0
	//grpFormat := "%s ~ %s"

	//dictCost := make(map[string]int)

	//dictPrice := make(map[string]int)
	dictPrice := make(map[int]int)

	dictBuyPoint := make(map[string]int)
	dictStage := make(map[string]int)

	sliceGainP := make([]float64, 0)
	sliceGainD := make([]float64, 0)
	sliceDaysHeld := make([]float64, 0)

	for _, o := range orders {
		//cost := o.Cost / float64(o.Share)

		//cgrp := math.Floor(cost / config.PriceInterval)
		//cgrps := strconv.FormatFloat(cgrp*config.PriceInterval, 'f', -1, 64)
		//cgrpe := strconv.FormatFloat((cgrp+1)*config.PriceInterval, 'f', -1, 64)

		//cgrpk := fmt.Sprintf(grpFormat, cgrps, cgrpe)

		//if val, ok := dictCost[cgrpk]; ok {
		//dictCost[cgrpk] = val + 1
		//} else {
		//dictCost[cgrpk] = 1
		//}

		grp := math.Floor(o.Price / config.PriceInterval)
		grps := int(grp * config.PriceInterval)

		//grps := strconv.FormatFloat(grp*config.PriceInterval, 'f', -1, 64)
		//grpe := strconv.FormatFloat((grp+1)*config.PriceInterval, 'f', -1, 64)

		//grpk := fmt.Sprintf(grpFormat, grps, grpe)

		//if val, ok := dictPrice[grpk]; ok {
		if val, ok := dictPrice[grps]; ok {
			//dictPrice[grpk] = val + 1
			dictPrice[grps] = val + 1
		} else {
			//dictPrice[grpk] = 1
			dictPrice[grps] = 1
		}

		buyPoint := strings.TrimSpace(o.BuyPoint)

		if val, ok := dictBuyPoint[buyPoint]; ok {
			dictBuyPoint[buyPoint] = val + 1
		} else {
			dictBuyPoint[buyPoint] = 1
		}

		sliceGainP = append(sliceGainP, o.GainP)
		sliceGainD = append(sliceGainD, o.GainD)
		sliceDaysHeld = append(sliceDaysHeld, float64(o.DaysHeld))

		stages := strconv.FormatFloat(math.Floor(o.Stage), 'f', -1, 64)

		if val, ok := dictStage[stages]; ok {
			dictStage[stages] = val + 1
		} else {
			dictStage[stages] = 1
		}

	}

	t := new(TransactionStat)
	t.TotalTrade = len(orders)
	//t.Cost = dictCost
	t.Price = dictPrice
	t.BuyPoint = dictBuyPoint
	t.Stage = dictStage

	var err error

	t.GainPMean, err = stats.Mean(sliceGainP)
	if err != nil {
		return nil, err
	}

	t.GainPMax, err = stats.Max(sliceGainP)
	if err != nil {
		return nil, err
	}

	t.GainPMin, err = stats.Min(sliceGainP)
	if err != nil {
		return nil, err
	}

	t.GainDMean, err = stats.Mean(sliceGainD)
	if err != nil {
		return nil, err
	}

	t.GainDMax, err = stats.Max(sliceGainD)
	if err != nil {
		return nil, err
	}

	t.GainDMin, err = stats.Min(sliceGainD)
	if err != nil {
		return nil, err
	}

	t.DaysHeldMean, err = stats.Mean(sliceDaysHeld)
	if err != nil {
		return nil, err
	}

	t.DaysHeldMax, err = stats.Max(sliceDaysHeld)
	if err != nil {
		return nil, err
	}

	t.DaysHeldMin, err = stats.Min(sliceDaysHeld)
	if err != nil {
		return nil, err
	}

	return t, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//type IBDCheckUp struct {
//ID string

//Symbol string

//RankInGroup           int
//CompositeRating       int
//MarketUptrend         string
//IndustryGroupRank     int
//GroupRSRating         string
//EPSRating             int
//EPSChgLastQtr         float64
//Last3QtrsAvgEPSGrowth float64
//NQtrsOfEPSAccel       int

//EPSEstChgCurrentQtr    float64
//EstimateRevisions      string
//LastQtrEarningsSuprise float64

//ThrYrEpsGrowthRate    float64
//NYrsOfAnnualEPSGrowth int
//EPSEstChgCurrentYr    float64

//SMRRating            string
//SalesChgLastQtr      float64
//ThrYrSalesGrowthRate float64
//AnnualPreTaxMargin   float64
//AnnualROE            float64
//DebtEquityRatio      float64

////Price          float64

//RSRating       int
//Off52WeekHigh  float64
//PriceVS50DayMA float64
//AvgVolume50Day int64

//MarketCapital int64
//AccDisRating  string
//UpDownVolume  float64
//ChgInFunds    float64
//QtrsOfIncFund int
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
