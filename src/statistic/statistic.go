package statistic

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"analysis"
	"fmt"
	"math"
	"transaction"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Statistic struct {
	StartDate      string
	EndDate        string
	LossThresholdP float64

	ChartGeneral     *analysis.ChartGeneral
	ChartIBD         *analysis.ChartIBD
	ChartMarketSmith *analysis.ChartMarketSmith

	TotalTrade int

	BattingAverage float64

	WinLossRatioP float64

	AdjustedWinLossRatioP float64

	ExpectedValueP float64

	KellyCriterionP float64

	WinLossRatioD float64

	AdjustedWinLossRatioD float64

	ExpectedValueD float64

	KellyCriterionD float64

	Gain *Trades
	Loss *Trades
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewStatistic(winners []*transaction.Trade, losers []*transaction.Trade) (*Statistic, error) {
	s := new(Statistic)

	var stat *Trades
	var err error

	if winners != nil && len(winners) != 0 {
		stat, err = newTrades(winners)
		if err != nil {
			return nil, err
		}

		s.Gain = stat
	} else {
		s.Gain = new(Trades)
	}

	if losers != nil && len(losers) != 0 {

		stat, err = newTrades(losers)
		if err != nil {
			return nil, err
		}

		s.Loss = stat
	} else {
		s.Loss = new(Trades)
	}

	s.TotalTrade = s.Gain.TotalTrade + s.Loss.TotalTrade
	s.BattingAverage = float64(s.Gain.TotalTrade) / float64(s.TotalTrade)

	s.WinLossRatioP = s.Gain.GainPMean / math.Abs(s.Loss.GainPMean)

	s.AdjustedWinLossRatioP = s.Gain.GainPMean * s.BattingAverage / math.Abs(s.Loss.GainPMean) * (1.0 - s.BattingAverage)

	s.ExpectedValueP = s.Gain.GainPMean*s.BattingAverage + s.Loss.GainPMean*(1.0-s.BattingAverage)

	s.KellyCriterionP = s.BattingAverage - ((1 - s.BattingAverage) / s.WinLossRatioP)

	s.WinLossRatioD = s.Gain.GainDMean / math.Abs(s.Loss.GainDMean)

	s.AdjustedWinLossRatioD = s.Gain.GainDMean * s.BattingAverage / math.Abs(s.Loss.GainDMean) * (1.0 - s.BattingAverage)

	s.ExpectedValueD = s.Gain.GainDMean*s.BattingAverage + s.Loss.GainDMean*(1.0-s.BattingAverage)

	s.KellyCriterionD = s.BattingAverage - ((1 - s.BattingAverage) / s.WinLossRatioD)

	return s, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *Statistic) FormatFloat(data float64) string {
	return fmt.Sprintf("%.4f", data)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
