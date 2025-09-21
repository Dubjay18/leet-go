package main

import (
	"fmt"
)

func hasDuplicate(nums []int) bool {
	numMap := make(map[int]int)
	for i, v := range nums {
		if _, ok := numMap[v]; ok {
			return true
		}
		numMap[v] = i
	}
	return false
}

func main() {
	nums := []int{1, 1, 3, 4, 5}
	fmt.Println(hasDuplicate(nums))
}
