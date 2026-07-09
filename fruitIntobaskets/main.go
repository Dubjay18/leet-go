package main


func totalFruit(fruits []int) int {
	count := make(map[int]int)
	l, total, result := 0, 0, 0

	for r := range(len(fruits)) {
		count[fruits[r]] += 1
		total += 1

		for len(count) > 2 {
			f := fruits[l]
			count[f] -= 1
			total--
			l++

			if count[f] == 0 {
				delete(count,f)
			}
		}

		result = max(result,total)
	}
	return  result
}