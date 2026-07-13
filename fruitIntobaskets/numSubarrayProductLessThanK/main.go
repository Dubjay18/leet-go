package main



func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	count := 0
	product := 1
	l := 0


	for r := range(nums) {
		product *= nums[r]
		for product >= k {
			product /= nums[l]
			l++
		}
		count += r - l + 1
	}
	return count
}