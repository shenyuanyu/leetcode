package task48

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		for j := i; j < len(matrix)-1-i; j++ {
			around(matrix, i, j)
		}
	}
}

func around(matrix [][]int, x, y int) {
	n := len(matrix) - 1
	for i := 0; i < 3; i++ {
		a, b := beReplace(n, x, y)
		matrix[x][y], matrix[a][b] = matrix[a][b], matrix[x][y]
		x, y = a, b
	}
}

func beReplace(n, x, y int) (int, int) {
	return n - y, x
}
