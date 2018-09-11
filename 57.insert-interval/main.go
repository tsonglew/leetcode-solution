package main

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) < 1 {
		return []Interval{newInterval}
	}
	var mn, mx int
	for mn, mx = 0, len(intervals)-1; mn <= mx; {
		mid := (mx + mn) / 2
		if intervals[mid].Start > newInterval.Start {
			mx = mid - 1
		} else if intervals[mid].Start < newInterval.Start {
			mn = mid + 1
		} else {
			mn, mx = mid, mid
			break
		}
	}
	if mn > mx {
		mn, mx = mx, mn
	}
	if mn < 0 {
		intervals = append([]Interval{newInterval}, intervals...)
		mn = 0
	} else if mx > len(intervals)-1 {
		intervals = append(intervals, newInterval)
		mx = len(intervals) - 1
	} else {
		intervals = append(intervals[0:mn+1], append([]Interval{newInterval}, intervals[mn+1:len(intervals)]...)...)
	}
	var i int
	if intervals[mn].End < intervals[mn+1].Start {
		mn++
	}
	for i = mn; i < len(intervals); i++ {
		if intervals[i].Start <= intervals[mn].End {
			intervals[mn].End = max(intervals[mn].End, intervals[i].End)
		} else {
			break
		}
	}
	result := append(intervals[0:mn+1], intervals[i:len(intervals)]...)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
