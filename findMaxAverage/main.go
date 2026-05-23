package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
  left := 0
  right := k
  maxSum := 0
  sum := 0
    for i:= 0; i < k; i++ {
      sum += nums[i]
    }
    maxSum = sum

  for right < len(nums) {
    sum = sum - nums[left] + nums[right]
      if sum > maxSum {
        maxSum = sum
      }
      left++
      right++
  }
  return float64(maxSum)/float64(k)
}

func main() {
  arr := []int{1,12,-5,-6,50,3}
  k := 4
  res := findMaxAverage(arr,k)
  fmt.Println(res)
}
