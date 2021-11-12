package backtest

import (
	"github.com/northberg/candlestick"
)

type CollectionMeta struct {
	Symbols    []string       `json:"symbols"`
	StartBlock int64          `json:"startBlock"`
	EndBlock   int64          `json:"endBlock"`
	Bot        *BotParameters `json:"bot"`
}

type BotParameters struct {
	Repository string      `json:"repository"`
	Entry      string      `json:"entry"`
	Scenarios  [][]float64 `json:"scenarios"`
	Leverage   int         `json:"leverage"`
	Segments   int64       `json:"segments"`
}

type ResultCollection struct {
	Id      string                   `json:"id"`
	Meta    CollectionMeta           `json:"meta"`
	Symbols map[string]*SimResultSet `json:"symbols"`
}

type TradeStatistics struct {
	WinRate            float64 `json:"winRate"`
	ReturnOnInvestment float64 `json:"returnOnInvestment"`
	ProfitPerTrade     float64 `json:"profitPerTrade"`
	NumberOfTrades     int     `json:"numberOfTrades"`
	EntryExitRatio     float64 `json:"entryExitRatio"`
	ProfitAndLoss      float64 `json:"profitAndLoss"`
	FeeProfitRatio     float64 `json:"feeProfitRatio"`
}

type TradeMeta struct {
	Statistics TradeStatistics `json:"statistics"`
	StartBlock int64           `json:"startBlock"`
	EndBlock   int64           `json:"endBlock"`
	Symbol     string          `json:"symbol"`
}

type SimResultSet struct {
	Meta        TradeMeta     `json:"meta"`
	Simulations []*SimResults `json:"simulations"`
}

type SimResults struct {
	Meta      TradeMeta `json:"meta"`
	Scenarios [][]float64
	Segments  []*SegmentResults `json:"segments"`
}

type SegmentResults struct {
	Meta   TradeMeta
	Orders []*LogOrderEntry `json:"orders"`
	Trades []*LogTradeEntry `json:"trades"`
}

type LogOrderEntry struct {
	Id        string
	TimeStamp int64
	Side      candlestick.OrderSide
	Type      candlestick.OrderKind
	Price     float64
	Amount    float64
}

type LogTradeEntry struct {
	Id        string
	TimeStamp int64
	Side      candlestick.OrderSide
	Type      candlestick.OrderKind
	Price     float64
	Amount    float64
	Realized  float64
	Fees      float64
}
