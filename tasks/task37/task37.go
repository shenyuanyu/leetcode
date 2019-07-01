package task37

import (
	"fmt"
)

func solveSudoku(board [][]byte) {
	var rowUsed [9]uint16
	var colUsed [9]uint16
	var minSudokuUsed [9]uint16

	isSolved := false
	for !isSolved {
		isSolved = recursiveSolveSudoku(&rowUsed, &colUsed, &minSudokuUsed, board)
	}
}

func recursiveSolveSudoku(rowUsed *[9]uint16, colUsed *[9]uint16, minSudokuUsed *[9]uint16, board [][]byte) bool {
	isSolved := false

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				t := ^rowUsed[i] & ^colUsed[j] & ^minSudokuUsed[getMinSudokuIndex(i, j)]
				rs := toBytes(t)
				if len(rs) == 1 {
					board[i][j] = rs[0]
				}
			}

			if board[i][j] != '.' {
				var t uint16 = 1 << (board[i][j] - '1')
				rowUsed[i] |= t
				colUsed[j] |= t
				minSudokuUsed[getMinSudokuIndex(i, j)] |= t
			} else {
				isSolved = false
			}
		}
	}

	printSudoku(board)
	return isSolved
}

func getMinSudokuIndex(i, j int) int {
	return i/3*3 + j/3
}

func toBytes(candidate uint16) []byte {
	var j uint16 = 1

	var resp []byte
	var i byte = '1'
	for ; i <= '9'; i++ {
		if candidate&j != 0 {
			resp = append(resp, i)
		}

		j <<= 1
	}

	return resp
}

func printSudoku(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			fmt.Printf("%c\t", board[i][j])
		}
		fmt.Println()
	}

	fmt.Println("------------------------------------")
}
