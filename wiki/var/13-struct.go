package main

import "fmt"

// 接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。

// Person 声明一个结构体类型
type Person struct {
	Name string
}

// Introduce 结构体方法
func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

// ChenJi 继承Person
type ChenJi struct {
	*Person
	Power int
}

func main() {
	goku := &ChenJi{
		Person: &Person{"alex"},
		Power:  9001,
	}

	goku.Introduce()

	fmt.Println(goku.Person.Name)
	fmt.Println(goku.Name)
}
