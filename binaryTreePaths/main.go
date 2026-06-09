package main

import (
	"fmt"
)

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

// func addPath(path string,toAdd string) {

// }
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}

	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		if path == "" {
			path = fmt.Sprintf("%v", node.Val)
		} else {
			path = fmt.Sprintf("%s->%v", path, node.Val)
		}

		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}

		dfs(node.Left, path)
		dfs(node.Right, path)
	}

	dfs(root, "")
	return res
}

func main() {
	vals := []any{1, 2, 3, nil, 5}
	root := buildTree(vals)

	paths := binaryTreePaths(root)
	fmt.Println(paths)
}
