package main

var ans2 []int

func postorder(root *Node) []int {
	ans2 = make([]int, 0)
	helper2(root)

	return ans2
}

func helper2(node *Node) {
	if node == nil {
		return
	}

	for _, child := range node.Children {
		helper2(child)
	}
	ans = append(ans, node.Val)
}

func postorder2(root *Node) []int {
	stack := NodeStack{}
	ans := []int{}

	if root == nil {
		return ans
	}

	stack.Push(root)

	for !stack.Empty() {
		node := stack.Pop()
		ans = append(ans, node.Val)
		if node.Children != nil {
			for _, child := range node.Children {
				stack.Push(child)
			}
		}
	}

	reverseNodeArray(ans)

	return ans
}

func reverseNodeArray(nodeArray []int) {
	length := len(nodeArray)
	for i := 0; i < length/2; i++ {
		nodeArray[i], nodeArray[length-1-i] = nodeArray[length-1-i], nodeArray[i]
	}
}
