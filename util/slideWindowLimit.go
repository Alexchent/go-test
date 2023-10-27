package util

type Node struct {
	curId   int
	counter int
}

// SlideWindowLimit 滑动窗口限流
// 思路：使用一个数组来保存每个时间窗口的请求数量，每次有请求过来时，就把当前时间戳对应的数组下标的值加1
// 然后判断当前时间戳对应的数组下标的值是否大于阈值，如果大于阈值，则拒绝请求

// 十进制转二进制
func DecimalToBinary(num int) string {
	if num == 0 {
		return "0"
	}
	res := ""
	for num > 0 {
		res = string(num%2+'0') + res
		num /= 2
	}
	return res
}

// 计算bitmap中1的个数
func CountOne(num int) int {
	count := 0
	for num > 0 {
		num &= num - 1
		count++
	}
	return count
}

// 计算bitmap中最长的连续1的个数
func CountOne2(num int) int {
	count := 0
	maxCount := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		} else {
			count = 0
		}
		if count > maxCount {
			maxCount = count
		}
		num >>= 1
	}
	return maxCount
}
