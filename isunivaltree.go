package main

var set map[int]bool

func isUnivalTree(root *TreeNode) bool {
	set = make(map[int]bool, 0)
	isUinvalTreePreorder(root)
	return len(set) == 1
}

func isUinvalTreePreorder(node *TreeNode) {
	if node == nil {
		return
	}

	set[node.Val] = true
	isUinvalTreePreorder(node.Left)
	isUinvalTreePreorder(node.Right)
}

func isUnivalTree2(root *TreeNode) bool {
	return isUinvalTreePreorder2(root, root.Val)
}

func isUinvalTreePreorder2(node *TreeNode, x int) bool {
	if node == nil {
		return true
	}

	isLeftSubTree := isUinvalTreePreorder2(node.Left, x)
	isRightSubTree := isUinvalTreePreorder2(node.Right, x)

	return node.Val == x && isLeftSubTree && isRightSubTree
}
