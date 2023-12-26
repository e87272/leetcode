package main

import (
	"log"
	"math"
	"sort"
)

func main() {

	log.Printf("threeSumClosest : %+v", threeSumClosest([]int{2, 5, 6, 7}, 16))
}

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	min := math.MinInt64
	closest := 0

	preDiff := 0.00
	for k, v := range nums {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		if k > 0 && math.Abs(float64(nums[k-1]+nums[k]+nums[len(nums)-1]-target)) > preDiff {
			continue
		} else {
			preDiff = math.Abs(float64(nums[0] + nums[1] + nums[len(nums)-1] - target))
		}
		for i := len(nums) - 1; i > k; i-- {
			if i < len(nums)-1 && nums[i] == nums[i+1] {
				continue
			}
			j := k + 1
			for j < i {
				if int(math.Abs(float64(v+nums[i]+nums[j]-target))) < min {
					if v+nums[i]+nums[j] == target {
						return target
					} else {
						min = int(math.Abs(float64(v + nums[i] + nums[j] - target)))
						closest = v + nums[i] + nums[j]
					}
				}
				j++
			}
		}
	}
	return closest
}
