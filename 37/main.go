package main

import (
	"log"
)

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	log.Printf("%+v", board)
}

func solveSudoku(board [][]byte) {
	solveSudokuMy(board)
}

func solveSudokuMy(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				numsCheck := getValidSudoku(board, i, j)
				for ch := '1'; ch <= '9'; ch++ {
					if !numsCheck[ch-'1'] {
						continue
					}
					board[i][j] = byte(ch)
					if solveSudokuMy(board) {
						return true
					}
				}
				board[i][j] = '.'
				return false
			}
		}
	}
	return true
}

func getValidSudoku(board [][]byte, i int, j int) [9]bool {
	numsCheck := [9]bool{true, true, true, true, true, true, true, true, true}

	for k := 0; k < 9; k++ {
		if board[i][k] != '.' {
			num := int(board[i][k] - '1')
			numsCheck[num] = false
		}
		if board[k][j] != '.' {
			num := int(board[k][j] - '1')
			numsCheck[num] = false
		}
		if board[int(i/3)*3+int(k/3)][int(j/3)*3+k%3] != '.' {
			num := int(board[int(i/3)*3+int(k/3)][int(j/3)*3+k%3] - '1')
			numsCheck[num] = false
		}
	}
	return numsCheck
}
