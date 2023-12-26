package main

import "log"

func main() {
	log.Printf("removeDuplicates : %+v", removeDuplicates([]int{1, 1, 2}))
}
func removeDuplicates(nums []int) int {
	count := 0
	nowV := -101
	for _, v := range nums {
		if nowV != v {
			nums[count] = v
			count = count + 1
			nowV = v
		}
	}
	return count
}
