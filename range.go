package main
import "fmt"

func main() {
	nums := []int{2,3,4}
	sum := 0
	//迭代 数组array、切片slice、通道channel或集合map的元素
	for index, num := range nums {
		sum += num
		if (num == 3) {
			fmt.Println("index:", index)
		}
	}
	fmt.Println("sum:", sum)
}