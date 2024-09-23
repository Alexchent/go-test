package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给字母编号
	id := map[string]int{}
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := id[a]; !has {
			id[a] = len(id)
		}
		if _, has := id[b]; !has {
			id[b] = len(id)
		}
	}

	u := &UnionFind{}
	for i, eq := range equations {
		u.Union(id[eq[0]], id[eq[1]], values[i])
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		a, oka := id[q[0]]
		b, okb := id[q[1]]
		if oka && okb && u.Find(a) == u.Find(b) {
			ans[i] = u.weight[a] / u.weight[b]
		} else {
			ans[i] = -1
		}
	}
	return ans
}

type UnionFind struct {
	parent []int     // 维护代表元
	weight []float64 // 维护权重
}

func NewUnionFind(n int) *UnionFind {
	fa := make([]int, n)
	we := make([]float64, n)
	for i := 0; i < n; i++ {
		fa[i] = i
		we[i] = 1
	}
	return &UnionFind{
		parent: fa,
		weight: we,
	}
}

func (u *UnionFind) Find(x int) int {
	if u.parent[x] != x {
		u.weight[x] *= u.weight[u.parent[x]] // 记录该元素到代表元的权重
		u.parent[x] = u.Find(u.parent[x])    // 路劲压缩，找到顶级作为该元素的代表元
	}
	return x
}

func (u *UnionFind) Union(x, y int, weight float64) {
	xp := u.Find(x)
	yp := u.Find(y)
	u.parent[xp] = yp
	u.weight[xp] = weight * u.weight[y] / u.weight[x]
}
