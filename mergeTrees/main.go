package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var res *TreeNode
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		res = &TreeNode{Val: root2.Val}
		res.Left = mergeTrees(nil, root2.Left)
		res.Right = mergeTrees(nil, root2.Right)
		return res
	}
	if root2 == nil {
		res = &TreeNode{Val: root1.Val}
		res.Left = mergeTrees(root1.Left, nil)
		res.Right = mergeTrees(root1.Right, nil)
		return res
	}

	res = &TreeNode{Val: root1.Val + root2.Val}
	res.Left = mergeTrees(root1.Left, root2.Left)
	res.Right = mergeTrees(root1.Right, root2.Right)

	return res
}