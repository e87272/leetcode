package main

import (
	"log"
	"sort"
)

func main() {
	log.Printf("removeDuplicates : %+v", removeElement([]int{1, 1, 2}, 1))
}
func removeElement(nums []int, val int) int {
	sort.Ints(nums)
	removeCount := 0
	for _, v := range nums {
		if val != v {
			nums[removeCount] = v
			removeCount = removeCount + 1
		}
	}
	return removeCount
}
