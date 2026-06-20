package main

import "fmt"


func minSubArrayLen(target int, nums []int) int {
	res := 0
	bucketSum := 0
	l := 0
	for r := range(len(nums)) {
		bucketSum += nums[r]
		for  bucketSum >= target {
			length := r - l + 1
			if length < res {
				res = length
			} else if length > res && res == 0 {
				res = length
			} 
			bucketSum -= nums[l]
			l++
		}
	}
	return res
}

func main() {
	result := minSubArrayLen(11,[]int{1,1,1,1,1,1,1,1})
	fmt.Println(result)
}