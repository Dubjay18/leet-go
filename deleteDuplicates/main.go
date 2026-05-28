package main

type ListNode struct {
	Val int
	Next *ListNode
}


func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head} // sentinel before the real head
	prev := dummy
	curr := head

		for curr != nil {
		if curr.Val == curr.Next.Val {
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