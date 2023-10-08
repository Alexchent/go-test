package main

import "fmt"

func main() {
	var slice1 []int
	fmt.Println(slice1 == nil) // true

	slice2 := []int{}
	//slice2 = append(slice2, 1)
	fmt.Println(slice2, len(slice2), cap(slice2)) // []

	slice3 := make([]int, 0)
	//slice3 = append(slice3, 2)
	fmt.Println(slice3, len(slice3), cap(slice3)) // []

}
