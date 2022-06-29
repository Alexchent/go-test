package main

import "fmt"

func main() {
	// 将切片b追加到切片a后面
	a := []int{1, 2, 3}
	b := []int{4, 5}

	a = append(a, b...)
	fmt.Println(a)
	a = append(a, 6)
	fmt.Println(a)

	// 复制切片a的元素到新的切片上
	c := make([]int, len(a))
	copy(c, a)
	fmt.Println(c)

	// 删除位于索引i的元素
	a = append(a[:0], a[1:]...)
	fmt.Println(a)

	// 切除出从i到j的元素
	a2 := append(a[:2], a[4:]...)
	fmt.Println(a2)
	fmt.Println(a) // 注意a切片的变化
}
