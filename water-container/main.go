package main

import "fmt"

func maxArea(height []int) int {
	l ,r := 0, len(height) - 1
	currMax := 0
	for l < r {
		a := (r-l) * min(height[l],height[r])
		if a > currMax {
			currMax = a
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return currMax
}



func main() {
c := maxArea([]int{1,8,6,2,5,4,8,3,7})
fmt.Println(c)
}