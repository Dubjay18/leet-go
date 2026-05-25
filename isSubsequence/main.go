package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	base := 0

	for i := 0; i < len(t); i++ {
		if s[base] == t[i] {
			if base == len(s)-1 {
				return true
			} else {
				base++
			}
		}
	}
	return false
}

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
