package main

import (
	"log"
	"sort"
)

func main() {
	log.Printf("%+v", combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return sum2(candidates, target, []int{})
}

func sum2(candidates []int, target int, result []int) [][]int {
	combinationSum := [][]int{}
	for k, v := range candidates {
		if v > target {
			break
		}
		if k > 0 && v == candidates[k-1] {
			continue
		}
		if target-v == 0 {
			resultCpy := make([]int, len(result)+1)
			copy(resultCpy, append(result, v))
			combinationSum = append(combinationSum, resultCpy)
		}
		if target-v > 0 {
			combinationSum = append(combinationSum, sum2(candidates[k+1:], target-v, append(result, v))...)
		}
	}
	return combinationSum
}
