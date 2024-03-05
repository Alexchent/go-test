package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //要产生随机数，必须初始化，否则每次产生的数都一样
	fmt.Println(rand.Intn(5))
	fmt.Println(rand.Intn(5))
}
