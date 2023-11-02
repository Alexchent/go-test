package util

import (
	"sync"
	"time"
)

// TokenLimit 令牌桶算法
// 弥补漏桶算法的问题
// 令牌桶算法的思路是这样的：系统会以一个恒定的速度往桶里放入令牌，而如果请求需要被处理，
// 则需要先从桶里获取一个令牌，当桶里没有令牌可取时，则拒绝服务
// 令牌桶算法的优点是可以方便的处理突发流量，且可以预留一定的资源给关键请求
// 令牌桶算法的缺点是实现相对复杂一些，且对短时间内的流量难以应对
type TokenLimit struct {
	// 令牌桶的容量
	capacity int
	// 当前令牌桶的令牌数量
	curToken int
	// 令牌放入速度，每秒放入多少个令牌
	rate int
	// 上次放入令牌的时间
	lastTime int64
	lock     sync.Mutex
}

func NewTokenLimit(capacity, rate int) *TokenLimit {
	return &TokenLimit{
		capacity: capacity,
		curToken: 0,
		rate:     rate,
		lastTime: time.Now().Unix(),
	}
}

func (t *TokenLimit) Allow() bool {
	t.lock.Lock()
	defer t.lock.Unlock()

	now := time.Now().Unix()
	// 计算从上次放入令牌到现在一共放入了多少令牌
	t.curToken += int((now - t.lastTime) * int64(t.rate))
	// 如果计算出来的令牌数量大于令牌桶的容量，则将令牌数量设置为令牌桶的容量
	if t.curToken > t.capacity {
		t.curToken = t.capacity
	}
	// 更新上次放入令牌的时间
	t.lastTime = now
	// 判断当前令牌桶的令牌数量是否大于0
	if t.curToken > 0 {
		// 如果大于0，则将令牌数量减1，并返回true
		t.curToken--
		return true
	}
	// 如果小于等于0，则返回false
	return false
}
