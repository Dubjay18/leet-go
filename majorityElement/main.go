package main

func majorityElement(nums []int) int {
	count := make(map[int]int)
	majority := 0
	for _, v := range nums {
		count[v]++
	}

	for k, v := range count{
		if count[majority] < v {
			majority = k
		}
	}
	return majority
 }

func main() {

}