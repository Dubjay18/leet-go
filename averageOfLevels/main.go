package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *
TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		var levelSum float64
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, levelSum/float64(levelSize))
	}

	return res
}