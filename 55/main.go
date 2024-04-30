package main

import (
	"log"
)

func main() {
	log.Printf("%+v", canJump([]int{3, 2, 1, 0, 4}))
}
func canJump(nums []int) bool {
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if maxIndex < i {
			return false
		}
		if i+nums[i] > maxIndex {
			maxIndex = i + nums[i]
		}
	}
	return true
}
