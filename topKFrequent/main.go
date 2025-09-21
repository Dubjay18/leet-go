package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	res := make([]int, k)
	for _, v := range nums {
		m[v] += 1
	}

	type kv struct {
		Key, Value int
	}
	var pairs []kv
	for k, v := range m {
		pairs = append(pairs, kv{k, v})
	}

	// Sort by value
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	for i := 0; i < k; i++ {
		res[i] = pairs[i].Key
	}

	return res
}
