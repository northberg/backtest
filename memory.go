package backtest

type Memory struct {
	storeFloat map[string]float64
	storeInt   map[string]int64
}

func NewMemory() *Memory {
	return &Memory{
		storeFloat: make(map[string]float64),
		storeInt:   make(map[string]int64),
	}
}

func (m *Memory) StoreInt64(key string, v int64) {
	m.storeInt[key] = v
}

func (m *Memory) ReadInt64(key string) int64 {
	return m.storeInt[key]
}

func (m *Memory) Store(key string, v int) {
	m.storeInt[key] = int64(v)
}

func (m *Memory) Read(key string) int {
	return int(m.storeInt[key])
}

func (m *Memory) StoreBool(key string, v bool) {
	if v {
		m.storeInt[key] = 1
	} else {
		m.storeInt[key] = 0
	}
}

func (m *Memory) ReadBool(key string) bool {
	if m.storeInt[key] == 0 {
		return false
	} else {
		return true
	}
}

func (m *Memory) StoreFloat(key string, v float64) {
	m.storeFloat[key] = v
}

func (m *Memory) ReadFloat(key string) float64 {
	return m.storeFloat[key]
}
