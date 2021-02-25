package main

var ans []int = []int{}

func preorder(root *Node) []int {
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)

	for i := 0; i < len(root.Children); i++ {
		preorder(root.Children[i])
	}

	return ans
}

func preorder2(root *Node) []int {
	return nil
}
