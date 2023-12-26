package main

import "log"

func main() {

	log.Printf("%+v", searchRange([]int{2, 2}, 2))

}
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	low, high := searchRange0(nums, target, 0, len(nums)-1)
	if high == -1 {
		return []int{-1, -1}
	}
	return []int{low, high}

}
func searchRange0(nums []int, target int, left int, right int) (int, int) {
	if left >= right {
		if nums[left] == target {
			return left, right
		} else {
			return -1, -1
		}
	} else if left < len(nums) && nums[left] > target {
		return -1, -1
	} else if right > 0 && nums[right] < target {
		return -1, -1
	}
	low := len(nums)
	high := -1
	mid := left + (right-left)/2

	if nums[mid] >= target {
		l, h := searchRange0(nums, target, left, mid-1)
		if l != -1 && l < low {
			low = l
		}
		if h > high {
			high = h
		}
	}
	if nums[mid] <= target {
		l, h := searchRange0(nums, target, mid+1, right)
		if l != -1 && l < low {
			low = l
		}
		if h > high {
			high = h
		}
	}
	if nums[mid] == target {
		if mid < low {
			low = mid
		}
		if mid > high {
			high = mid
		}
	}

	return low, high
}
