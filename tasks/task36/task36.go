package task36

// 验证数独是否合法
func IsValidSudoku(board [][]byte) bool {
	// row and column
	for i := 0; i < 9; i++ {
		if !isValidSudokuRow(board, i) {
			return false
		}

		if !isValidSudokuCol(board, i) {
			return false
		}
	}
	if !isValidSudokuGrid(board) {
		return false
	}

	return true
}

func isValidSudokuRow(board [][]byte, r int) bool {
	tmp := make(map[byte]struct{})
	for i := 0; i < 9; i++ {
		t := board[r][i]
		if t == '.' {
			continue
		}
		if t < '0' || t > '9' {
			return false
		}

		if _, ok := tmp[t]; ok {
			return false
		}
		tmp[t] = struct{}{}
	}
	return true
}

func isValidSudokuCol(board [][]byte, c int) bool {
	tmp := make(map[byte]struct{})
	for i := 0; i < 9; i++ {
		t := board[i][c]
		if t == '.' {
			continue
		}
		if t < '0' || t > '9' {
			return false
		}

		if _, ok := tmp[t]; ok {
			return false
		}
		tmp[t] = struct{}{}
	}
	return true
}

func isValidSudokuGrid(board [][]byte) bool {

	for m := 0; m < 9; m += 3 {
		for n := 0; n < 9; n += 3 {
			tmp := make(map[byte]struct{})
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					t := board[i+m][j+n]
					if t == '.' {
						continue
					}
					if t < '0' || t > '9' {
						return false
					}

					if _, ok := tmp[t]; ok {
						return false
					}
					tmp[t] = struct{}{}
				}
			}
		}
	}
	return true
}
