package main

import (
	"fmt"
	"slices"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)
	res := nums[0] + nums[1] + nums[2]

	for i, v := range nums {
		if i >= len(nums)-1 {
			break
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			tempSum := v + nums[l] + nums[r]
			diff := abs(target - tempSum)
			resDiff := abs(target - res)
			if diff <= resDiff {
				res = tempSum
			}
			if tempSum < target {
				l++
			} else if tempSum > target {
				r--
			} else {
				return target
			}
		}
	}
	return res

}

func main() {
	res := threeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2)
	fmt.Println(res, "ans1")
	// res2 := threeSumClosest([]int{-1,2,1,-4},1)
	// fmt.Println(res2,"ans2")
}
