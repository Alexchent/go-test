package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// 广度优先遍历，找到两个节点的过程中，记录每个节点的父节点
	parent := make(map[*TreeNode]*TreeNode)
	parent = preOrder(root, parent)

	visited := map[*TreeNode]bool{}

	for p != nil {
		visited[parent[p]] = true
		p = parent[p]
	}
	for q != nil {
		if visited[parent[q]] {
			return q
		}
		q = parent[q]
	}
	return nil
}

func preOrder(root *TreeNode, parent map[*TreeNode]*TreeNode) map[*TreeNode]*TreeNode {
	if root == nil {
		return parent
	}
	if root.Left != nil {
		parent[root.Left] = root
		parent = preOrder(root.Left, parent)
	}
	if root.Right != nil {
		parent[root.Right] = root
		parent = preOrder(root.Right, parent)
	}
	return parent
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 广度优先遍历，找到两个节点的过程中，记录每个节点的父节点
	queue := []*TreeNode{root}
	parent := make(map[*TreeNode][]*TreeNode)
	parent[root] = queue
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			parent[node.Left] = append(parent[node], node)
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = append(parent[node], node)
			queue = append(queue, node.Right)
		}
	}
	plist := parent[p]
	qlist := parent[q]
	plen := len(parent[p])
	qlen := len(parent[q])
	if plen > qlen {
		plist = plist[:qlen]
	} else {
		qlist = qlist[:qlen]
	}

	for i := len(plist) - 1; i >= 0; i-- {
		if plist[i] == qlist[i] {
			return plist[i]
		}
	}
	return nil
}
