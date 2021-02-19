package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(height(node.Left), height(node.Right)) + 1
}

func isBalanced2(root *TreeNode) bool {
	return height2(root) >= 0
}

func height2(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := height2(node.Left)
	rightHeight := height2(node.Right)

	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}
