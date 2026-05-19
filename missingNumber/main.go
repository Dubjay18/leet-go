package main


func missingNumber(nums []int) int {
	  seen := make(map[int]bool)

    for _, num := range nums {
        seen[num] = true
    }

    result := 0
    for i := 1; i <= len(nums); i++ {
        if !seen[i] {
            result = i
        }
    }

    return result
}

func main() {

}