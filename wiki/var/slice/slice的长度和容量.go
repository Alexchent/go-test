package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice[:0]) // []
	fmt.Println(slice[:1]) // [0]
	slice1 := slice[2:4]   // [2,3]
	fmt.Println(slice1)
	fmt.Printf("slice1的长度是%d,cap是%d\n", len(slice1), cap(slice1))

	slice2 := slice1[0:4]

	fmt.Println(slice2)

	sl_from := []int{1, 2, 3} // [1,2,3]
	sl_to := make([]int, 5)   // [0,0,0,0,0]
	//复制切片，增加切片的容量
	n := copy(sl_to, sl_from) // [1,2,3,0,0]
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	sl3 = append(sl3, []int{7, 8}...)
	fmt.Println(sl3)

	fmt.Println(strings.Repeat("-", 10) + "slice初始化" + strings.Repeat("-", 10))

	var a []int
	fmt.Println(a == nil) // true
	a = append(a, 1)
	fmt.Println(a)

	fmt.Println(strings.Repeat("-", 10))

	b := []int{}
	fmt.Println(b == nil) // false
	fmt.Println(b)
	b = append(b, 2)
	fmt.Println(b)

	fmt.Println(strings.Repeat("-", 10))

	c := make([]int, 0)
	fmt.Println(c == nil) // false
	c = append(c, 3)
	fmt.Println(c)
}
