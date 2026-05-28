package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	seen := make(map[*ListNode]bool)
	curr := head
	for curr != nil  {
		if !seen[curr] {
			seen[curr] = true
			curr = curr.Next
			continue
		}
		return true
	}
	return false
}