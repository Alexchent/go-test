package solution

// You are given an implementation of a function Solution that,given a positive integer N,
//prints to standard output another integer using the digits of N,which was formed by reversing a decimal representation of N.
// The leading zeros of the resulting integer should be not be printed by the function.
// For example, given N = 123 the function should print 321, and given N = 100 the function should print 1.
// Assume that:
// N is an integer within the range [1..1,000,000,000].
// In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.
// Compare this snippet from test/solution_test.go:

//func Solution(N int) {
//	var enable_print int
//	enable_print = N % 10
//	for N > 0 {
//		if enable_print == 0 && N%10 != 0 {
//			enable_print = 1
//		} else if enable_print == 1 {
//			fmt.Print(N % 10)
//		}
//		N = N / 10
//	}
//}

func Solution(A []int) int {
	if len(A) >= 100000 {
		return -1
	}

	var count int
	allzero := true
	for i := 0; i < len(A); i++ {
		if A[i] != 0 {
			allzero = false
		}
		var sum int
		for j := i; j < len(A); j++ {
			sum += A[j]
			if sum == 0 {
				count++
				if count >= 1000000000 {
					return -1
				}
			}
		}
	}
	if allzero {
		return -1
	}
	return count
}

// An infrastrunture consisting of N cities numbered from 1 to N, and M bidirectional roads between them is given.
// Roads do not intersect apart from at their start and end points (they can pass through underground tunnels to avoid collisions).
// For each pair of cities directly connected by a road, let's define their network rank as the total number of roads that are connected to either of the two cities.
// Write a function:
// func Solution(N int, A []int, B []int) int
// that, given two arrays A and B consisting of M integers each and an integer N, where A[i] and B[i] are cities at the two ends of the i-th road,
// returns the maximal network rank in the whole infrastructure.
// The function should return −1 if the infrastructure has no roads.
// For example, given N = 4, A = [1, 2, 3, 3], B = [2, 3, 1, 4], the function should return 4,
// as explained above.
// Given N = 4, A = [1, 2, 4, 5], B = [2, 3, 5, 6], the function should return −1,
// as the cities number 4 and 6 are not connected.
// Assume that:
// N is an integer within the range [1..100,000];
// M is an integer within the range [0..200,000];
// each element of arrays A, B is an integer within the range [1..N];
// A[i] ≤ B[i].
// In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.
// Compare this snippet from test/solution_test.go:
//func TestSolution(t *testing.T) {
//	if p := Solution(4, []int{1, 2, 3, 3}, []int{2, 3, 1, 4}); p != 4 {
//		t.Errorf("no pass 1 but %+v got", p)
//	}
//	if p := Solution(4, []int{1, 2, 4, 5}, []int{2, 3, 5, 6}); p != -1 {
//		t.Errorf("no pass 2 but %+v got", p)
//	}
//}

func Solution2(N int, A []int, B []int) int {
	var max int
	for i := 0; i < len(A); i++ {
		var count int
		for j := 0; j < len(A); j++ {
			if A[i] == A[j] || A[i] == B[j] || B[i] == A[j] || B[i] == B[j] {
				count++
			}
		}
		if count > max {
			max = count
		}
	}
	return max
}
