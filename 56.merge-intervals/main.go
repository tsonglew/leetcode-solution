package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	result := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= result[len(result)-1].End {
			if result[len(result)-1].End < intervals[i].End {
				result[len(result)-1].End = intervals[i].End
			}
		} else {
			result = append(result, Interval{intervals[i].Start, intervals[i].End})
		}
	}
	return result
}

func main() {
	fmt.Println(merge([]Interval{
		Interval{1, 4},
		Interval{0, 4},
	}))
}
