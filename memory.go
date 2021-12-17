package backtest

type Memory struct {
	Floats   map[string]float64 `json:"floats"`
	Integers map[string]int64   `json:"integers"`
}

func NewMemory() *Memory {
	return &Memory{
		Floats:   make(map[string]float64),
		Integers: make(map[string]int64),
	}
}

func (m *Memory) StoreInt64(key string, v int64) {
	m.Integers[key] = v
}

func (m *Memory) ReadInt64(key string) int64 {
	return m.Integers[key]
}

func (m *Memory) Store(key string, v int) {
	m.Integers[key] = int64(v)
}

func (m *Memory) Read(key string) int {
	return int(m.Integers[key])
}

func (m *Memory) StoreBool(key string, v bool) {
	if v {
		m.Integers[key] = 1
	} else {
		m.Integers[key] = 0
	}
}

func (m *Memory) ReadBool(key string) bool {
	if m.Integers[key] == 0 {
		return false
	} else {
		return true
	}
}

func (m *Memory) StoreFloat(key string, v float64) {
	m.Floats[key] = v
}

func (m *Memory) ReadFloat(key string) float64 {
	return m.Floats[key]
}
