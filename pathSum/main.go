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

		if nums[i] != nil {
			node.Left = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i >= len(nums) {
			break
		}

		if nums[i] != nil {
			node.Right = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// If it's a leaf, check whether the remaining target equals the node value
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	// Otherwise, subtract current node value and recurse
	remaining := targetSum - root.Val
	return hasPathSum(root.Left, remaining) || hasPathSum(root.Right, remaining)
}

func main() {
	vals := []any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, nil, nil, 1}
	root := buildTree(vals)
	targetSum := 22

	result := hasPathSum(root, targetSum)
	fmt.Println(result) // should print true
}
