package main

import (
	"fmt"
	"time"
)

// 利用内存缓存提升函数运行效率
// 修改后的斐波那契数列性能有了质的飞跃

const LTM = 41

var fibs [LTM]uint64

func main() {
	start := time.Now()
	defer fmt.Println(time.Now().Sub(start))

	for i := 0; i < LTM; i++ {
		fmt.Println(fibonacci2(i))
	}
}

func fibonacci2(n int) (res uint64) {
	// 初始fibs [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	fibs[n] = res
	return
}
