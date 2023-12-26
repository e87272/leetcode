package main

import "log"

func main() {
	log.Printf("%+v", jump([]int{2, 9, 6, 5, 7, 0, 7, 2, 7, 9, 3, 2, 2, 5, 7, 8, 1, 6, 6, 6, 3, 5, 2, 2, 6, 3}))
	// log.Printf("%+v", jump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}))
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	times := jumpTo(nums, 0, 0)
	return times
}

func jumpTo(nums []int, nowIndex int, times int) int {
	maxIndex := 0
	for i := 1; i <= nums[nowIndex]; i++ {
		if nowIndex+i >= len(nums)-1 {
			return times + 1
		}
		if nowIndex+i+nums[nowIndex+i] > maxIndex+nums[maxIndex] {
			maxIndex = nowIndex + i
		}
	}
	return jumpTo(nums, maxIndex, times+1)
}
