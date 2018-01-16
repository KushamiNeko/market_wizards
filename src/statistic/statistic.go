package statistic

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"math"
	"strconv"
	"transaction"

	"github.com/montanaflynn/stats"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Statistic struct {
	TotalTrade int

	BattingAverage float64

	WinLossRatio float64

	Gain *TransactionStat
	Loss *TransactionStat
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewStatistic(winner []*transaction.Order, losser []*transaction.Order) (*Statistic, error) {
	s := new(Statistic)

	stat, err := NewTransactionStat(winner)
	if err != nil {
		return nil, err
	}

	s.Gain = stat

	stat, err = NewTransactionStat(losser)
	if err != nil {
		return nil, err
	}

	s.Loss = stat

	s.TotalTrade = s.Gain.TotalTrade + s.Loss.TotalTrade
	s.BattingAverage = float64(s.Gain.TotalTrade) / float64(s.TotalTrade)

	s.WinLossRatio = s.Gain.GainPMean / math.Abs(s.Loss.GainPMean)

	return s, nil
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

	Cost map[string]int

	Price map[string]int

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

	//DayHold int `datastore:",omitempty" json:",omitempty"`
	DayHoldMean float64
	DayHoldMax  float64
	DayHoldMin  float64

	//Stage float64
	Stage map[string]int

	//Note string `datastore:",noindex"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewTransactionStat(orders []*transaction.Order) (*TransactionStat, error) {

	grpPrice := 10.0

	dictCost := make(map[string]int)
	dictPrice := make(map[string]int)
	dictBuyPoint := make(map[string]int)
	dictStage := make(map[string]int)

	sliceGainP := make([]float64, 0)
	sliceDayHold := make([]float64, 0)

	for _, o := range orders {
		cost := o.Cost / float64(o.Share)

		cgrp := math.Floor(cost / grpPrice)
		cgrps := strconv.FormatFloat(cgrp*grpPrice, 'f', -1, 64)
		cgrpe := strconv.FormatFloat((cgrp+1)*grpPrice, 'f', -1, 64)

		cgrpk := fmt.Sprintf("%s~%s", cgrps, cgrpe)

		if val, ok := dictCost[cgrpk]; ok {
			dictCost[cgrpk] = val + 1
		} else {
			dictCost[cgrpk] = 1
		}

		grp := math.Floor(o.Price / grpPrice)
		grps := strconv.FormatFloat(grp*grpPrice, 'f', -1, 64)
		grpe := strconv.FormatFloat((grp+1)*grpPrice, 'f', -1, 64)

		grpk := fmt.Sprintf("%s~%s", grps, grpe)

		if val, ok := dictPrice[grpk]; ok {
			dictPrice[grpk] = val + 1
		} else {
			dictPrice[grpk] = 1
		}

		if val, ok := dictBuyPoint[o.BuyPoint]; ok {
			dictBuyPoint[o.BuyPoint] = val + 1
		} else {
			dictBuyPoint[o.BuyPoint] = 1
		}

		sliceGainP = append(sliceGainP, o.GainP)
		sliceDayHold = append(sliceDayHold, float64(o.DayHold))

		stages := strconv.FormatFloat(math.Floor(o.Stage), 'f', -1, 64)

		if val, ok := dictStage[stages]; ok {
			dictStage[stages] = val + 1
		} else {
			dictStage[stages] = 1
		}

	}

	t := new(TransactionStat)
	t.TotalTrade = len(orders)
	t.Cost = dictCost
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

	t.DayHoldMean, err = stats.Mean(sliceDayHold)
	if err != nil {
		return nil, err
	}

	t.DayHoldMax, err = stats.Max(sliceDayHold)
	if err != nil {
		return nil, err
	}

	t.DayHoldMin, err = stats.Min(sliceDayHold)
	if err != nil {
		return nil, err
	}

	return t, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type IBDCheckUp struct {
	ID string

	Symbol string

	RankInGroup           int
	CompositeRating       int
	MarketUptrend         string
	IndustryGroupRank     int
	GroupRSRating         string
	EPSRating             int
	EPSChgLastQtr         float64
	Last3QtrsAvgEPSGrowth float64
	NQtrsOfEPSAccel       int

	EPSEstChgCurrentQtr    float64
	EstimateRevisions      string
	LastQtrEarningsSuprise float64

	ThrYrEpsGrowthRate    float64
	NYrsOfAnnualEPSGrowth int
	EPSEstChgCurrentYr    float64

	SMRRating            string
	SalesChgLastQtr      float64
	ThrYrSalesGrowthRate float64
	AnnualPreTaxMargin   float64
	AnnualROE            float64
	DebtEquityRatio      float64

	//Price          float64

	RSRating       int
	Off52WeekHigh  float64
	PriceVS50DayMA float64
	AvgVolume50Day int64

	MarketCapital int64
	AccDisRating  string
	UpDownVolume  float64
	ChgInFunds    float64
	QtrsOfIncFund int
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
