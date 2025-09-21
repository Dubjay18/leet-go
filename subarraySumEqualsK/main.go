package main

import "fmt"

func subarraySum(nums []int, k int) int {
	//result count
	var c int
	// prefix sum
	ps := map[int]int{0: 1}
	// current sum
	cSum := 0
	for _, v := range nums {
		cSum += v
		if j, ok := ps[cSum-k]; ok {
			c += j
		}
		ps[cSum] = 1 + ps[cSum]

	}
	return c
}

func main() {
	// Test the subarraySum function
	a := []int{1, 2, 3, 3}
	result := subarraySum(a, 3)
	fmt.Println(result)
}
