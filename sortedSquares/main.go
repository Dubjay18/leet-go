package main

func sortedSquares(nums []int) []int {
	for i, v := range nums {
		nums[i] = v * v
	}
	// sort array in descending order in O(nlogn)
	n := len(nums)

	for i := 0; i < n-1; i++ {
		swapped := false

		// two pointers: j and j+1
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}

		// stop if already sorted
		if !swapped {
			break
		}
	}
	return nums
}
