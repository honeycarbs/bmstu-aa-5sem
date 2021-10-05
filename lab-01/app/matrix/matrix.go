package matrix

import "fmt"

type IMtx [][]int

func IMtxInit(rows, cols, filler int) IMtx {
	mtx := make(IMtx, rows)
	for i := range mtx {
		mtx[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			mtx[i][j] = filler
		}
	}
	return mtx
}

func IMtxLog(mtx IMtx, rows, cols int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(mtx[i][j])
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}