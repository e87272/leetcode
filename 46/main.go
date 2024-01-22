package main

import "log"

func main() {
	log.Printf("%+v", permute([]int{1, 2, 3}))
	// log.Printf("%+v", jump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}))
}
func permute(nums []int) [][]int {
	permute := [][]int{}
	index := 0
	permute = append(permute, nums)
	for index < len(nums)-1 {
		for _, ary := range permute {
			for i := index + 1; i < len(nums); i++ {
				newAry := make([]int, len(ary))
				copy(newAry, ary)
				newAry[i] = ary[index]
				newAry[index] = ary[i]
				permute = append(permute, newAry)
			}
		}
		index = index + 1
	}
	return permute
}
