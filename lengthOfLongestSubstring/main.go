package main

import "fmt"


func lengthOfLongestSubstring(s string) int {
	bucket := map[byte]bool{}
	left := 0
	result := 0
	for r := range(len(s)) {
		fmt.Println(r,left,s[r],s[left])
		for bucket[s[r]]  {
			bucket[s[left]] = false
			left++
		}
		bucket[s[r]] = true
		result = max(result,r - left + 1)
	} 
	return result

}


func main() {
	s :=  "abcabcbb"
	r := lengthOfLongestSubstring(s)
	fmt.Println(r)
}