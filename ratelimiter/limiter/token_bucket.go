package limiter

import (
	"math"
	"sync"
	"time"
)

type TokenBucketLimiter struct {
	capacity     int
	tokens       int
	refillRate   int //per second refill rate
	lastRefilled time.Time
	mutex        sync.Mutex
}

func NewTokenBucketLimiter(capacity int, refillRate int) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		capacity:     capacity,
		tokens:       capacity,
		refillRate:   refillRate,
		lastRefilled: time.Now(),
	}
}

func (t *TokenBucketLimiter) AllowRequestCheck() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	now := time.Now()
	timeElapsedSinceLastCheck := now.Sub(t.lastRefilled).Seconds()
	tokenstoAdd := int(timeElapsedSinceLastCheck) * t.refillRate
	t.tokens = int(math.Min(float64(t.capacity), float64(t.tokens+tokenstoAdd)))
	t.lastRefilled = t.lastRefilled.Add(time.Duration(int(timeElapsedSinceLastCheck)) * time.Second) ///this line is required to prevent time leakage

	if t.tokens > 0 {
		t.tokens--
		return true
	}
	return false
}
