package main

import "fmt"

// Animal 定义一个接口
type Animal interface {
	Eat(food string)
	getAge() uint8
}

type Cat struct {
	name string
	age  uint8
}

type dog struct {
	name string
	age  uint8
}

// Eat 结构体Cat的Eat方法
func (c *Cat) Eat(food string) {
	//panic("implement me")
	fmt.Println("cat eat: ", food)
	c.age++
}

func (d *dog) Eat(food string) {
	fmt.Println("dog eat: ", food)
	d.age += 2
}

// Age 结构体Cat的Age方法
func (c *Cat) getAge() uint8 {
	//panic("implement me")
	return c.age
}

func (d *dog) getAge() uint8 {
	//panic("implement me")
	return d.age
}

func Test(animal Animal) {
	animal.Eat("食物")
	fmt.Println(animal.getAge())
}

func main() {
	//var _ Animal = &dog{}
	//var _ Animal = (*dog)(nil)

	jojo := &Cat{
		name: "jojo",
		age:  3,
	}

	jinmao := &dog{
		name: "jinbo",
		age:  5,
	}

	Test(jojo)
	Test(jinmao)

	//var _ Animal = (*Cat)(nil)
}
