package sort

// SortPop 冒泡排序
// 思路：比较相邻的两个元素，如果前一个比后一个大，则交换位置
// 这样一轮下来，最大的元素就在最后面了
// 重复上述步骤，直到没有任何一对数字需要比较
func SortPop(arr2 *[]int) {
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

// a, b链表，请链表焦点
// 思路：先求出两个链表的长度，然后让长的链表先走两个链表长度之差的步数
// 然后两个链表一起走，当两个链表的节点相同时，就是焦点
// 例如：
// a: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7
// b: 8 -> 9 -> 5 -> 6 -> 7
// a的长度为7，b的长度为5，a先走7-5=2步，然后两个链表一起走，当两个链表的节点相同时，就是焦点
type ListNode struct {
	Val  int
	Next *ListNode
}

func IntersectionNode(a, b *ListNode) *ListNode {
	aLen, bLen := 0, 0
	for p := a; p != nil; p = p.Next {
		aLen++
	}
	for p := b; p != nil; p = p.Next {
		bLen++
	}
	// 让长的链表先走两个链表长度之差的步数
	for aLen > bLen {
		a = a.Next
		aLen--
	}
	for bLen > aLen {
		b = b.Next
		bLen--
	}
	// 两个链表一起走，当两个链表的节点相同时，就是焦点
	for a != b {
		a = a.Next
		b = b.Next
	}
	return a
}
