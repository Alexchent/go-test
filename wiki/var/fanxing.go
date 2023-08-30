package main

import "fmt"

func main() {
	// 泛型函数调度
	s := Add[int](1, 2)
	fmt.Println(s)

	s1 := Add[float32](1.2, 1.3)
	fmt.Println(s1)
}

// Add 泛型函数声明
func Add[T int | float32 | float64](a, b T) T {
	return a + b
}
