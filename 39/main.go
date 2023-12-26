package main

import "log"

func main() {
	log.Printf("%+v", combinationSum([]int{7, 3, 2}, 18))
}

func combinationSum(candidates []int, target int) [][]int {
	return sum(candidates, target, []int{})
}

func sum(candidates []int, target int, result []int) [][]int {
	combinationSum := [][]int{}
	for k, v := range candidates {
		if target-v == 0 {
			resultCpy := make([]int, len(result)+1)
			copy(resultCpy, append(result, v))
			combinationSum = append(combinationSum, resultCpy)
		}
		if target-v > 0 {
			combinationSum = append(combinationSum, sum(candidates[k:], target-v, append(result, v))...)
		}
	}
	return combinationSum
}
