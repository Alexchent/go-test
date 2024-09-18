package graph

// 433. 最小基因变化
// 基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
//
// 假设我们需要调查从基因序列 start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
//
// 例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
// 另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。（变化后的基因必须位于基因库 bank 中）
//
// 给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end 所需的最少变化次数。如果无法完成此基因变化，返回 -1 。
//
// 注意：起始基因序列 start 默认是有效的，但是它并不一定会出现在基因库中。
func minMutation(startGene string, endGene string, bank []string) int {
	// hashmap存储基因库，用于检测基因变化是否合法
	// 初始基因`startGene`作为当前量
	// 从初始基因`current`开始，每次改变一个片段成为新基因`newGene`, 该片段向(A,C,G,T)个方向变化
	// 如果变化是有效的（在基因库中），那么这个新基因`newGene`作为当前量，再重复上面的步骤
	if startGene == endGene {
		return 0
	}
	// 基因库存储到hashmap
	bankSet := make(map[string]bool)
	for _, gen := range bank {
		bankSet[gen] = true
	}
	// 判断endGene是否在基因库中
	if !bankSet[endGene] {
		return -1
	}
	visited := map[string]bool{}
	return dfs(startGene, endGene, bankSet, visited, 0)
}

func dfs(current string, end string, genes map[string]bool, visited map[string]bool, step int) int {
	// 如果能找到，那边
	if current == end {
		return step
	}
	visited[current] = true
	minStep := -1
	// 变化基因序列上的每个点
	for i := 0; i < len(current); i++ {
		currentChar := current[i]
		for _, char := range []byte{'A', 'C', 'G', 'T'} {
			// 不同才算是发生变化
			if currentChar != char {
				// 生成新的基因
				newGene := current[:i] + string(char) + current[i:]
				// 新基因合法性验证
				if genes[newGene] && !visited[newGene] {
					stepForNew := dfs(newGene, end, genes, visited, step+1)
					if stepForNew != -1 && (minStep == -1 || stepForNew < minStep) {
						minStep = stepForNew
					}
				}
			}
		}
	}
	return minStep
}
