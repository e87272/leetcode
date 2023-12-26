package main

import (
	"log"
	"sort"
)

func main() {

	log.Printf("threeSumClosest : %+v", fourSum([]int{2, 2, 2, 2, 2}, 8))
}

func fourSum(nums []int, target int) [][]int {
	sumMap := [][]int{}
	sort.Ints(nums)
	for k, v := range nums {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		newTarget := target - v
		threeMap := threeSum(nums[k+1:], newTarget)
		for _, threeNum := range threeMap {
			sumMap = append(sumMap, append([]int{v}, threeNum...))
		}
	}
	return sumMap
}

func threeSum(nums []int, target int) [][]int {
	sumMap := [][]int{}
	for k, v := range nums {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		j := k + 1
		for i := len(nums) - 1; i > k; i-- {
			if i < len(nums)-1 && nums[i] == nums[i+1] {
				continue
			}
			for j < i {
				if v+nums[i]+nums[j] < target {
					j++
				} else if v+nums[i]+nums[j] == target {
					sumMap = append(sumMap, []int{v, nums[j], nums[i]})
					break
				} else {
					break
				}
			}
		}
	}
	return sumMap
}
