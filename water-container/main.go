package main

func maxArea(heights []int) int {
	l, r := 0, len(heights)
	max_area := 0
	for l < r {
		area := min(heights[l], heights[r])*r - l
		max_area = max(max_area, area)
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return max_area
}
func main() {
	h := []int{5, 9, 2, 1, 4}

	println(maxArea(h))
}
