package code

type Node struct {
	Val  int
	Next *Node
}

// ReverseListNode Reverse 链表反转
func ReverseListNode(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 初始化一个前置节点
	var pre *Node
	for head != nil {
		// 设置当前节点的下一个节点为前置节点
		head.Next, pre = pre, head
		// 移动到下一个节点
		head = head.Next
	}
	return pre
}

// IsIntersect 判断两个链表是否相交
// 思路：如果两个链表相交，那么相交点之后的长度是相同的，所以先计算两个链表的长度，
// 然后长的链表先走，直到两个链表的长度相同，然后两个链表同时走，如果相交，那么一定会在相交点相遇
func IsIntersect(head1, head2 *Node) bool {
	if head1 == nil || head2 == nil {
		return false
	}

	// 两个链表的长度
	len1, len2 := 0, 0
	for head1 != nil {
		len1++
		head1 = head1.Next
	}

	for head2 != nil {
		len2++
		head2 = head2.Next
	}

	// 长的链表先走
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			head1 = head1.Next
		}
	} else {
		for i := 0; i < len2-len1; i++ {
			head2 = head2.Next
		}
	}

	// 两个链表同时走
	for head1 != nil && head2 != nil {
		if head1 == head2 {
			return true
		}

		head1 = head1.Next
		head2 = head2.Next
	}

	return false
}

// IsCircle 判断链表是否有环
// 思路：使用快慢指针，如果有环，那么快指针一定会追上慢指针
func IsCircle(head *Node) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将两个升序链表合并为一个升序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	node := &ListNode{}
	res := node // 当前节点
	for list1 != nil && list2 != nil {
		if list2.Val >= list1.Val {
			res.Next = list1
			list1 = list1.Next
		} else {
			res.Next = list2
			list2 = list2.Next
		}
		res = res.Next // res移动到下一个节点
	}
	// 此时还有一个链表没有走完，将其拼接到结果链表
	if list1 != nil {
		res.Next = list1
	} else {
		res.Next = list2
	}
	return node.Next
}
