package main

var head *TreeNode = &TreeNode{Val: -1, Left: nil, Right: nil}
var prev *TreeNode = nil

func convertBiNode(root *TreeNode) *TreeNode {
	convertBiNodeInorder(root)
	return head.Right
}

func convertBiNodeInorder(node *TreeNode) {
	if node == nil {
		return
	}

	convertBiNodeInorder(node.Left)

	// 第一个结点
	if prev == nil {
		prev = node
		head.Right = node
	} else {
		prev.Right = node
		prev = node
	}
	node.Left = nil
	convertBiNodeInorder(node.Right)
}
