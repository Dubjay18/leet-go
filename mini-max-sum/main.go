package main

import "fmt"


func miniMaxSum(arr []int32) {
    // Write your code here
	sum := int64(0)
	m := int32(arr[0])
	mx := int32(0)
	for _,v := range arr {
		sum += int64(v)
		if m > v {
			m = v
		}
		if mx < v {
			mx = v
		}
	}
	fmt.Println(m)
	max := sum - int64(m)
	mini := sum - int64(mx)
	 fmt.Println(mini,max)
}


func main() {
	a := []int32{7,69,2,221,8974}
	miniMaxSum(a)
}