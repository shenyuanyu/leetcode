package tasks

import (
	"fmt"
	"reflect"
)

type position struct {
	Row int
	Col int
}

// 解决数独问题
func SolveSudoku(board [][]byte) {
	notSolvePosPossibleMap := make(map[position]map[byte]struct{})
	// 初始化
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			k := board[i][j]
			if k == '.' { // 未确定
				pos := position{i, j}
				notSolvePosPossibleMap[pos] = generateMapByByteRange('1', '9')

				excludeByRow(pos, notSolvePosPossibleMap, board)
				excludeByCol(pos, notSolvePosPossibleMap, board)
				excludeByTinySudoku(pos, notSolvePosPossibleMap, board)

				if len(notSolvePosPossibleMap[pos]) == 1 {
					for key := range notSolvePosPossibleMap[pos] {
						board[i][j] = key
						delete(notSolvePosPossibleMap, pos)
						break
					}
				}
			}
		}
	}

	// 循环检测
	prev := 0
	for len(notSolvePosPossibleMap) != 0 && prev != len(notSolvePosPossibleMap) {
		prev = len(notSolvePosPossibleMap)
		for pos := range notSolvePosPossibleMap {
			excludeByRow(pos, notSolvePosPossibleMap, board)
			excludeByCol(pos, notSolvePosPossibleMap, board)
			excludeByTinySudoku(pos, notSolvePosPossibleMap, board)

			if len(notSolvePosPossibleMap[pos]) == 1 {
				for key := range notSolvePosPossibleMap[pos] {
					board[pos.Row][pos.Col] = key
					delete(notSolvePosPossibleMap, pos)
					break
				}
			}
		}
	}
}

// 生成start到end的byte数组, [start, end]
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

	notSolvePos := make([]position, 0, 9)
	for j := 0; j < 9; j++ {
		if pos.Col == j {
			continue
		}

		k := board[pos.Row][j]
		if k == '.' { // 未确定位置, 加入队列
			notSolvePos = append(notSolvePos, position{pos.Row, j})
		} else { // 已确定位置, 直接排除
			if _, ok := posPossible[k]; ok {
				delete(posPossible, k)
			}
		}
	}

	excludeByPossible(pos, notSolvePos, notSolvePosPossible)
	excludeByNotPossible(pos, notSolvePos, notSolvePosPossible)
}

// 根据列排除可能性
func excludeByCol(pos position, notSolvePosPossible map[position]map[byte]struct{}, board [][]byte) {
	
	posPossible := notSolvePosPossible[pos]

	notSolvePos := make([]position, 0, 9)
	for i := 0; i < 9; i++ {
		if pos.Row == i {
			continue
		}

		k := board[i][pos.Col]
		if k == '.' { // 未确定位置, 加入队列
			notSolvePos = append(notSolvePos, position{i, pos.Col})
		} else { // 已确定位置, 直接排除
			if _, ok := posPossible[k]; ok {
				delete(posPossible, k)
			}
		}
	}

	excludeByPossible(pos, notSolvePos, notSolvePosPossible)
	excludeByNotPossible(pos, notSolvePos, notSolvePosPossible)

}

// 根据小九宫排除
func excludeByTinySudoku(pos position, notSolvePosPossible map[position]map[byte]struct{}, board [][]byte) {
	posPossible := notSolvePosPossible[pos]

	// 确定小九宫起始位置
	row := pos.Row / 3 * 3
	col := pos.Col / 3 * 3

	notSolvePos := make([]position, 0, 9)
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if i == pos.Row && j == pos.Col {
				continue
			}

			k := board[i][j]
			if k == '.' { // 未确定
				notSolvePos = append(notSolvePos, position{i, j})
			} else {
				if _, ok := posPossible[k]; ok {
					delete(posPossible, k)
				}
			}
		}
	}

	// 可能性排除法
	excludeByPossible(pos, notSolvePos, notSolvePosPossible)
	excludeByNotPossible(pos, notSolvePos, notSolvePosPossible)
}

func excludeByPossible(pos position, notSolvePos []position,
	notSolvePosPossibleMap map[position]map[byte]struct{}) {

	posPossibleMap := notSolvePosPossibleMap[pos]

	// 计算每个可能的出现率
	count := make([]int, 0, len(notSolvePos))
	for i := 0; i < len(notSolvePos)-1; i++ {
		k := 0
		for j := i + 1; j < len(notSolvePos); j++ {
			if reflect.DeepEqual(notSolvePosPossibleMap[notSolvePos[i]], notSolvePosPossibleMap[notSolvePos[j]]) {
				k++
				notSolvePos = deletePositionArray(notSolvePos, j)
				j--
			}
		}
		count = append(count, k)
	}

	// 根据出现率删除
	for i := 0; i < len(count); i++ {
		if len(notSolvePosPossibleMap[notSolvePos[i]]) == count[i] {
			for key := range notSolvePosPossibleMap[notSolvePos[i]] {
				delete(posPossibleMap, key)
			}
		}
	}
}

func excludeByNotPossible(pos position, notSolvePos []position,
	notSolvePosPossibleMap map[position]map[byte]struct{}) {

	posPossibleMap := notSolvePosPossibleMap[pos]

	for posPossible := range posPossibleMap {
		b := true
		for _, notPos := range notSolvePos {
			if _, ok := notSolvePosPossibleMap[notPos][posPossible]; ok {
				b = false
			}
		}

		if b {
			notSolvePosPossibleMap[pos] = map[byte]struct{}{
				posPossible: struct{}{},
			}
			return
		}
	}

}

func deletePositionArray(arr []position, i int) []position {
	return append(arr[:i], arr[i+1:]...)
}

// PrintBoard ...
func PrintBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%s\t", string(board[i][j]))
		}
		fmt.Println()

	}
}
