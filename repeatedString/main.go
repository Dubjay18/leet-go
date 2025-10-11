package main

import "fmt"

func repeatedString(s string, n int64) int64 {
	// Write your code here

	var count int64
	var fc, sc int64
	var reminderLen int64
	for _, v := range s {
		if v == 'a' {
			fc++
		}
	}
	reminderLen = n % int64(len(s))
	for i, v := range s {
		if i >= int(reminderLen) {
			break
		}
		if v == 'a' {
			sc++
		}
	}
	count = fc + sc
	return count
}

func main() {
	c := repeatedString("aba", 10)
	fmt.Println(c)
}
