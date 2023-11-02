package util

import "time"

// CounterLimit 计数器限流
// 思路：使用一个计数器来记录当前时间窗口的请求数量，每次有请求过来时，就把计数器加1
// 然后判断当前计数器的值是否大于阈值，如果大于阈值，则拒绝请求
// 如果小于阈值，则接收请求
// 实现简单，但是没有很好的处理突增流量的能力
type CounterLimit struct {
	// 当前时间窗口的请求数量
	curCount int
	// 时间窗口的开始时间
	startTime int64
	// 时间窗口周期
	cycle time.Duration
	// 时间窗口内最大请求数量
	maxCount int
}

func NewCounterLimit(cycle time.Duration, maxCount int) *CounterLimit {
	return &CounterLimit{
		curCount:  0,
		startTime: time.Now().UnixNano(),
		cycle:     cycle,
		maxCount:  maxCount,
	}
}

func (c *CounterLimit) Allow() bool {
	// 判断当前时间是否已经超过了时间窗口的结束时间
	if time.Now().UnixNano()-c.startTime > int64(c.cycle) {
		// 如果超过了时间窗口的结束时间，则重置计数器和时间窗口的开始时间
		c.startTime = time.Now().UnixNano()
		c.curCount = 0
		return true
	}
	// 如果没有超过时间窗口的结束时间，则判断当前时间窗口的请求数量是否大于阈值
	if c.curCount >= c.maxCount {
		return false
	}
	// 如果没有超过阈值，则将当前时间窗口的请求数量加1
	c.curCount++
	return true
}

// SlideWindowLimit 滑动窗口限流
// 思路：使用一个数组来保存每个时间窗口的请求数量，每次有请求过来时，就把当前时间戳对应的数组下标的值加1
// 然后判断当前时间戳对应的数组下标的值是否大于阈值，如果大于阈值，则拒绝请求
// 如果小于阈值，则接收请求

// 漏桶算法
// 流量均匀，无法处理突增流量

// 令牌桶算法
// 弥补漏桶算法的问题
