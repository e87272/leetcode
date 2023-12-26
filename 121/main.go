package main

import (
	"log"
)

func main() {
	log.Printf("%+v", maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
func maxProfit(prices []int) int {
	min := prices[0]
	max := prices[0]
	nowMaxProfit := 0
	for _, v := range prices {
		if v < min {
			max = 0
			min = v
			if max-min > nowMaxProfit {
				nowMaxProfit = max - min
			}
		}
		if v > max {
			max = v
			if max-min > nowMaxProfit {
				nowMaxProfit = max - min
			}
		}
	}

	return nowMaxProfit
}
