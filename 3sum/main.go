package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}

	for i, v := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if i >= len(nums)-1 {
			break
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			tempSum := v + nums[l] + nums[r]
			if tempSum > 0 {
				r--
			} else if tempSum < 0 {
				l++
			} else {
				a := []int{v, nums[l], nums[r]}
				res = append(res, a)
				// Move both pointers
				l++
				r--

				// Skip duplicates on the left
				for l < r && nums[l] == nums[l-1] {
					l++
				}

				// Skip duplicates on the right
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}

		}
	}
	return res
}

func main() {

	ans := threeSum([]int{0, 0, 0, 0})
	fmt.Println(ans)
}
