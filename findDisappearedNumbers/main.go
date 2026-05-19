package main





func findDisappearedNumbers(nums []int) []int {
    seen := make(map[int]bool)

    for _, num := range nums {
        seen[num] = true
    }

    results := []int{}
    for i := 1; i <= len(nums); i++ {
        if !seen[i] {
            results = append(results, i)
        }
    }

    return results
}

func main() {
	arr := []int{4,3,2,7,8,2,3,1}
	findDisappearedNumbers(arr)
}