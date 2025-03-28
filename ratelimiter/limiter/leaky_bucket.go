package limiter

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type LeakyBucketLimiter struct {
	capacity       int
	leakedInterval time.Duration
	queue          []*http.Request
	mutex          sync.Mutex
	ctx            context.Context
	cancelFunc     context.CancelFunc
}

func NewLeakyBucketLimiter(capacity int, leakedInterval time.Duration) *LeakyBucketLimiter {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &LeakyBucketLimiter{
		capacity:       capacity,
		leakedInterval: leakedInterval,
		queue:          make([]*http.Request, 0),
		ctx:            ctx,
		cancelFunc:     cancelFunc,
	}
}

func (l *LeakyBucketLimiter) Addrequest(r *http.Request) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if len(l.queue) < l.capacity {
		l.queue = append(l.queue, r)
		return true
	}
	return false
}

func (l *LeakyBucketLimiter) ProcessRequest() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if len(l.queue) > 0 {
		fmt.Println(l.queue[0])
		l.queue = l.queue[1:]
	}
}

func (l *LeakyBucketLimiter) StartLeaking() {
	ticker := time.NewTicker(l.leakedInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				l.ProcessRequest()
			case <-l.ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}

func (l *LeakyBucketLimiter) StopLeaking() {
	l.cancelFunc()
}

func (l *LeakyBucketLimiter) AllowRequestCheck() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if len(l.queue) < l.capacity {
		return true
	}
	return false
}