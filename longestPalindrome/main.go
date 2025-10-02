package longestpalindrome

import "fmt"

func longestpalindrome(s string) string {
	result := ""
	maxLen := 0

	for i := range s {
		l, r := i, i
		// check odd
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r-l)+1 > maxLen {
				maxLen = (r - l) + 1
				result = s[l : r+1]
			}
			l--
			r++
		}
		// check even
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r-l)+1 > maxLen {
				maxLen = (r - l) + 1
				result = s[l : r+1]
			}
			l--
			r++
		}
	}

	return result
}

func main() {
	r := longestpalindrome("babad")
	fmt.Println(r)
}
