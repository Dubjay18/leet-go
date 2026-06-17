package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func findDuplicates(nums []int) []int {
	var res []int
	for _,v := range nums {
		index := abs(v) -1

		if nums[index] < 0{
			res = append(res, abs(v))
		} else {
			nums[index] = -nums[index]
		}
	}
	return res  
}

func main() {
	arr := []int{4,3,2,7,8,2,3,1}
	r := findDuplicates(arr)
	fmt.Println(r)
}