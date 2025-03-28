package limiter

type RateLimiter interface {
	AllowRequestCheck() bool
}
