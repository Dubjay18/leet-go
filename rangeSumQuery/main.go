package main

type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {
	n := NumArray{}
	n.prefix = []int{}
	cur := 0
	for _, v := range nums {
		cur += v
		n.prefix = append(n.prefix, cur)
	}
	return n
}

func (this *NumArray) SumRange(left int, right int) int {
	rightSum := this.prefix[right]
	leftSum := 0
	if left > 0 {
		leftSum = this.prefix[left-1]
	}
	return rightSum - leftSum
}
