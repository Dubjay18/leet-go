package main

import (
	"fmt"
	"sort"
)




type Interval struct {
	Start, End int
}


func CanAttendMeetings(intervals []*Interval) bool {
	// sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	for i:= 1; i < len(intervals); i++ {
		i1 := intervals[i-1]
		i2 := intervals[i]

		if i1.End > i2.Start {
			return false
		}

	}

	return true
}



func main() {
	intervals := []*Interval{
		{Start: 0, End: 30},
		{Start: 5, End: 10},
		{Start: 15, End: 20},
	}

	fmt.Println(CanAttendMeetings(intervals))
}