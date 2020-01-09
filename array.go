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


    //声明一个多维数组
	var a = [5][2]int {{0,0},{1,2},{2,4},{3,6},{4,8}}
    var j int
	for  i = 0; i < 5; i++ {
      for j = 0; j < 2; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
      }
   }
}