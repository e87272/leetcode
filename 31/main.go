package main

import (
	"log"
	"sort"
)

func main() {
	nums := []int{5, 1, 1}
	nextPermutation(nums)
	log.Printf("%+v", nums)
}

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	increment := nums[0] < nums[1]
	changeIndex := 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] != increment {
			changeIndex = i
			increment = nums[i] < nums[i+1]
		}
	}
	if increment {
		temp := nums[len(nums)-1]
		nums[len(nums)-1] = nums[len(nums)-2]
		nums[len(nums)-2] = temp
	} else if changeIndex == 0 {
		sort.Ints(nums[changeIndex:])
	} else {
		for i := len(nums) - 1; i >= changeIndex; i-- {
			if nums[i] > nums[changeIndex-1] {
				temp := nums[changeIndex-1]
				nums[changeIndex-1] = nums[i]
				nums[i] = temp
				sort.Ints(nums[changeIndex:])
				break
			}
		}
	}
}
