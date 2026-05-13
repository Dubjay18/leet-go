package main

import (
	"fmt"
	"math"
)

func checkSign(n int32) string {
	if n > 0 {
		return "positive"
	} else if n < 0 {
		return "negative"
	}
	return "zero"
}

func plusMinus(arr []int32) {
	mp := make(map[string]int)
	for _,v := range arr {
		switch checkSign(v) {
		case "positive":
			mp["positive"]++
			fmt.Println(v,"p")
		case "negative":
			mp["negative"]++
			fmt.Println(v,"n")
		case "zero":
			mp["zero"]++
			fmt.Println(v,"z")
		default:
			mp["zero"]++
			fmt.Println(v,"z1")
		}
	}
	pr :=  float64(mp["positive"])/float64(len(arr))
	nr := float64(mp["negative"])/float64(len(arr))
	zr := float64(mp["zero"])/float64(len(arr))
	fmt.Println(mp["positive"],len(arr))
	fmt.Println(math.Round(pr*1e6) / 1e6)
	fmt.Println(math.Round(nr*1e6) / 1e6)
	fmt.Println(math.Round(zr*1e6) / 1e6)

}

func main() {
	arr := []int32{-4,3,-9,0,4,1}
	plusMinus(arr)
}