package code

import "math/rand"

// 实现RandomizedSet 类
// 1. insert(val) 当元素val 不存在时，插入val 并返回true，否则返回false
// 2. remove(val) 当元素val 存在时，删除val 并返回true，否则返回false
// 3. getRandom() 随机返回现有集合中的一个元素，每个元素被返回的概率应该相同
// 要求：所有操作的时间复杂度都是O(1)

type RandomizedSet struct {
	nums   []int       // 存储元素
	idxMap map[int]int // 存储元素的索引
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:   make([]int, 0),
		idxMap: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.idxMap[val]; !ok {
		this.nums = append(this.nums, val)
		this.idxMap[val] = len(this.nums) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.idxMap[val]; ok {
		// 将最后一个元素放到要删除的元素位置上
		lastIdx := len(this.nums) - 1
		lastVal := this.nums[lastIdx]
		this.nums[idx] = lastVal
		this.idxMap[lastVal] = idx
		// 删除最后一个元素
		this.nums = this.nums[:lastIdx]
		delete(this.idxMap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
