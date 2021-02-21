package main

func levelOrderBottom(root *TreeNode) [][]int {
	ans := [][]int{}

	if root == nil {
		return ans
	}

	queue := []*TreeNode{}

	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		ansItem := []int{}
		for i := 0; i < size; i++ {
			node := queue[0]
			val := node.Val
			ansItem = append(ansItem, val)

			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, ansItem)
	}

	length := len(ans)

	for i := 0; i < length/2; i++ {
		ans[i], ans[length-1-i] = ans[length-1-i], ans[i]
	}

	return ans
}
