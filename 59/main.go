package main

import (
	"log"
)

func main() {
	log.Printf("%+v", generateMatrix(4))
}

func generateMatrix(n int) [][]int {
	matrix := [][]int{}
	for i := 0; i < n; i++ {
		temp := []int{}
		for j := 0; j < n; j++ {
			temp = append(temp, 0)
		}
		matrix = append(matrix, temp)
	}
	count := 0
	i := 0
	j := 0
	direct := 1
	for count < n*n {
		count = count + 1
		switch direct {
		case 0:
			matrix[i][j] = count
			if i-1 < 0 || matrix[i-1][j] != 0 {
				direct = 1
				j = j + 1
			} else {
				i = i - 1
			}
		case 1:
			matrix[i][j] = count
			if j+1 == n || matrix[i][j+1] != 0 {
				direct = 2
				i = i + 1
			} else {
				j = j + 1
			}
		case 2:
			matrix[i][j] = count
			if i+1 == n || matrix[i+1][j] != 0 {
				direct = 3
				j = j - 1
			} else {
				i = i + 1
			}
		case 3:
			matrix[i][j] = count
			if j-1 < 0 || matrix[i][j-1] != 0 {
				direct = 0
				i = i - 1
			} else {
				j = j - 1
			}
		}
	}

	return matrix
}
