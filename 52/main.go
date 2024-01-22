package main

import (
	"log"
)

func main() {
	v := totalNQueens(5)
	log.Printf("%+v", v)
}

var result [][]string

func totalNQueens(n int) int {

	queenMap := map[[2]int]bool{}

	return addQueen(n, 0, queenMap, 0)
}

func addQueen(n int, row int, queenMap map[[2]int]bool, count int) int {

	if len(queenMap) == n {
		// log.Printf("queenMap : %+v", queenMap)
		return count + 1
	}

	for i := 0; i < n; i++ {
		// log.Printf("%+v", queenMap)
		if checkQueen(queenMap, row, i) {
			queenMap[[2]int{row, i}] = true
			count = addQueen(n, row+1, queenMap, count)
			delete(queenMap, [2]int{row, i})
		}
	}

	return count
}

func checkQueen(queenMap map[[2]int]bool, row int, col int) bool {

	for k, _ := range queenMap {
		if k[1] == col || row-k[0] == col-k[1] || row-k[0] == k[1]-col {
			return false
		}
	}

	return true
}
