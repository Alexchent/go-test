package unionfind

import "fmt"

// 990. 等式方程的可满足性
func EquationsPossible(equations []string) bool {
	// 并查集判断联通性
	union := NewUnionFind(26)

	for _, item := range equations {
		// 将所有可联通的进行合并
		if item[1] == '=' {
			a := int(item[0] - 'a')
			b := int(item[3] - 'a')
			union.union(a, b)
		}
	}
	fmt.Println(union.parent)
	for _, item := range equations {
		// 将所有可联通的进行合并
		if item[1] == '!' {
			a := int(item[0] - 'a')
			b := int(item[3] - 'a')
			if union.find(a) == union.find(b) {
				return false
			}
		}
	}

	return true
}

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	// 初始化，每个成员的代表元是自己本身
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &UnionFind{p}
}

// find 查找代表元
func (u *UnionFind) find(x int) int {
	// 退出条件，代表元是自身
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x]) // 路径压缩
	}
	return u.parent[x]
}

func (u *UnionFind) union(a, b int) {
	pa, pb := u.find(a), u.find(b)
	// 代表元不一样则进行合并
	if pa != pb {
		u.parent[pa] = pb
	}
}
