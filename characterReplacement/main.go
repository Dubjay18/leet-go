package main

import "fmt"

func characterReplacement(s string, k int) int {
	res := 0
	l := 0
	count := make(map[byte]int, 26)
	maxFreq := 0
	for r := range len(s) {
		count[s[r]]++
		maxFreq = max(maxFreq,count[s[r]])
		for ((r -l) + 1 - maxFreq) > k {
			count[s[l]]--
			l++
		}
		res = max(res,  (r -l) + 1)
	}
	return  res
}

func main() {
	result := characterReplacement("AABABBA",1)
	fmt.Println(result)
}