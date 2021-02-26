package main

// 层序遍历
func maxDepth3(root *Node) int {
	if root == nil {
		return 0
	}

	queue := NodeQueue{}
	queue.Enqueue(root)
	depth := 0

	for !queue.Empty() {
		length := queue.Len()

		for i := 0; i < length; i++ {
			node := queue.Dequeue()
			if node.Children != nil {
				for i := 0; i < len(node.Children); i++ {
					queue.Enqueue(node.Children[i])
				}
			}
		}
		depth++
	}

	return depth
}

func maxDepth2(root *Node) int {
	if root == nil {
		return 0
	}

	result := 1

	for _, node := range root.Children {
		result = max(result, 1+maxDepth2(node))
	}

	return result
}
