package main

// interface 实现多态
import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("手续费 0%")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("手续费 15%")
}

func main() {
	nokia := new(NokiaPhone)
	hello(nokia)

	iphone := new(IPhone)
	hello(iphone)
}

func hello(p Phone) {
	p.call()
}
