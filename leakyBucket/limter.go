package leakyBucket

import (
	"sync"
	"time"
)

/*
漏桶算法：
流量整形或速率限制控制数据注入到网络的速率，平滑网络上的突发流量。
将请求放入漏桶中，漏通满了请求会被拒绝,以恒定的速度将请求从桶中拿出，这里使用惰性方案可以使用协程定时清除。
*/

type LeakyBucket struct {
	capacity int64
	currSize int64
	rate     int64
	interval int64

	lastTime time.Time

	mu sync.Mutex
}

func NewLeakyBucket(capacity, rate int64, opts ...Option) *LeakyBucket {

	bucket := &LeakyBucket{
		capacity: capacity,
		lastTime: time.Now(),
		rate:     rate,
		interval: 1000, //default 1s
		mu:       sync.Mutex{},
	}

	for _, opt := range opts {
		opt(bucket)
	}
	return bucket
}

type Option func(bucket *LeakyBucket)

func WithSize(size int64) Option {
	return func(bucket *LeakyBucket) {
		bucket.currSize = size
	}
}

/*
 漏桶出水的时间间隔   val 100 表示 100ms放水一次
*/
func WithMilliseconds(val int64) Option {
	return func(bucket *LeakyBucket) {
		bucket.interval = val
	}
}

func (m *LeakyBucket) Take() bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	//惰性放水 如果当前距离上一次防水时间超过放水间隔则放水
	interval := time.Now().Sub(m.lastTime).Milliseconds()
	if interval > m.interval {
		m.lastTime = time.Now()
		//水量最小0
		m.currSize = minInt64(0, m.currSize-(m.rate*int64(interval/m.interval)))
	}

	//  当前大小不超过容量则放行
	if m.currSize+1 <= m.capacity {
		m.currSize++
		return true
	}

	//漏桶满了则拒绝请求
	return false
}

func (m *LeakyBucket) GetSize() int64 {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.currSize
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
