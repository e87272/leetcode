package main

import "log"

func main() {
	log.Printf("%+v", isValidSudoku([][]byte{
		{'.', '.', '.', '.', '.', '.', '5', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'9', '3', '.', '.', '2', '.', '4', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '3', '4', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '5', '2', '.', '.'},
	}))
}
func isValidSudoku(board [][]byte) bool {
	countRow := [9][9]bool{}
	countCol := [9][9]bool{}
	countNight := [9][9]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := int(board[i][j] - '1')
			if countRow[i][num] {
				return false
			}
			if countCol[j][num] {
				return false
			}
			if countNight[int(i/3)*3+int(j/3)][num] {
				return false
			}
			countRow[i][num] = true
			countCol[j][num] = true
			countNight[int(i/3)*3+int(j/3)][num] = true
		}
	}
	return true
}
