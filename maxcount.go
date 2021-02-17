package main

func maxCount(m int, n int, ops [][]int) int {
	row, col := m, n
	for _, op := range ops {
		row, col = min(row, op[0]), min(col, op[1])
	}

	return row * col
}

func maxCount2(m int, n int, ops [][]int) int {
	row, col := m, n
	matrix := make([][]int, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]int, col)
	}

	count := 0

	for _, op := range ops {
		for i := 0; i < op[0]; i++ {
			for j := 0; j < op[1]; j++ {
				matrix[i][j]++
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[0][0] == matrix[i][j] {
				count++
			}
		}
	}

	return count
}
