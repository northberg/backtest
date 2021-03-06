package backtest

import (
	"github.com/northberg/candlestick"
)

type TradeStatistics struct {
	WinRate        float64 `json:"winRate"`
	NumberOfTrades int     `json:"numberOfTrades"`
	EntryExitRatio float64 `json:"entryExitRatio"`
	ProfitPerTrade float64 `json:"profitPerTrade"`
	Realized       float64 `json:"realized"`
	Unrealized     float64 `json:"unrealized"`
	FeeProfitRatio float64 `json:"feeProfitRatio"`
}

type MasterMeta struct {
	Id         string          `json:"id"`
	BotId      string          `json:"botId"`
	Name       string          `json:"name"`
	CreatedOn  int64           `json:"createdOn"`
	Symbols    []string        `json:"symbols"`
	Scenarios  [][]float64     `json:"scenarios"`
	Leverage   int             `json:"leverage"`
	Segments   int64           `json:"segments"`
	StartBlock int64           `json:"startBlock"`
	EndBlock   int64           `json:"endBlock"`
	Statistics TradeStatistics `json:"statistics"`
}

type MasterResult struct {
	Scenarios []*ScenarioResult `json:"scenarios"`
	Meta      MasterMeta        `json:"meta"`
}

type SymbolMeta struct {
	Statistics TradeStatistics `json:"statistics"`
	Symbol     string          `json:"symbol"`
}

type SymbolResult struct {
	Segments []*SegmentResult `json:"segments"`
	Meta     SymbolMeta       `json:"meta"`
}

type ScenarioMeta struct {
	Statistics TradeStatistics `json:"statistics"`
	Parameters []float64       `json:"parameters"`
}

type ScenarioResult struct {
	Symbols map[string]*SymbolResult `json:"symbols"`
	Meta    ScenarioMeta             `json:"meta"`
}

type SegmentMeta struct {
	StartBlock int64           `json:"startBlock"`
	EndBlock   int64           `json:"endBlock"`
	Statistics TradeStatistics `json:"statistics"`
}

type SegmentResult struct {
	Meta   SegmentMeta      `json:"meta"`
	Orders []*LogOrderEntry `json:"orders"`
	Trades []*LogTradeEntry `json:"trades"`
}

type LogOrderEntry struct {
	Id        string                   `json:"id"`
	TimeStamp int64                    `json:"timeStamp"`
	Side      candlestick.OrderSide    `json:"side"`
	Type      candlestick.OrderKind    `json:"type"`
	Hedge     candlestick.PositionSide `json:"hedge"`
	Price     float64                  `json:"price"`
	Amount    float64                  `json:"amount"`
}

type LogTradeEntry struct {
	Id        string                   `json:"id"`
	TimeStamp int64                    `json:"timeStamp"`
	Side      candlestick.OrderSide    `json:"side"`
	Type      candlestick.OrderKind    `json:"type"`
	Hedge     candlestick.PositionSide `json:"hedge"`
	Price     float64                  `json:"price"`
	Amount    float64                  `json:"amount"`
	Realized  float64                  `json:"realized"`
	Fees      float64                  `json:"fees"`
	Position  float64                  `json:"position"`
	Entry     float64                  `json:"entry"`
}
