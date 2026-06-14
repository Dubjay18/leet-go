package main


func construct2DArray(original []int, m int, n int) [][]int {
    if len(original) != n * m {
		return [][]int{}
	}

	res := make([][]int, m)

	for r := range m {
		start := r * n
		end := start + n
		res = append(res, original[start:end])
	}
	return  res
}