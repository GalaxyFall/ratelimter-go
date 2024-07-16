package ratelimter

import (
	"ratelimter/fixedWindow"
	"ratelimter/leakyBucket"
	"ratelimter/slidingWindow"
	"ratelimter/tokenBucket"
)

const (
	FixedWindowRateLimter = iota
	LeakyBucketRateLimter
	SlidingWindowRateLimter
	TokenBucketRateLimter
)

type RateLimter interface {
	Take() bool
	GetSize() int64
}

type Params struct {
	//window
	Length int64

	//bucket
	Capacity int64
	Rate     int64

	Limter int
}

func NewDefaultRateLimter(params Params) RateLimter {
	switch params.Limter {
	case FixedWindowRateLimter:
		return fixedWindow.NewFixedWindow(params.Length)
	case LeakyBucketRateLimter:
		return leakyBucket.NewLeakyBucket(params.Capacity, params.Rate)
	case SlidingWindowRateLimter:
		return slidingWindow.NewSlidingWindow(params.Length)
	case TokenBucketRateLimter:
		return tokenBucket.NewTokenBucket(params.Capacity, params.Rate)
	}
	return nil
}
