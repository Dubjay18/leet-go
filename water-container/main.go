package main

import "math"

func maxArea(heights []int) int {
	var max_area int
	n := len(heights)
	for p1 := 0; p1 < n; p1++ {
		for p2 := p1 + 1; p2 < n; p2++ {
			length := math.Min(float64(heights[p1]), float64(heights[p2]))
			width := p2 - p1
			area := length * float64(width)
			max_area = int(math.Max(float64(max_area), float64(area)))
		}
	}
	return max_area
}
func main() {
	h := []int{5, 9, 2, 1, 4}

	println(maxArea(h))
}
