package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1

	for l <= r {
		mid := l + (r-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if l == len(letters) {
		return letters[0]
	}

	return letters[l]
}

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c')))
	fmt.Println(string(nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'z')))

}
