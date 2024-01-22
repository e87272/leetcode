package main

import (
	"log"
)

func main() {
	v := solveNQueens(5)
	log.Printf("%+v", v)
}

var result [][]string

func solveNQueens(n int) [][]string {

	if n == 1 {
		return [][]string{{"Q"}}
	}
	result = [][]string{}
	for i := 0; i < n; i++ {
		queenAry := [][]int{}
		for j := 0; j < n; j++ {
			queenAry = append(queenAry, make([]int, n))
		}
		log.Printf("i : %+v", i)
		addQueen(i, 0, queenAry)
	}

	return result
}

func addQueen(i int, j int, queenAry [][]int) bool {

	if queenAry[i][j] != 0 {
		return false
	}
	road := [][]int{}
	for k := 0; k < len(queenAry); k++ {
		if queenAry[i][k] == 1 {
			return false
		} else if queenAry[i][k] == 0 {
			road = append(road, []int{i, k})
		}
		if queenAry[k][j] == 1 {
			return false
		} else if queenAry[k][j] == 0 {
			road = append(road, []int{k, j})
		}
	}
	k := 0
	for i-k >= 0 && j-k >= 0 {
		if queenAry[i-k][j-k] == 1 {
			return false
		} else if queenAry[i-k][j-k] == 0 {
			road = append(road, []int{i - k, j - k})
		}
		k++
	}
	k = 0
	for i-k >= 0 && j+k < len(queenAry) {
		if queenAry[i-k][j+k] == 1 {
			return false
		} else if queenAry[i-k][j+k] == 0 {
			road = append(road, []int{i - k, j + k})
		}
		k++
	}
	k = 0
	for i+k < len(queenAry) && j-k > 0 {
		if queenAry[i+k][j-k] == 1 {
			return false
		} else if queenAry[i+k][j-k] == 0 {
			road = append(road, []int{i + k, j - k})
		}
		k++
	}
	k = 0
	for i+k < len(queenAry) && j+k < len(queenAry) {
		if queenAry[i+k][j+k] == 1 {
			return false
		} else if queenAry[i+k][j+k] == 0 {
			road = append(road, []int{i + k, j + k})
		}
		k++
	}

	for _, v := range road {
		queenAry[v[0]][v[1]] = 2
	}
	queenAry[i][j] = 1
	// log.Printf("queenAry : %+v", queenAry)

	if j+1 == len(queenAry) {
		return true
	}

	if j+1 < len(queenAry) {
		for k := 0; k < len(queenAry); k++ {
			// log.Printf("len(queenAry) : %+v k : %+v  j+1 : %+v", len(queenAry), k, j+1)
			copyQueenAry := [][]int{}
			for _, copyRow := range queenAry {
				copy := append([]int{}, copyRow...)
				copyQueenAry = append(copyQueenAry, copy)
			}
			ok := addQueen(k, j+1, queenAry)
			if ok {
				queen := []string{}
				for _, row := range queenAry {
					queenStr := ""
					for _, v := range row {
						if v == 1 {
							queenStr = queenStr + "Q"
						} else {
							queenStr = queenStr + "."
						}
					}
					queen = append(queen, queenStr)
				}
				result = append(result, queen)
				log.Printf("queen : %+v", queen)
			}
			queenAry = copyQueenAry
		}
	}

	return false
}
