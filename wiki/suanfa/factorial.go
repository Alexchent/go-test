package sunfa

// factorial 阶乘
//func main() {
//	fmt.Println(pow(2, 5))
//}

func pow(x, n int) int {
	if n == 1 {
		return x
	}

	return pow(x, n-1) * x
}
