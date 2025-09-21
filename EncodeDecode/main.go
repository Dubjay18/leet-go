package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encodedStrs := make([]string, len(strs))
	for _, v := range strs {
		var en []string
		n := strconv.Itoa(len(v))
		en = append(en, n)
		en = append(en, "#")
		en = append(en, v)
		encodedStrs = append(encodedStrs, strings.Join(en, ""))
	}
	return strings.Join(encodedStrs, "")
}

func (s *Solution) Decode(str string) []string {
	var result []string
	i := 0

	for i < len(str) {
		// Find the position of "#"
		delimPos := strings.Index(str[i:], "#")
		if delimPos == -1 {
			break
		}

		delimPos += i // Adjust position relative to the start

		// Extract the length value
		lenStr := str[i:delimPos]
		length, err := strconv.Atoi(lenStr)
		if err != nil {
			break
		}

		// Extract the actual string
		start := delimPos + 1
		end := start + length
		if end > len(str) {
			break
		}

		decodedStr := str[start:end]
		result = append(result, decodedStr)

		// Move to the next encoded string
		i = end
	}

	return result
}

func main() {
	so := Solution{}
	fmt.Println(so.Encode([]string{"leet", "co#de", "", "😀"}))
}
