package backtest

type BotMetrics struct {
	Elapsed     int64 `json:"elapsed"`
	StartTime   int64 `json:"startTime"`
	TotalBlocks int   `json:"totalBlocks"`
	Progress    int   `json:"progress"`
	Running     bool  `json:"running"`
}

type BotRunVariables struct {
	Start     int64       `json:"startBlock"`
	End       int64       `json:"endBlock"`
	Segments  int64       `json:"segments"`
	Symbols   []string    `json:"symbols"`
	Scenarios [][]float64 `json:"scenarios"`
	Leverage  int         `json:"leverage"`
}