package DMtx

import matrix "../matrix"
import utils "../utils"

// import "fmt"

func LevenshteinMatrixIterative(src, dest []rune) int {
	srcLen, destLen := len(src), len(dest)

	rows, cols := srcLen + 1, destLen + 1
	DMtx := matrix.IMtxInit(rows, cols, 0)

	for i := 1; i < rows; i++ {
		DMtx[i][0] = i
	}
	for j := 1; j < cols; j++ {
		DMtx[0][j] = j
	}

	match := 0
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if src[i - 1] == dest[j - 1] {
				match = 0
			} else {
				match = 1
			}
			insert  := DMtx[i][j - 1] + 1
			delete  := DMtx[i - 1][j] + 1
			replace := DMtx[i - 1][j - 1] + match

			min := utils.MinOfInt(insert, delete, replace)
			DMtx[i][j] = min
		}
	}
	return DMtx[rows - 1][cols - 1]
}

func LevenshteinPartialMemo(src, dest []rune) int {
	srcLen, destLen := len(src), len(dest)

	rows, cols := srcLen + 1, destLen + 1
	DMtx := matrix.IMtxInit(rows, cols, -1)

	ans := _partialMemoHelper(src, dest, srcLen, destLen, DMtx)
	return ans
}

func _partialMemoHelper(src, dest []rune, i, j int, 
						DMtx matrix.IMtx) int {
	if (utils.MinOfInt(i, j) == 0) {
		return utils.Max2Int(i, j)
	}
	if DMtx[i][j] != -1 {
		return DMtx[i][j]
	}

	match := 1
	if src[i - 1] == dest[j - 1] {
		match = 0
	}

	insert  := _partialMemoHelper(src, dest, i, j - 1, DMtx) + 1
	delete  := _partialMemoHelper(src, dest, i - 1, j, DMtx) + 1
	replace := match + _partialMemoHelper(src, dest, i - 1, j - 1, DMtx)

	min := utils.MinOfInt(insert, delete, replace)
	DMtx[i][j] = min
	
	return DMtx[i][j]
}

func LevenshteinRecursive(src, dest []rune) int {
	srcLen, destLen := len(src), len(dest)

	ans := _lRecursiveHelper(src, dest, srcLen, destLen)

	return ans
}

func _lRecursiveHelper(src, dest []rune, i, j int) int {
	if (utils.MinOfInt(i, j) == 0) {
		return utils.Max2Int(i, j)
	}

	match := 1
	if (src[i - 1] == dest[j - 1]) {
		match = 0
	}

	insert := _lRecursiveHelper(src, dest, i, j - 1) + 1
	delete := _lRecursiveHelper(src, dest, i - 1, j) + 1
	replace := match + _lRecursiveHelper(src, dest, i - 1, j - 1)

	min := utils.MinOfInt(insert, delete, replace)
	return min
}

func DamerauLevenshteinIterative(src, dest []rune) int{
	srcLen, destLen := len(src), len(dest)

	rows, cols := srcLen + 1, destLen + 1
	DMtx := matrix.IMtxInit(rows, cols, 0)

	for i := 1; i < rows; i++ {
		DMtx[i][0] = i
	}
	for j := 1; j < cols; j++ {
		DMtx[0][j] = j
	}

	match := 0
	min := 0
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if src[i - 1] == dest[j - 1] {
				match = 0
			} else {
				match = 1
			}

			insert := DMtx[i][j - 1] + 1
			delete := DMtx[i - 1][j] + 1
			replace := DMtx[i - 1][j - 1] + match
			transpose := -1

			if i > 1 && j > 1 {
				transpose = DMtx[i - 2][j - 2] + 1
			}

			if transpose != -1 && src[i - 1] == dest[j - 2] && src[i - 2] == dest[j - 1] {
				min = utils.MinOfInt(insert, delete, replace, transpose)
			} else {
				min = utils.MinOfInt(insert, delete, replace)
			}
			DMtx[i][j] = min
		}
	}
	DMtx[srcLen][destLen] = min
	return DMtx[srcLen][destLen]
}

func DamerauLevenshteinRecursive(src, dest []rune) int {
	srcLen, destLen := len(src), len(dest)

	ans := _dRecursiveHelper(src, dest, srcLen, destLen)

	return ans
}

func _dRecursiveHelper(src, dest []rune, i, j int) int {
	if (utils.MinOfInt(i, j) == 0) {
		return utils.Max2Int(i, j)
	}

	match := 1
	if (src[i - 1] == dest[j - 1]) {
		match = 0
	}

	insert := _dRecursiveHelper(src, dest, i, j - 1) + 1
	delete := _dRecursiveHelper(src, dest, i - 1, j) + 1
	replace := match + _dRecursiveHelper(src, dest, i - 1, j - 1)

	transpose := -1

	if i > 1 && j > 1 {
		transpose =_dRecursiveHelper(src, dest, i - 2, j - 2) + 1
	}

	min := 0
	if transpose != -1 && src[i - 1] == dest[j - 2] && src[i - 2] == dest[j - 1] {
		min = utils.MinOfInt(insert, delete, replace, transpose)
	} else {
		min = utils.MinOfInt(insert, delete, replace)
	}
	return min
}