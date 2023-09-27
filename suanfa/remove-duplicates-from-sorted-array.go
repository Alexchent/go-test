package main

import "fmt"

// 双指针经典算法

// 去除有序数组中的重复项, 返回移除后新数组的长度
// 快慢指针, 快指针发现与慢指针对应的数值不重复，则慢指针前进一位，同时修改原数组对应的数值
func RemoveDuplicatesFromSortedArray(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	fast, slow := 0, 0
	for fast < len(arr) {
		if arr[fast] != arr[slow] {
			slow++
			arr[slow] = arr[fast]
		}
		fast++
	}

	return arr[0:slow]
}

type ListNode struct {
	Val  string
	key  string
	Next *ListNode
}

// deleteDuplicates 函数实现从排序链表中删除重复元素。
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 初始化双指针slow, fast
	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			// slow指针与快指针相等的值赋值为快指针的值
			slow.Next = fast
			// 将指针后移，操作下一个节点
			slow = slow.Next
		}
		// 将指针后移，操作下一个节点
		fast = fast.Next
	}
	// 断开与后面重复元素的连接
	slow.Next = nil
	return head
}

func main() {
	arr := []int{1, 1, 1, 2, 2, 3, 4, 5, 5, 6}
	res := RemoveDuplicatesFromSortedArray(arr)
	fmt.Println(res)
}

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
func moveZeroesFast(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
