package backtest

import cs "github.com/northberg/candlestick"

type BotConfig struct {
	Version    BotVersion `json:"version"`
	Funds      float64    `json:"funds"`
	Symbols    []string   `json:"symbols"`
	Parameters []float64  `json:"parameters"`
	Leverage   int        `json:"leverage"`
}

type BotVersion struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Commit      string `json:"commit"`
	Description string `json:"description"`
}

type BotInstance struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedOn int64  `json:"createdOn"`
	StartTime int64  `json:"startTime"`
	StopTime  int64  `json:"stopTime"`
	Status    string `json:"status"`
	Log       string `json:"log"`
	Error     string `json:"error"`
}

type TradeState struct {
	InstanceId string            `json:"instanceId"`
	StateId    string            `json:"stateId"`
	Balance    float64           `json:"balance"`
	Funds      float64           `json:"funds"`
	Symbol     string            `json:"symbol"`
	Time       int64             `json:"time"`
	Parameters []float64         `json:"parameters"`
	Memory     Memory            `json:"memory"`
	Leverage   int               `json:"leverage"`
	Position   cs.HedgedPosition `json:"position"`
}
