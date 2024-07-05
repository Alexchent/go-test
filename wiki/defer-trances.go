package main

import "fmt"

// 延迟执行，注意如果是defer func其参数是预执行的
// 嵌套的defer先进后出
// defer 延迟执行，即使发生panic

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	//b()

	a := f()
	fmt.Println(a)
}

// 执行结果如下
//entering: b
//in b
//entering: a
//in a
//leaving: a
//leaving: b

// defer + 闭包
func f() (r int) {
	// defer 执行此时已经记录了参数 r = 0 ， 因此return 对r的更改对其没有影响
	defer func(r int) {
		r = r + 5
	}(r)
	return 1 // 返回 1
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5 // return后执行，此时t = 5， (t = t+5) === 10
	}()
	// 返回时 t = 5, 因为defer还没有执行
	return t
}
