package main

import (
	"log"
	"sort"
)

func main() {
	log.Printf("%+v", merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}
func merge(intervals [][]int) [][]int {
	mergeIntervals := [][]int{}

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
