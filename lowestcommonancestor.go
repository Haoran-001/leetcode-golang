package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath, qPath := getPath(root, p), getPath(root, q)
	var ans *TreeNode

	for i := 0; i < len(pPath) && i < len(qPath) && pPath[i] == qPath[i]; i++ {
		ans = pPath[i]
	}

	return ans
}

func getPath(root, targetNode *TreeNode) (path []*TreeNode) {
	node := root

	for node != targetNode {
		path = append(path, node)
		if targetNode.Val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	path = append(path, node)

	return
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	tempNode := root

	for {
		if tempNode.Val > p.Val && tempNode.Val > q.Val {
			tempNode = tempNode.Left
		} else if tempNode.Val < p.Val && tempNode.Val < q.Val {
			tempNode = tempNode.Right
		} else {
			break
		}
	}

	return tempNode
}

func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor3(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor3(root.Right, p, q)
	} else {
		return root
	}
}
