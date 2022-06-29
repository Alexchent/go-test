package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("now:", time.Now())                           // 2022-06-20 15:01:23.750947 +0800 CST m=+0.000140004
	fmt.Println("now string:", time.Now().String())           // 2022-06-20 15:01:23.750947 +0800 CST m=+0.000140004
	fmt.Println("now format:", time.Now().Format("20060102")) // 20220620
	fmt.Println(time.Now().Date())                            // 2022 June 20
	fmt.Println("now unix:", time.Now().Unix())               // 1655708838
	fmt.Println("now month:", time.Now().Month())             // June
	fmt.Println("now day:", time.Now().Day())                 // 20
	fmt.Println("now hour:", time.Now().Hour())               // 15

	fmt.Println("ns：", time.Nanosecond)  // 1ns
	fmt.Println("s：", time.Second)       // 1s
	fmt.Println("ms：", time.Millisecond) // 1ms
	fmt.Println("minute：", time.Minute)  // 1ms

	fmt.Println("week：", time.Sunday) // Sunday
	fmt.Println("week：", time.Monday) // Monday

	fmt.Println("month：", time.May) // Sunday
}
