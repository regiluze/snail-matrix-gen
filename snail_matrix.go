package snail_matrix

import ()

func matrix_wrapper_with(matrix [][]int, number int) [][]int {
	if len(matrix) == 0 {
		return [][]int{[]int{number}}
	}
	newMatrixSize := len(matrix[0]) + 2
	newMatrix := make([][]int, 0, newMatrixSize)
	topLine := make([]int, newMatrixSize)
	for i := range topLine {
		topLine[i] = number
	}
	newMatrix = append(newMatrix, topLine)
	for _, line := range matrix {
		newLine := make([]int, newMatrixSize)
		newLine = append([]int{number}, line...)
		newLine = append(newLine, number)
		newMatrix = append(newMatrix, newLine)
	}
	bottomLine := make([]int, newMatrixSize)
	for i := range bottomLine {
		bottomLine[i] = number
	}
	return append(newMatrix, bottomLine)
}

func ssnail_matrix_for(number int) [][]int {
	var snailMatrix [][]int
	for i := 0; i <= number; i++ {
		snailMatrix = matrix_wrapper_with(snailMatrix, i)
	}
	return snailMatrix
}

func snail_matrix_for(number int) [][]int {
	var snailMatrix [][]int
	for i := 0; i <= number; i++ {
		snailMatrix = matrix_wrapper_with(snailMatrix, i)
	}
	return snailMatrix
}
