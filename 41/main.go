package main

import (
	"log"
	"sort"
)

func main() {
	log.Printf("%+v", firstMissingPositive([]int{1, 2, 0}))
}

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	smallest := 1
	for _, v := range nums {
		if v > 0 {
			if smallest == v {
				smallest++
			} else if smallest-1 == v {
				continue
			} else {
				break
			}
		}
	}
	return smallest
}
