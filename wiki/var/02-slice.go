package main

import "fmt"

// 切片有称为动态数组，相对于数组而言，它的长度是不固定的，可以动态的追加和删除
func main() {
	//初始化一个切片
	s := []int{1, 2, 3}
	fmt.Println(s)
	fmt.Printf("len:%d, cap:%d, s:%v\n", len(s), cap(s), s)

	//初始化一个长度为3，容量为5的切片
	k := make([]int, 3, 5)
	fmt.Println(k)

	fmt.Println(s[1:]) // {2,3}
	fmt.Println(s[:])  // {1,2,3}
	fmt.Println(s[:2]) // {1,2}

	//像切片中追加元素，注意数据类型必须一样
	s = append(s, 4)
	s = append(s, 5, 6) //可以同时追加多个
	fmt.Println(s)

	test713()
}

//假设有字符串 str，那么 str[len(str)/2:] + str[:len(str)/2] 的结果是什么？
func test713() {
	str := "0123456"
	str3 := str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(str3) // 3456012
}
