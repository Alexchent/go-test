package hs

// 79、单词搜索
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
func exist(board [][]string, word string) bool {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	h, w := len(board), len(board[0])

	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	var check func(x, y, k int) bool
	check = func(x, y, k int) bool {
		// 搜索的位置不符合，剪枝
		if board[x][y] != string(word[k]) {
			return false
		}
		// 完成一路搜索
		if k == len(word)-1 {
			return true
		}
		visited[x][y] = true
		defer func() {
			visited[x][y] = false
		}()
		// 对上下左右四个方向进行深度优先遍历
		for _, d := range directions {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < h && ny >= 0 && ny < w && !visited[nx][ny] {
				if check(nx, ny, k) {
					return true
				}
			}
		}
		return false
	}

	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
