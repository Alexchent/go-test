package main

import "fmt"

// JiFen 泛型切片
type JiFen[T int | float64] []T

// Arr 泛型map
type Arr[Key int | string, Value int | string] map[Key]Value

// HaoShenYin 泛型struct
type HaoShenYin[T string | int | float64] struct {
	Title string
	Value T
}

type MyStruct[S int | string, P map[S]string] struct {
	Name    string
	Content S
	Job     P
}

func main() {
	// 泛型函数调度
	s := Add[int](1, 2)
	fmt.Println(s)
	s1 := Add[float32](1.2, 1.3)
	fmt.Println(s1)
	// 特别的调用方式
	fmt.Println(Add(1, 2))

	j := JiFen[int]{1, 2}
	fmt.Println(j)
	fmt.Println(JiFen[float64]{1, 2})

	arr1 := make(Arr[int, string])
	arr1[1] = "hello"
	fmt.Println(arr1)

	arr2 := make(Arr[string, string])
	arr2["name"] = "刘诗诗"
	fmt.Println(arr2)

	hi := HaoShenYin[string]{Title: "Hello", Value: "kitty"}
	fmt.Println(hi)

	he := HaoShenYin[int]{Title: "Hello", Value: 100}
	fmt.Println(he)

	fmt.Println(MyStruct[string, map[string]string]{Name: "Hello", Content: "kitty", Job: map[string]string{"al": "kitty", "b": "good"}})

	fmt.Println(Sum(1.1, 2.0))
}

// Add 泛型函数声明
func Add[T int | float32 | float64](a, b T) T {
	return a + b
}

type myInt interface {
	int | int8 | int16 | int32 | int64
}

type myUint interface {
	uint | uint8 | uint16 | uint32
}

type myFloat interface{ float32 | float64 }

type MyNumber interface{ myInt | myUint | myFloat }

func Sum[T MyNumber](a, b T) T {
	return a + b
}
