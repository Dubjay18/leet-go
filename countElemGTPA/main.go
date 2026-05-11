package main

import "fmt"

func countResponseTimeRegressions(responseTimes []int32) int32 {
	// Write your code here
	sum := 0
	var count int32
	for i, v := range responseTimes {
		if i > 0 {
			average := sum / (i)
			if v > int32(average) {
				count++
			}
		}
		fmt.Println(sum)
		sum += int(v)
	}
	fmt.Println("i", count)
	return count
}

func main() {

	responseTimes := []int32{100, 200, 150, 300}
	countResponseTimeRegressions(responseTimes)

}
