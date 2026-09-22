package striver_dsa_sheet

// Problem: https://leetcode.com/problems/set-matrix-zeroes/
// https://www.youtube.com/shorts/gIs7OmSVL4E

func setZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])
	firstRowZero := false
	firstColZero := false
	// Checking first col
	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	//Checking first row
	for i := 0; i < cols; i++ {
		if matrix[0][i] == 0 {
			firstRowZero = true
			break
		}
	}

	//makring first row and column if any current element is 0

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Now based on first row and column setting other left array as 0 where needed

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// now since other rest of matrix is done only first col and row are left for marking

	if firstColZero {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}

	if firstRowZero {
		for i := 0; i < cols; i++ {
			matrix[0][i] = 0
		}
	}
}
