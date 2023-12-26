package main

import "log"

func main() {
	log.Printf("%+v", searchInsert([]int{1, 3, 5, 6, 7}, 2))

}
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2
	if nums[left] >= target {
		return left
	}
	if nums[right] == target {
		return right
	}
	if nums[right] < target {
		return right + 1
	}
	for left+1 < right {
		log.Printf("left , right, mid : %+v , %+v , %+v", left, right, mid)
		if nums[mid] > target {
			right = mid
		}
		if nums[mid] < target {
			left = mid
		}
		if nums[mid] == target {
			return mid
		}
		mid = left + (right-left)/2
	}
	if target > nums[mid] {
		return mid + 1
	}
	return mid
}
