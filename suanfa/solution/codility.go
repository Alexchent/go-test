package solution

import (
	"fmt"
	"strconv"
)

// 10进制转为2进制，判断二进制中连续0的最大个数
// 9 -> 1001 -> 2
// 529 -> 1000010001 -> 4
// 20 -> 10100 -> 1
// 15 -> 1111 -> 0
func BinaryGap(N int) int {
	binary := DecimalToBinary(N)
	maxGap := 0
	currentGap := 0
	for _, v := range binary {
		if v == '0' {
			currentGap++
		} else {
			if currentGap > maxGap {
				maxGap = currentGap
			}
			currentGap = 0
		}
	}
	return maxGap
}

// 10进制转为2进制
func DecimalToBinary(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}

	for n > 0 {
		reminder := n % 2
		result = strconv.Itoa(reminder) + result
		n = n / 2
	}
	return result
}

// 给定一个包含N个整数的数组A，返回数组中未出现的最小正整数（> 0）。
// 例如，给定A=[1,3,6,4,1,2]，返回5。
// 给定A=[1,2,3]，返回4。
// 给定A=[-1,-3]，返回1。
func FindMissingInteger(A []int) int {
	B := make([]int, 100000)
	for _, v := range A {
		if v > 0 {
			B[v-1] = 1
		}
	}

	for i, vb := range B {
		if vb != 1 {
			return i + 1
		}
	}
	return 0
}

// 节省内存空间的解法
func FindMissing2(A []int) int {
	res := 1
	maxlen := len(A)
	// 从头部开始比对，如果
	for res < maxlen {
		if res == A[res-1] {
			res++
			continue
		}
		if res < A[res-1] {
			A = append(A, A[res-1])
			maxlen++
		}
	}
	return res
}

// SumZeroSlice 一个由N个整数组成的数组A，有坐标P>=Q, 如果P到Q之间的元素之和是0，则称为A的一个零和区间
// 例如，数组A = [2, -2, 3, 0, 4, -7]，有5个零和区间，分别是[0,0]，[2,3]，[0,5]，[3,3]，[4,5]
// 写一个函数，给定一个由N个整数组成的数组A，返回A中零和区间的个数
// 例如，给定数组A = [2, -2, 3, 0, 4, -7]，函数应该返回5，如上述解释
// 假定：
// N是一个整数，范围[0..100,000]
// 数组A的每个元素的取值范围[-10,000..10,000]
// 在你的解决方案中，关注正确性。你的解决方案的性能不是评估的重点。
// Compare this snippet from test/solution_test.go:
func SumZeroSlice(A []int) int {
	var count int
	for i := 0; i < len(A); i++ {
		var sum int
		for j := i; j < len(A); j++ {
			sum += A[j]
			if sum == 0 {
				count++
			}
		}
	}
	return count
}

func ReverseInt(N int) {
	var enablePrint int
	enablePrint = N % 10
	for N > 0 {
		if enablePrint == 0 && (N/10)%10 != 0 {
			enablePrint = 1
		} else if enablePrint == 1 {
			fmt.Print(N % 10)
		}
		N = N / 10
	}
}
