package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 20

	/* 使用 if 语句判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	} else {
		fmt.Printf("a 大于等于20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)

	season := Season(7)
	fmt.Println(season)

	Loop()
}

func Season(month int) string {
	switch month {
	case 1, 2, 12:
		return "winter"
	case 3, 4, 5:
		return "spring"
	case 6, 7, 8:
		return "summer"
	case 9, 10, 11:
		return "Autumn"
	}
	return "unknown"
}

func Loop() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is %d, and j is : %d\n", i, j)
		}
	}
}
