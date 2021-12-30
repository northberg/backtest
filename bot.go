package backtest

import (
	cs "github.com/northberg/candlestick"
	"strconv"
	"time"
)

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

func (l *BotRunLog) HasError() bool {
	return l.Error != ""
}

func (l *BotRunLog) Elapsed() int64 {
	if l.StopTime == 0 {
		return time.Now().UTC().Unix() - l.StartTime
	} else {
		return l.StopTime - l.StartTime
	}
}

func NewBotInstance(id string, name string) *BotInstance {
	return &BotInstance{
		Id:        id,
		Name:      name,
		Status:    "",
		CreatedOn: time.Now().UTC().Unix(),
		Logs:      make([]*BotRunLog, 0),
	}
}

type BotInstance struct {
	Id        string       `json:"id"`
	Name      string       `json:"name"`
	Status    string       `json:"status"`
	CreatedOn int64        `json:"createdOn"`
	Logs      []*BotRunLog `json:"logs"`
}

func (i *BotInstance) LastLog() *BotRunLog {
	if len(i.Logs) == 0 {
		return nil
	}
	return i.Logs[len(i.Logs)-1]
}

func (i *BotInstance) NewRun() *BotRunLog {
	logger := new(BotRunLog)
	logger.Id = i.Id + ":" + strconv.Itoa(len(i.Logs))
	logger.StartTime = time.Now().UTC().Unix()
	i.Logs = append(i.Logs, logger)
	return logger
}

type TradeState struct {
	InstanceId string            `json:"instanceId"`
	BotId      string            `json:"botId"`
	StateId    string            `json:"stateId"`
	LastFill   float64           `json:"lastFill"`
	Balance    float64           `json:"balance"`
	Funds      float64           `json:"funds"`
	Symbol     string            `json:"symbol"`
	Time       int64             `json:"time"`
	Parameters []float64         `json:"parameters"`
	Memory     *Memory           `json:"memory"`
	Leverage   int               `json:"leverage"`
	Position   cs.HedgedPosition `json:"position"`
}
