package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// removeElements removes all nodes with value `val` and returns the new head.
// Uses a dummy sentinel to simplify head removals and updates pointers
// in a single pass without special-casing the head.
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head} // sentinel before the real head
	prev := dummy
	curr := head

	for curr != nil {
		if curr.Val == val {
			// skip curr by linking prev to curr.Next
			prev.Next = curr.Next
			// advance curr to the node after the removed one
			curr = prev.Next
		} else {
			// keep curr, move both pointers forward
			prev = curr
			curr = curr.Next
		}
	}

	return dummy.Next
}

func main() {
	// Example: remove 7 from [7,7,7,7]
	a := &ListNode{Val: 7}
	a.Next = &ListNode{Val: 7}
	a.Next.Next = &ListNode{Val: 7}
	a.Next.Next.Next = &ListNode{Val: 7}

	removed := removeElements(a, 7)
	for node := removed; node != nil; node = node.Next {
		println(node.Val)
	}
}
