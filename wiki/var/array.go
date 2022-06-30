package main

import "fmt"

func main() {
	//初始化7个int型成员的数组
	var week = [7]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(week[1])
	//修改数组成员
	week[1] = 10
	fmt.Println(week[1])

	//初始化7个string型成员的数组
	weekname := [7]string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}
	fmt.Println(weekname[1])

	//初始化不确定长度的int型数组
	userids := [...]int{1, 2, 3, 9, 2, 3, 6}
	fmt.Println(userids, len(userids), cap(userids))

	/* n是一个长度为10，元素类型为int的数组*/
	var n [10]int

	var i int

	for i = 0; i < 10; i++ {
		//修改元素的值
		n[i] = i + 100
		fmt.Printf("%d 位的值为 %d\n", i, n[i])
	}

	fmt.Println(n)

	//声明一个多维数组
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var j int
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
