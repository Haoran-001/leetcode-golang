package main

func maxDepth(root *TreeNode) int {
	return nodeDepth(root)
}

func nodeDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(nodeDepth(node.Left), nodeDepth(node.Right)) + 1
}
