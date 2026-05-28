package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l1 := list1
	l2 := list2
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return dummy.Next
}

func main() {
	// Example: merge [1,2,4] and [1,3,4]
	a := &ListNode{Val: 1}
	a.Next = &ListNode{Val: 2}
	a.Next.Next = &ListNode{Val: 4}

	b := &ListNode{Val: 1}
	b.Next = &ListNode{Val: 3}
	b.Next.Next = &ListNode{Val: 4}

	merged := mergeTwoLists(a, b)
	for node := merged; node != nil; node = node.Next {
		println(node.Val)
	}
}
