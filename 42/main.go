package main

import "log"

func main() {
	log.Printf("%+v", trap2([]int{4, 2, 3}))
}

func trap(height []int) int {
	now := 0
	trapping := 0
	tempTrap := 0
	lastTrap := 0
	for lastTrap < len(height) {
		now = height[lastTrap]
		tempTrap = 0
		for i := lastTrap; i < len(height); i++ {
			if height[i] < now {
				tempTrap = tempTrap + now - height[i]
			} else {
				lastTrap = i
				now = height[i]
				trapping = trapping + tempTrap
				tempTrap = 0
			}
		}
		if tempTrap != 0 {
			height[lastTrap] = height[lastTrap] - 1
		} else {
			lastTrap = lastTrap + 1
		}
	}
	return trapping
}

func trap2(height []int) int {
	left := 0
	right := len(height) - 1
	trapping := 0
	for left < right {
		if height[left] < height[right] {
			if height[left+1] < height[left] {
				trapping = trapping + height[left] - height[left+1]
				height[left+1] = height[left]
			}
			left = left + 1
		} else {
			if height[right-1] < height[right] {
				trapping = trapping + height[right] - height[right-1]
				height[right-1] = height[right]
			}
			right = right - 1
		}
	}
	return trapping
}
