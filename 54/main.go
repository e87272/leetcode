package main

import (
	"log"
)

func main() {
	log.Printf("%+v", spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
func spiralOrder(matrix [][]int) []int {
	orderNums := []int{}
	col := 0
	row := 0
	flag := 1 //0:up 1:right 2:down 3:left
	for i := 0; i < len(matrix)*len(matrix[0]); i++ {
		orderNums = append(orderNums, matrix[row][col])
		matrix[row][col] = -101
		switch flag {
		case 0:
			if row-1 == -1 || matrix[row-1][col] == -101 {
				flag = 1
				col = col + 1
			} else {
				row = row - 1
			}
		case 1:
			if col+1 == len(matrix[0]) || matrix[row][col+1] == -101 {
				flag = 2
				row = row + 1
			} else {
				col = col + 1
			}
		case 2:
			if row+1 == len(matrix) || matrix[row+1][col] == -101 {
				flag = 3
				col = col - 1
			} else {
				row = row + 1
			}
		case 3:
			if col-1 == -1 || matrix[row][col-1] == -101 {
				flag = 0
				row = row - 1
			} else {
				col = col - 1
			}
		}
	}
	return orderNums
}
