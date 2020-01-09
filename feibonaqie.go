package main 

import "fmt"

func fi(n int) int {
	if (n < 2) {
		return n
	}

	return fi(n-1) + fi(n-2)
}

func main() {
	var i int
	for i=0; i<10; i++ {
		fmt.Printf("%d\t", fi(i))
	}
}