package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(nums []any) *TreeNode {
	if len(nums) == 0 || nums[0] == nil {
		return nil
	}
	root := &TreeNode{Val: nums[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(nums) {
		node := queue[0]
		queue = queue[1:]
		// left child
		if nums[i] != nil {
			node.Left = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i >= len(nums) {
			break
		}

		// right child
		if nums[i] != nil {
			node.Right = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func main() {
	vals := []any{3, 9, 20, nil, nil, 15, 7}
	root := buildTree(vals)

	depth := maxDepth(root)
	fmt.Println(depth) // should print 3
}
