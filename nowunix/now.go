package nowunix

import (
	"sync/atomic"
	"time"
)

var clock int64

// 每秒校准，返回秒级时间戳
func init() {
	clock = time.Now().Unix()

	go func() {
		for {
			atomic.StoreInt64(&clock, time.Now().Unix()) // 时间戳校准
			time.Sleep(time.Second)                      // sleep
		}
	}()
}

func Now() int64 { return atomic.LoadInt64(&clock) }

// 返回当前时间的 time.Time
func NowTime() time.Time {
	now := Now()
	if now >= 1e9 {
		return time.Unix(now/1e9, now%1e9)
	} else {
		return time.Unix(now, 0)
	}
}
