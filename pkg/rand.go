package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 使用当前时间的纳秒值作为种子
	// 要产生随机数，必须初始化，否则每次产生的数都一样
	//rand.Seed(time.Now().Unix())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成 0 到 100 之间的随机数
	fmt.Println(rand.Intn(101))
}
