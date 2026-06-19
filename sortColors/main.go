package main


// bucket sort
func sortColors(nums []int) {
	bucket := map[int]int{}

	for _,v := range nums {
		bucket[v]++
	}
	for i :=0 ;i < bucket[0]; i++ {
		nums[i] = 0
	}

	for i := bucket[0]; i < bucket[1]+bucket[0]; i++ {
		nums[i] = 1
	}

	for i := bucket[1]+bucket[0]; i < bucket[2]+bucket[1]+bucket[0]; i++ {
		nums[i] = 2
	}
}

func main() {
	sortColors([]int{2,0,2,1,1,0})
}