package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		// Sort characters in the string to create a key
		chars := strings.Split(str, "")
		sort.Strings(chars)
		key := strings.Join(chars, "")

		// Group strings by their sorted character key
		groups[key] = append(groups[key], str)
	}

	// Convert map values to result slice
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
