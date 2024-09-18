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
func minMutation2(startGene string, endGene string, bank []string) int {
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
	// 广度优先遍历
	q := []string{startGene}
	for step := 0; q != nil; step++ {
		list := q
		q = nil
		for _, cur := range list {
			for i, char := range cur {
				for _, c := range "ACGT" {
					if char != c {
						newGene := cur[:i] + string(c) + cur[i+1:]
						if bankSet[newGene] {
							if newGene == endGene {
								return step + 1
							}
							delete(bankSet, newGene)
							q = append(q, newGene)
						}
					}
				}
			}
		}
	}
	return -1
}
