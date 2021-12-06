package backtest

type BotConfig struct {
	BotId      string    `json:"botId"`
	Funds      float64   `json:"funds"`
	Symbols    []string  `json:"symbols"`
	Parameters []float64 `json:"parameters"`
	Leverage   int       `json:"leverage"`
}

type BotVersion struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type BotInstance struct {
	Id        string     `json:"id"`
	Version   BotVersion `json:"version"`
	Name      string     `json:"name"`
	CreatedOn int64      `json:"createdOn"`
	StartTime int64      `json:"startTime"`
	StopTime  int64      `json:"stopTime"`
	Status    string     `json:"status"`
	Log       string     `json:"log"`
	Error     string     `json:"error"`
}
