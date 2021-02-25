package main

var intSlice []int

func minDiffInBST(root *TreeNode) int {
	intSlice = make([]int, 0)
	nodesPreorder(root)

	minDistance := 2<<31 - 1

	for i := 1; i < len(intSlice); i++ {
		curDistance := intSlice[i] - intSlice[i-1]
		if curDistance < minDistance {
			minDistance = curDistance
		}
	}

	return minDistance
}

func nodesPreorder(node *TreeNode) {
	if node == nil {
		return
	}
	nodesPreorder(node.Left)
	intSlice = append(intSlice, node.Val)
	nodesPreorder(node.Right)
}
