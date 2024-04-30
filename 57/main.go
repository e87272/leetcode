package main

import (
	"log"
	"sort"
)

func main() {
	log.Printf("%+v", insert2([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}

func insert(intervals [][]int, newInterval []int) [][]int {

	mergeIntervals := [][]int{}
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	nowMax := intervals[0][1]
	mergeIntervals = append(mergeIntervals, []int{intervals[0][0], intervals[0][1]})
	for i := 1; i < len(intervals); i++ {
		if nowMax < intervals[i][0] {
			mergeIntervals = append(mergeIntervals, []int{intervals[i][0], intervals[i][1]})
		}
		if intervals[i][1] > nowMax {
			mergeIntervals[len(mergeIntervals)-1][1] = intervals[i][1]
			nowMax = intervals[i][1]
		}
	}

	return mergeIntervals
}

func insert2(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append([][]int{newInterval}, intervals...)
	}
	if intervals[0][0] > newInterval[1] {
		return append([][]int{newInterval}, intervals...)
	}
	if intervals[len(intervals)-1][1] < newInterval[0] {
		return append(intervals, newInterval)
	}
	nowMax := newInterval[1]
	mergeFlag := false
	start := 0
	var i int
	for i = 0; i < len(intervals); i++ {
		if !mergeFlag && intervals[i][1] >= newInterval[0] {
			if intervals[i][0] > newInterval[1] {
				intervals = append(intervals[:i], append([][]int{newInterval}, intervals[i:]...)...)
				break
			}
			if intervals[i][0] > newInterval[0] {
				intervals[i][0] = newInterval[0]
			}
			mergeFlag = true
			start = i
		}
		if mergeFlag {
			if nowMax < intervals[i][0] {
				intervals[start][1] = nowMax
				intervals = append(intervals[:start+1], intervals[i:]...)
				break
			} else if nowMax < intervals[i][1] {
				intervals[start][1] = intervals[i][1]
				intervals = append(intervals[:start+1], intervals[i+1:]...)
				break
			}
		}
	}

	if mergeFlag && nowMax > intervals[start][1] {
		intervals[start][1] = nowMax
		if nowMax > intervals[len(intervals)-1][1] {
			intervals = intervals[:start+1]
		}
	}

	return intervals
}
