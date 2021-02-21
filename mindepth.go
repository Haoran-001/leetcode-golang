package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 左右子树如果有一棵为空, 则不参与比较, 直接返回另一棵非空子树+1
	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}

	if root.Left != nil && root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{}
	count := []int{}
	queue = append(queue, root)
	count = append(count, 1)

	for i := 0; i < len(queue); i++ {
		node := queue[i]
		depth := count[i]

		if node.Left == nil && node.Right == nil {
			return depth
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
			count = append(count, depth+1)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
			count = append(count, depth+1)
		}
	}

	return 0
}
