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

type BotRunLog struct {
	Id        string `json:"id"`
	StartTime int64  `json:"startTime"`
	StopTime  int64  `json:"stopTime"`
	Output    string `json:"output"`
	Error     string `json:"error"`
}

type BotInstance struct {
	Id        string       `json:"id"`
	Name      string       `json:"name"`
	Status    string       `json:"status"`
	CreatedOn int64        `json:"createdOn"`
	Logs      []*BotRunLog `json:"logs"`
}

type TradeState struct {
	InstanceId string            `json:"instanceId"`
	StateId    string            `json:"stateId"`
	LastFill   float64           `json:"lastFill"`
	Balance    float64           `json:"balance"`
	Funds      float64           `json:"funds"`
	Symbol     string            `json:"symbol"`
	Time       int64             `json:"time"`
	Parameters []float64         `json:"parameters"`
	Memory     Memory            `json:"memory"`
	Leverage   int               `json:"leverage"`
	Position   cs.HedgedPosition `json:"position"`
}
