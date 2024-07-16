package tokenBucket

import (
	"sync"
	"time"
)

/*
令牌桶算法：
以恒定的速率把令牌放到桶中，当令牌桶满了拒绝请求。
允许一定程度的突然性流量。
*/

type TokenBucket struct {
	mu       sync.Mutex
	capacity int64
	currSize int64
	rate     int64 //放令牌速率  默认rate/s
	interval int64
	lastTime time.Time
}

func NewTokenBucket(capacity, rate int64, opts ...Option) *TokenBucket {

	bucket := &TokenBucket{
		mu:       sync.Mutex{},
		capacity: capacity,
		currSize: capacity, //默认是满的
		rate:     rate,
		interval: 1000, //default 1s
		lastTime: time.Now(),
	}

	for _, opt := range opts {
		opt(bucket)
	}

	return bucket
}

type Option func(bucket *TokenBucket)

func WithSize(size int64) Option {
	return func(bucket *TokenBucket) {
		bucket.currSize = size
	}
}

/*
 令牌桶的放令牌的时间间隔   val 100 表示 100ms放一次
*/
func WithMilliseconds(val int64) Option {
	return func(bucket *TokenBucket) {
		bucket.interval = val
	}
}

func (m *TokenBucket) Take() bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	//惰性放令牌 最大不超过容量
	interval := time.Now().Sub(m.lastTime).Milliseconds()
	if interval > m.interval {
		m.lastTime = time.Now()
		m.currSize = maxInt64(m.capacity, m.currSize+m.rate*(m.interval/interval))
	}

	//尝试取出令牌  如果有令牌则放行
	if m.currSize-1 >= 0 {
		m.currSize--
		return true
	}

	return false
}

func (m *TokenBucket) GetSize() int64 {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.currSize
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}
