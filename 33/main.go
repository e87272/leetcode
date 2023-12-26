package main

import "log"

func main() {

	log.Printf("search : %+v", search([]int{5, 1, 2, 3, 4}, 1))
}
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for right >= left {
		mid := left + (right-left)/2
		log.Printf("%+v %+v %+v", left, mid, right)
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
