package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
root.Left,root.Right = root.Right,root.Left
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	
	if left == nil && right == nil {
		return root
	}
	return  nil
}