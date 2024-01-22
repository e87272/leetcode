package main

import (
	"log"
)

func main() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	log.Printf("%+v", matrix)
	// log.Printf("%+v", jump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}))
}

func rotate(matrix [][]int) {
	// sideLen := (len(matrix)-1) / 2
	for x := 0; x < len(matrix)/2; x++ {
		for y := x; y < len(matrix[x])-1-x; y++ {
			nowTemp := matrix[x][y]
			tempX := y
			tempY := -x + len(matrix) - 1
			nextTemp := matrix[tempX][tempY]
			matrix[tempX][tempY] = nowTemp
			for tempX != x || tempY != y {
				nowTemp = nextTemp
				nowX := tempX
				tempX = tempY
				tempY = -nowX + len(matrix) - 1
				nextTemp = matrix[tempX][tempY]
				matrix[tempX][tempY] = nowTemp
			}
		}
	}
}
