package main

import "fmt"
import "unsafe"


func main(){
	const LENGTH = 2
	const WIDTH = 1.5
	var area int

	area = LENGTH*WIDTH

	
	fmt.Printf("面积为 : %d ", area)
	fmt.Println()

    //枚举类型
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	fmt.Println(a,b,c)

    //iota
    //iota，特殊常量，可以认为是一个可以被编译器修改的常量。
    //iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
    //const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
    const (
    	A = iota
    	B = iota
    	C = iota
    )
    
    fmt.Println(A,B,C)
}
