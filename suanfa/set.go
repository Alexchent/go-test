package main

type Set struct {
	m map[interface{}]struct{}
}

// 创建一个空集合
func NewSet() *Set {
	return &Set{make(map[interface{}]struct{})}
}

// 添加元素
func (s *Set) Add(item interface{}) {
	s.m[item] = struct{}{}
}

// 删除元素
func (s *Set) Remove(item interface{}) {
	delete(s.m, item)
}

// 判断元素是否存在
func (s *Set) Contains(item interface{}) struct{} {
	return s.m[item]
}

// 返回集合大小
func (s *Set) Size() int {
	return len(s.m)
}
