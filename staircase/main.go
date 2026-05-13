package main

import (
	"fmt"
	"strings"
)



func staircase(n int32) {
	for v := range n {
		nm := v+1
		spaces := strings.Repeat(" ", int(n-nm))
		hashes := strings.Repeat("#",int(nm))
		fmt.Println(spaces+hashes)
	}	
}

func main() {
	staircase(6)
}
