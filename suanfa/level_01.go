package main

// Fibonacci 斐波那契数列
func Fibonacci(n int) []int {
	if n == 0 {
		return []int{}
	} else if n == 1 {
		return []int{0}
	}

	sequence := []int{0, 1}

	for i := 2; i < n; i++ {
		next := sequence[i-1] + sequence[i-2]
		sequence = append(sequence, next)
	}

	return sequence
}
