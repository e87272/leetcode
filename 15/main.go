package main

import (
	"log"
	"sort"
)

func main() {

	log.Printf("threeSum : %+v", threeSum([]int{0, 0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	sumMap := [][]int{}

	sort.Ints(nums)

	for k, v := range nums {
		if v > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		j := k + 1
		for i := len(nums) - 1; i > k; i-- {
			if nums[i] < 0 {
				break
			}
			if i < len(nums)-1 && nums[i] == nums[i+1] {
				continue
			}
			for j < i {
				if v+nums[i]+nums[j] < 0 {
					j++
				} else if v+nums[i]+nums[j] == 0 {
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
