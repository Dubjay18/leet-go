package main


func moveZeroes(nums []int) {
	j := 0
	for _, v  := range nums {
		if v != 0 {
			nums[j] = v
		j++
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	arr := []int{0,1,0,3,12}
	moveZeroes(arr)
}