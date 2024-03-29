package nowunixnano

import (
	"sync/atomic"
	"time"
)

var clock int64

// 每 100毫秒 校准，返回纳秒级时间戳
func init() {
	clock = time.Now().UnixNano()

	go func() {
		for {
			atomic.StoreInt64(&clock, time.Now().UnixNano()) // 时间戳校准
			for i := 0; i < 9; i++ {                         // 时间戳累加
				time.Sleep(100 * time.Millisecond)
				atomic.AddInt64(&clock, int64(100*time.Millisecond))
			}
			time.Sleep(100 * time.Millisecond) // sleep
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
