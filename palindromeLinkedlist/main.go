package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// helper: reverse a list and return new head
	var reverse func(*ListNode) *ListNode
	reverse = func(node *ListNode) *ListNode {
		var prev *ListNode
		curr := node
		for curr != nil {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		return prev
	}

	// helper: return end of first half
	endOfFirstHalf := func(head *ListNode) *ListNode {
		slow, fast := head, head
		for fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}

	// find end of first half and start of second half
	firstHalfEnd := endOfFirstHalf(head)
	secondStart := reverse(firstHalfEnd.Next)

	// compare first half and reversed second half
	p1 := head
	p2 := secondStart
	result := true
	for p2 != nil {
		if p1.Val != p2.Val {
			result = false
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// restore the list (optional)
	firstHalfEnd.Next = reverse(secondStart)

	return result
}
