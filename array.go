package main

import "fmt"

func main(){
	/* n是一个长度为10，元素类型为int的数组*/
	var n [10]int

	var i int

	for i = 0; i<10; i++ {
		//修改元素的值
		n[i] = i +100
		fmt.Printf("%d 位的值为 %d\n", i, n[i])
	}
}