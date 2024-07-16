package slidingWindow

import (
	"container/list"
	"sync"
	"time"
)

type SlidingWindow struct {
	mu           sync.Mutex
	list         *list.List
	windowLength int64 //滑动窗口大小 表示这个时间间隔放行的数量
	interVal     int64 //窗口间隔  多少时间算一个窗口间隔 默认1s
}

func NewSlidingWindow(length int64, opts ...Option) *SlidingWindow {

	window := &SlidingWindow{
		mu:           sync.Mutex{},
		list:         list.New(),
		windowLength: length,
		interVal:     1000,
	}

	for _, opt := range opts {
		opt(window)
	}

	return window
}

type Option func(window *SlidingWindow)

/*
  窗口间隔 多少时间算一个窗口间隔
  val 100 表示 100ms内允许 windowLength个请求
*/
func WithMilliseconds(val int64) Option {
	return func(window *SlidingWindow) {
		window.interVal = val
	}
}

func (m *SlidingWindow) Take() bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	//如果窗口内现在的流量未到滑动窗口限制大小则放行  尾插首取
	if int64(m.list.Len()) < m.windowLength {
		m.list.PushFront(time.Now().UnixMilli())
		return true
	}
	//惰性清除窗口  如果窗口满了则查看最先进来的是否超过时间间隔
	backTime := m.list.Back().Value.(int64)
	// 未超出则限流
	if time.Now().UnixMilli()-backTime <= m.interVal {
		return false
	}
	//如果超出则代表最早的过期了 清除且放行
	m.list.Remove(m.list.Back())
	m.list.PushFront(time.Now().UnixMilli())

	return true
}

func (m *SlidingWindow) GetSize() int64 {
	m.mu.Lock()
	defer m.mu.Unlock()

	return int64(m.list.Len())
}
