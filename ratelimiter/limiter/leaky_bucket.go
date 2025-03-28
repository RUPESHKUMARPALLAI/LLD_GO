package limiter

import (
	"sync"
	"time"
)

type TokenBucketLimiter struct {
	capacity   int
	tokens     int
	refillRate time.Duration
	lastRefilled time.Time
	mutex sync.Mutex
}

func NewTokenBucketLimiter() *TokenBucketLimiter {
   return &TokenBucketLimiter{
	capaci
   }
}