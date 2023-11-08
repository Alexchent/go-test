package sort

// PopSort 冒泡排序
// 思路：比较相邻的两个元素，如果前一个比后一个大，则交换位置
// 这样一轮下来，最大的元素就在最后面了
// 重复上述步骤，直到没有任何一对数字需要比较
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func PopSort(arr2 *[]int) {
	arr := *arr2
	num := len(arr)
	for i := 0; i < num-1; i++ {
		for j := 0; j < num-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// QuickSort 快速排序
// 思路：取一个基准值，将数组分为两部分，左边的都比基准值小，右边的都比基准值大
// 然后递归的对左右两部分进行快速排序
// 时间复杂度：O(nlogn)
// 空间复杂度：O(logn)
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int
	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pivot), right...)
}
