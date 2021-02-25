package main

func isCousins(root *TreeNode, x int, y int) bool {
	depthX, parentNodeX := currentNodeDepth(root, x)
	depthY, parentNodeY := currentNodeDepth(root, y)

	return depthX == depthY && parentNodeX != parentNodeY
}

func currentNodeDepth(root *TreeNode, x int) (int, *TreeNode) {
	queue := TreeNodeQueue{}
	queue.Enqueue(root)
	depth := 0

	if root.Val == x {
		return -1, nil
	}

	for !queue.Empty() {
		length := queue.Len()
		depth++
		for i := 0; i < length; i++ {
			node := queue.Dequeue()
			if node.Left != nil && node.Left.Val == x {
				return depth + 1, node
			}
			if node.Right != nil && node.Right.Val == x {
				return depth + 1, node
			}

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
	}

	return -1, nil
}

var depthMap map[int]int
var parentMap map[int]*TreeNode

func isCousins2(root *TreeNode, x int, y int) bool {
	depthMap = make(map[int]int, 0)
	parentMap = make(map[int]*TreeNode, 0)
	nodeDFS(root, nil)

	return depthMap[x] == depthMap[y] && parentMap[x] != parentMap[y]
}

func nodeDFS(node *TreeNode, parentNode *TreeNode) {
	if node == nil {
		return
	}

	if parentNode == nil {
		depthMap[node.Val] = 1
	} else {
		depthMap[node.Val] = depthMap[parentNode.Val] + 1
	}
	parentMap[node.Val] = parentNode

	nodeDFS(node.Left, node)
	nodeDFS(node.Right, node)
}
