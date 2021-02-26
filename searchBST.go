package main

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	}

	if val > root.Val {
		return searchBST(root.Right, val)
	}

	return nil
}

func searchBST2(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if val < root.Val {
			root = root.Left
		} else if val > root.Val {
			root = root.Right
		}
	}

	return root
}
