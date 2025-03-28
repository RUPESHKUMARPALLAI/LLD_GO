package limiter

import "time"

func NewRateLimiter(limiterType string, capacity int, interval time.Duration) RateLimiter {
	switch limiterType {
	case "leaky":
		return NewLeakyBucketLimiter(capacity, interval)
	case "token":
		return NewTokenBucketLimiter(capacity, interval)
	default:
		return nil
	}
}
