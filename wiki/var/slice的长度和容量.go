package main

import "fmt"

func main() {
	slice := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice1 := slice[5:7]
	fmt.Println(slice1)
	fmt.Printf("slice1的长度是%d,cap是%d\n", len(slice1), cap(slice1))

	slice2 := slice1[0:4]

	fmt.Println(slice2)

	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	//复制切片，增加切片的容量
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
