package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// First pass: calculate left products
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Second pass: calculate right products and multiply with left products
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 1}))

}
