package tasks

import (
	"reflect"
)

type position struct {
	Row int
	Col int
}

// 解决数独问题
func solveSudoku(board [][]byte) {

}

func generateMapByByteRange(start, end byte) map[byte]struct{} {
	hash := make(map[byte]struct{})
	for i := start; i <= end; i++ {
		hash[i] = struct{}{}
	}
	return hash
}

// 根据行排除可能性
func excludeByRow(pos position, notSolvePosPossible map[position]map[byte]struct{}, board [][]byte) {
	posPossible := notSolvePosPossible[pos]

	notSolvePos := make([]position, 0, len(board[pos.Row]))
	for j := 0; j < len(board[pos.Row]); j++ {
		if pos.Col == j {
			continue
		}

		k := board[pos.Row][j]
		if k == '.' {
			notSolvePos = append(notSolvePos, position{pos.Row, j})
		} else {
			if _, ok := posPossible[k]; ok {
				delete(posPossible, k)
			}
		}
	}

}

func excludeByPossible(notSolvePos []position, notSolvePosPossibleMap map[position]map[byte]struct{}) map[byte]struct{} {
	count := make([]int, 0, len(notSolvePos))
	for i := 0; i < len(notSolvePos)-1; i++ {
		k := 0
		for j := i; j < len(notSolvePos); j++ {
			if reflect.DeepEqual(notSolvePosPossibleMap[notSolvePos[i]], notSolvePosPossibleMap[notSolvePos[j]]) {
				k++
				notSolvePos = deletePositionArray(notSolvePos, j)
				j--
			}
		}
		count = append(count, k)
	}

	i

}

func deletePositionArray(arr []position, i int) []position {
	return append(arr[:i], arr[i+1:]...)
}
