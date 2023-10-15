package code

import "fmt"

// 在有N个不重复元素的数组中，元素的取值范围是[1,N+1]，找出数组中缺失的元素

func FindEle(A []int) int {
	if len(A) == 0 {
		return 0
	}

	B := make([]int, len(A)+1)

	for _, v := range A {
		B[v-1] = 1
	}
	fmt.Println(B)
	for i, v := range B {
		if v == 0 {
			return i + 1
		}
	}
	return 0
}
