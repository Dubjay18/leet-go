package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSameTree checks whether two trees are identical.
func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

// isSubtree returns true if subRoot is a subtree of root.
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func main() {
	// Example:
	//   root:    3
	//          / \
	//         4   5
	//        / \
	//       1   2
	//   subRoot:
	//         4
	//        / \
	//       1   2

	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}

	sub := &TreeNode{Val: 4}
	sub.Left = &TreeNode{Val: 1}
	sub.Right = &TreeNode{Val: 2}

	fmt.Println(isSubtree(root, sub)) // expected: true
}
