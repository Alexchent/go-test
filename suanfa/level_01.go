package main

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

var factorialCache = map[int]int{}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}

	if result, ok := factorialCache[n]; ok {
		return result
	}

	result := n * Factorial(n-1)
	factorialCache[n] = result
	return result
}
