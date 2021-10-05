package multiply

import (
	mtx "../mtx"
)

func Algebraic(mtx_left, mtx_right mtx.Matrix) mtx.Matrix {
	var product mtx.Matrix
	product = mtx.Allocate(mtx_left.Rows, mtx_right.Cols)
	
	for i := 0; i < product.Rows; i++ {
		for j := 0; j < product.Cols; j++ {
			for k := 0; k < mtx_right.Rows; k++ {
				product.Values[i][j] = product.Values[i][j] + mtx_left.Values[i][k] * mtx_right.Values[k][j]
			}
		}
	}

	return product
}

func WinogradOptimized(mtx_left, mtx_right mtx.Matrix) mtx.Matrix {
	var (
		row_count, col_count_left, col_count_right int
		product mtx.Matrix
		rows_precomp, cols_precomp []int
	)
	row_count = mtx_left.Rows
	col_count_left = mtx_left.Cols
	col_count_right = mtx_right.Cols

	product = mtx.Allocate(row_count, col_count_right)

	rows_precomp = __precomputeRowsOptimized(rows_precomp, mtx_left, row_count, col_count_left)
	cols_precomp = __precomputeColsOptimized(cols_precomp, mtx_right, col_count_left, col_count_right)

	var i, j, k int
	is_odd := col_count_left % 2
	for  i = 0; i < row_count; i++ {
		for j = 0; j < col_count_right; j++ {
			product.Values[i][j]= -(rows_precomp[i] + cols_precomp[j])
			for k = 1; k < col_count_left; k+= 2 {
				product.Values[i][j] += (mtx_left.Values[i][k - 1] +
					+ mtx_right.Values[k][j]) * (mtx_left.Values[i][k] + mtx_right.Values[k - 1][j])
			}
			if (is_odd == 1) {
				product.Values[i][j] += mtx_left.Values[i][col_count_left - 1] * +
				 mtx_right.Values[col_count_left - 1][j]
			}
		}
	}

	return product
}

func __precomputeRowsOptimized(rows_precomp []int, mtx_left mtx.Matrix, row_count, col_count int) []int {
	rows_precomp = make([]int, row_count)
	for i := 0; i < row_count; i++ {
		for j := 1; j < col_count; j+= 2 {
			rows_precomp[i] += mtx_left.Values[i][j - 1] * mtx_left.Values[i][j]
		}
	}
	return rows_precomp
}

func __precomputeColsOptimized(cols_precomp []int, mtx_right mtx.Matrix, col_count_left, col_count_right int) []int {
	cols_precomp = make([]int, col_count_right)
	for i := 0; i < col_count_right; i++ {
		for j := 1; j < col_count_left; j+=2 {
			cols_precomp[i] += (mtx_right.Values[j - 1][i] * mtx_right.Values[j][i])
		}
	}
	return cols_precomp
}

func Winograd(mtx_left, mtx_right mtx.Matrix) mtx.Matrix {
	var (
		row_count, col_count_left, col_count_right int
		product mtx.Matrix
		rows_precomp, cols_precomp []int
	)
	row_count = mtx_left.Rows
	col_count_left = mtx_left.Cols
	col_count_right = mtx_right.Cols

	product = mtx.Allocate(row_count, col_count_right)

	rows_precomp = __precomputeRows(rows_precomp, mtx_left, row_count, col_count_left)
	cols_precomp = __precomputeCols(cols_precomp, mtx_right, col_count_left, col_count_right)
	
	for  i := 0; i < row_count; i++ {
		for j := 0; j < col_count_right; j++ {
			product.Values[i][j]= -(rows_precomp[i] + cols_precomp[j])
			for k := 1; k < col_count_left; k+= 2 {
				product.Values[i][j] = product.Values[i][j] + (mtx_left.Values[i][k - 1] +
					+ mtx_right.Values[k][j]) * (mtx_left.Values[i][k] + mtx_right.Values[k - 1][j])
			}
		}
	}
	if (col_count_left % 2 != 0) {
		for  i := 0; i < row_count; i++ {
			for j := 0; j < col_count_right; j++ {
				product.Values[i][j] = product.Values[i][j] +  mtx_left.Values[i][col_count_left - 1] * mtx_right.Values[col_count_left - 1][j]
			}
		}
	}
	return product
}

func __precomputeRows(rows_precomp []int, mtx_left mtx.Matrix, row_count, col_count int) []int {
	rows_precomp = make([]int, row_count)
	for i := 0; i < row_count; i++ {
		for j := 1; j < col_count; j+= 2 {
			rows_precomp[i] = rows_precomp[i] + mtx_left.Values[i][j - 1] * mtx_left.Values[i][j]
		}
	}
	return rows_precomp
}
func __precomputeCols(cols_precomp []int, mtx_right mtx.Matrix, col_count_left, col_count_right int) []int {
	cols_precomp = make([]int, col_count_right)
	for i := 0; i < col_count_right; i++ {
		for j := 1; j < col_count_left; j+=2 {
			cols_precomp[i] = cols_precomp[i] + (mtx_right.Values[j - 1][i] * mtx_right.Values[j][i])
		}
	}
	return cols_precomp
}