package main

import "fmt"

func main(){

	a := 100
	b := 200
	var ret int
    ret = max(a,b)
    fmt.Printf("最大值是：%d\n", ret)


    x,y := swap("google", "webkit")
    fmt.Println(x,y)

}

//声明参数名称和类型 声明返回值类型
func max(num1, num2 int) int {
	var res int
	if (num1 > num2) {
		res = num1
	} else {
		res = num2
	}
	return res
}

//函数可以返回多个值
func swap(x, y string) (string, string) {
	return y, x
}


