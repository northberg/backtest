package backtest

import (
	"github.com/northberg/candlestick"
)

type CollectionMeta struct {
	Name       string          `json:"name"`
	CreatedOn  int64           `json:"createdOn"`
	Symbols    []string        `json:"symbols"`
	StartBlock int64           `json:"startBlock"`
	EndBlock   int64           `json:"endBlock"`
	Bot        *BotParameters  `json:"bot"`
	Statistics TradeStatistics `json:"statistics"`
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
	Meta      TradeMeta         `json:"meta"`
	Segments  []*SegmentResults `json:"segments"`
}

type SegmentResults struct {
	Meta   TradeMeta        `json:"meta"`
	Orders []*LogOrderEntry `json:"orders"`
	Trades []*LogTradeEntry `json:"trades"`
}

type LogOrderEntry struct {
	Id        string                `json:"id"`
	TimeStamp int64                 `json:"timeStamp"`
	Side      candlestick.OrderSide `json:"side"`
	Type      candlestick.OrderKind `json:"type"`
	Price     float64               `json:"price"`
	Amount    float64               `json:"amount"`
}

type LogTradeEntry struct {
	Id        string                `json:"id"`
	TimeStamp int64                 `json:"timeStamp"`
	Side      candlestick.OrderSide `json:"side"`
	Type      candlestick.OrderKind `json:"type"`
	Price     float64               `json:"price"`
	Amount    float64               `json:"amount"`
	Realized  float64               `json:"realized"`
	Fees      float64               `json:"fees"`
}
