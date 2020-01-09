package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 这样写也可以，更像 While 语句形式
	sum = 1
    for sum <= 10{
                sum += sum
    }
    fmt.Println(sum)

    //for-each range 循环 可以对字符串、数组、切片进行迭代输出元素
    strings := []string{"google","runoob"}
    for i, s := range strings {
    	fmt.Println(i,s)
    }

    numbers := [6]int{1,2,3,5}
    for i,x := range numbers {
    	fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
    }

    /*结果

16
0 google
1 runoob
第 0 位 x 的值 = 1
第 1 位 x 的值 = 2
第 2 位 x 的值 = 3
第 3 位 x 的值 = 5
第 4 位 x 的值 = 0
第 5 位 x 的值 = 0

    */
}