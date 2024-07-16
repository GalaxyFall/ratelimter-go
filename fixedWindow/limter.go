package fixedWindow

import (
	"sync"
	"time"
)

/*
固定窗口:
每开启一个新的窗口，在窗口时间大小内，可以通过窗口请求上限个请求。
该算法主要是会存在临界问题，如果流量都集中在两个窗口的交界处，那么突发流量会是设置上限的两倍。
*/

type FixedWindow struct {
	mu           sync.Mutex
	windowLength int64 //窗口大小 窗口间隔内允许的请求数
	interVal     int64 //多久算一个时间间隔

	count    int64 //记录窗口内请求数
	lastTime time.Time
}

func NewFixedWindow(length int64, opts ...Option) *FixedWindow {

	window := &FixedWindow{
		mu:           sync.Mutex{},
		windowLength: length,
		interVal:     1000, //默认1s
		lastTime:     time.Now(),
	}

	for _, opt := range opts {
		opt(window)
	}

	return window
}

type Option func(window *FixedWindow)

/*
 固定窗口的时间间隔   val 100 表示 100ms一个间隔
*/
func WithMilliseconds(val int64) Option {

	return func(window *FixedWindow) {
		window.interVal = val
	}
}

func (m *FixedWindow) Take() bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	interVal := time.Now().Sub(m.lastTime).Milliseconds()
	//如果间隔超出则重新开一个窗口
	if interVal > m.interVal {
		m.count = 0
		m.lastTime = time.Now()
	}

	//如果当前固定窗口内处理的请求数超过则限流
	if m.count+1 > m.windowLength {
		return false
	}

	m.count++
	return true
}

func (m *FixedWindow) GetSize() int64 {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.count
}
