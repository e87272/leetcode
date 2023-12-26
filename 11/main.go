package main

import (
	"log"
)

func main() {
	area := maxArea([]int{0, 2})
	log.Printf("area : %+v", area)
}

func maxArea(height []int) int {

	var max int

	for i := 0; i < len(height); i++ {
		if height[i]*(len(height)-1) < max || height[i] == 0 {
			continue
		}
		for j := int(max/height[i]) + 1; j < len(height); j++ {
			if height[i] < height[j] {
				if height[i]*(j-i) > max {
					max = height[i] * (j - i)
				}
			} else {
				if height[j]*(j-i) > max {
					max = height[j] * (j - i)
				}
			}
		}
	}

	return max
}
