package main

/** 字典树/前缀树
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 *
 */

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func NewWordDictionary() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	node := this
	for _, char := range word {
		c := char - 'a'
		if node.children[c] == nil {
			node.children[c] = &WordDictionary{}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(index int, node *WordDictionary) bool
	dfs = func(index int, node *WordDictionary) bool {
		// 搜索到底部
		if index == len(word) {
			return node.isEnd
		}
		ch := word[index]
		if ch != '.' {
			child := node.children[ch]
			if child != nil && dfs(index+1, child) {
				return true
			}
		} else {
			for i := range node.children {
				child := node.children[i]
				if child != nil && dfs(index+1, child) {
					return true
				}
			}
		}
		return false
	}

	return dfs(0, this)
}
