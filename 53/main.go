package main

import (
	"log"
)

func main() {
	v := maxSubArray([]int{5, 4, -1, 7, 8})
	log.Printf("%+v", v)
}
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	sum := nums[0]
	left := 0
	for i := 1; i < len(nums); i++ {
		if sum+nums[i]-nums[left] > sum+nums[i] {
			sum = sum + nums[i] - nums[left]
			left = left + 1
		} else {
			sum = sum + nums[i]
		}
		if sum < 0 {
			sum = nums[i]
			left = i
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
