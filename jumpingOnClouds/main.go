package main

func jumpingOnClouds(c []int32) int32 {
	var jumps int32
	for i := 0; i < len(c)-1; {
		if i+2 < len(c) && c[i+2] == 0 {
			i += 2
		} else {
			i++
		}
		jumps++
	}
	return jumps
}

func main() {
	p := []int32{0, 0, 1, 0, 0, 1, 0}
	jumpingOnClouds(p)
}
