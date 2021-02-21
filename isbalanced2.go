package main

func isBalanced2_1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(depth(root.Left)-depth(root.Right)) <= 1 && isBalanced2_1(root.Left) && isBalanced2_1(root.Right)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return min(depth(node.Left), depth(node.Right)) + 1
}

func isBalanced2_2(root *TreeNode) bool {
	return depth2(root) != -1
}

func depth2(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight, rightHeight := depth2(node.Left), depth2(node.Right)

	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) >= 2 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}
