package main

import (
	"log"
	"sort"
)

func main() {
	log.Printf("%+v", permuteUnique([]int{1, 1, 2, 2}))
	// log.Printf("%+v", jump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}))
}

var ans [][]int

func permuteUnique(nums []int) [][]int {
	used := make([]bool, len(nums))
	ans = [][]int{}
	sort.Ints(nums)
	solve(used, nums, []int{})
	return ans
}

func solve(used []bool, nums []int, nowAry []int) {

	if len(nowAry) == len(nums) {
		tmp := make([]int, len(nowAry))
		copy(tmp, nowAry)
		ans = append(ans, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		nowNum := nums[i]
		nowAry = append(nowAry, nowNum)
		used[i] = true
		solve(used, nums, nowAry)
		used[i] = false
		nowAry = nowAry[:len(nowAry)-1]
	}
	return
}
