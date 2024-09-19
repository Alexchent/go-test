package graph

// 单词接龙 ToDo
// 字典 wordList 中从单词 beginWord 到 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 -> s2 -> ... -> sk：
//
// 每一对相邻的单词只差一个字母。
// 对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
// sk == endWord
// 给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0 。
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := map[string]bool{}
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return 0
	}
	minChange := 0 // 最少变化次数
	for i, c := range beginWord {
		if byte(c) != endWord[i] {
			minChange++
		}
	}
	if minChange < len(wordList) {
		return 0
	}
	// 开始广度优先遍历
	queue := []string{beginWord}
	for step := 0; queue != nil; step++ {
		list := queue
		queue = nil
		for _, cur := range list {
			for i, x := range cur {
				if byte(x) != endWord[i] {
					newC := cur[:i] + string(endWord[i]) + cur[i+1:]
					if wordSet[newC] {
						if newC == endWord {
							return step
						}
						delete(wordSet, newC)
						queue = append(queue, newC)
					}
				}
			}
		}
	}
	return 0
}
