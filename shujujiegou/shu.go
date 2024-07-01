package main

// 广度优先遍历
// 借助队列来实现，队列先进先出，广度优先逐层递进
// 层序遍历
type TreeNode struct {
	val   any
	left  *TreeNode
	right *TreeNode
	root  *TreeNode // 根
}

// 二叉搜索树，定位节点
// 二叉搜索树遵循以下条件
// 1. 对于根节点，左子树中所有节点的值 < 根节点的值 < 右子树中所有节点的值。
// 2. 任意节点的左、右子树也是二叉搜索树，即同样满足条件 1.
func (bst *TreeNode) search(num int) *TreeNode {
	// 从根节点开始遍历
	node := bst.root
	for node != nil {
		if num > node.val.(int) {
			node = node.right
		} else if num < node.val.(int) {
			node = node.left
		} else {
			break
		}
	}
	return node
}
