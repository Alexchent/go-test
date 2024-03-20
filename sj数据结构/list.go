package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 将list作为双向队列使用
	deque := list.New()

	// 元素入队
	deque.PushBack(21)  // 添加到队尾
	deque.PushBack(22)  // 添加到队尾
	deque.PushFront(11) // 添加到对首
	deque.PushFront(10) // 添加到对首

	size := deque.Len()
	fmt.Println(size)

	for size > 0 {
		a := deque.Front() // 访问队首
		deque.Remove(a)    // 弹出元素

		fmt.Println(a.Value)
		size--
	}

}
