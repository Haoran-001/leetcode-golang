package main

func sumRootToLeaf(root *TreeNode) int {
	return sumRootToLeafPreorder(root, 0)
}

func sumRootToLeafPreorder(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}

	sum = sum<<2 + node.Val

	if node.Left == nil && node.Right == nil {
		return sum
	}

	leftSubTreeSum := sumRootToLeafPreorder(node.Left, sum)
	rightSubTreeSum := sumRootToLeafPreorder(node.Right, sum)

	return leftSubTreeSum + rightSubTreeSum
}
