package backtest

type Simulation struct {
	Config     *SimConfig    `json:"config"`
	Instance   *BotInstance  `json:"instance"`
	LastMetric *SimMetrics   `json:"lastMetric"`
	Results    *MasterResult `json:"results"`
	Candidate  bool          `json:"candidate"`
}

type SimMetrics struct {
	Elapsed      int64 `json:"elapsed"`
	StartTime    int64 `json:"startTime"`
	TotalBlocks  int   `json:"totalBlocks"`
	Progress     int   `json:"progress"`
	Finished     bool  `json:"finished"`
	CurrentBlock int   `json:"currentBlock"`
	Running      bool  `json:"running"`
}

type SimConfig struct {
	Version   BotVersion  `json:"version"`
	Label     string      `json:"description"`
	Start     int64       `json:"startBlock"`
	End       int64       `json:"endBlock"`
	Symbols   []string    `json:"symbols"`
	Scenarios [][]float64 `json:"scenarios"`
	Leverage  int         `json:"leverage"`
	Segments  int64       `json:"segments"`
}
