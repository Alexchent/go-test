package code

// https://leetcode.cn/problems/lru-cache/?envType=study-plan-v2&envId=top-interview-150

type LRUCache struct {
	cap        int
	cache      map[int]*TNode
	head, tail *TNode
}

type TNode struct {
	key, val  int
	next, pre *TNode
}

func ConstructorLRUCache(capacity int) LRUCache {
	head, tail := new(TNode), new(TNode)
	head.next = tail
	tail.pre = head

	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*TNode, capacity),
		head:  head,
		tail:  tail,
	}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.moveToFont(n)
	return n.val
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.cache[key]
	if ok {
		n.val = value
		this.moveToFont(n)
		return
	}

	if len(this.cache) == this.cap {
		last := this.tail.pre
		this.remove(last)
		delete(this.cache, last.key)
	}
	node := &TNode{
		key: key,
		val: value,
	}
	this.pushFont(node)
	this.cache[key] = node
}

func (this *LRUCache) remove(n *TNode) {
	n.pre.next = n.next
	n.next.pre = n.pre
	n.pre = nil
	n.next = nil
}

func (this *LRUCache) pushFont(n *TNode) {
	n.pre = this.head
	n.next = this.head.next
	// 注意先修改next节点的pre
	this.head.next.pre = n
	this.head.next = n
}

func (this *LRUCache) moveToFont(n *TNode) {
	this.remove(n)
	this.pushFont(n)
}
