package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	fmt.Println(arr[0])
	// getting the first diagonal
	s1 := 0
	s2 := 0
	for i := 0; i < len(arr[0]); i++ {
		fmt.Println(arr[i][i])
		s1 += int(arr[i][i])
	}
	for i, j := 0, len(arr[0])-1; i < len(arr[0]) && j >= 0; i, j = i+1, j-1 {
		fmt.Println(arr[i][j])
		s2 += int(arr[i][j])
	}
	return int32(math.Abs(float64(s1-s2)))
}

func main() {
	a := []int32{11, 2, 4}
	r1 := []int32{4, 5, 6}
	r2 := []int32{10, 8, -12}
	arr := [][]int32{a, r1, r2}
	diagonalDifference(arr)
}
