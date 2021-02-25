package main

var array1, array2 []int

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	array1 = make([]int, 0)
	array2 = make([]int, 0)

	leafSimilarDFS(root1, &array1)
	leafSimilarDFS(root2, &array2)

	return equalOfTwoIntArray(array1, array2)
}

func leafSimilarDFS(node *TreeNode, array *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*array = append(*array, node.Val)
		return
	}

	leafSimilarDFS(node.Left, array)
	leafSimilarDFS(node.Right, array)
}
