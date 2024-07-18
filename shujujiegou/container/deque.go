package container

import "fmt"

// 基于数组实现的双向队列
type arrayDeque struct {
	nums    []int // 用于存储双向队列元素的数组
	front   int   // 队首指针，指向队首元素
	queSize int   // 队列长度
	cap     int   // 队列容量
}

// 初始化队列
func newArrayDeque(cap int) *arrayDeque {
	return &arrayDeque{
		nums:    make([]int, cap),
		front:   0,
		queSize: 0,
		cap:     cap,
	}
}

// 获取队列长度
func (d *arrayDeque) size() int {
	return d.queSize
}

// 判断队列是否为空
func (d *arrayDeque) isEmpty() bool {
	return d.queSize == 0
}

// 计算环形数组索引
func (d *arrayDeque) index(i int) int {
	// 取余实现数组首尾相连
	return (i + d.cap) % d.cap
}

// 队首入队
func (d *arrayDeque) pushLeft(num int) {
	if d.queSize == d.cap {
		fmt.Println("队列满了")
		return
	}

	d.front = d.index(d.front - 1) // 移动队首指针，左移
	d.nums[d.front] = num
	d.queSize++
}

// 队首出队
func (d *arrayDeque) popLeft() any {
	num := d.peekLeft()
	d.front = d.index(d.front + 1)
	d.queSize--
	return num
}

// 队尾入队
func (d *arrayDeque) pushRight(num int) {
	if d.queSize == d.cap {
		fmt.Println("队列满了")
		return
	}
	rear := d.index(d.front + d.queSize)
	d.nums[rear] = num
	d.queSize++
}

// 队尾出队
func (d *arrayDeque) popRight() any {
	num := d.peekRight()
	d.queSize--
	return num
}

// 访问队首元素
func (d *arrayDeque) peekLeft() any {
	if d.isEmpty() {
		return nil
	}

	return d.nums[d.front]
}

// 访问队尾元素
func (d *arrayDeque) peekRight() any {
	if d.isEmpty() {
		return nil
	}
	last := d.index(d.front + d.queSize - 1)
	return d.nums[last]
}

func (d *arrayDeque) toSlice() {
	fmt.Println(d)
}
