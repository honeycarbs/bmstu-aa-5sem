package mtx

import (
	"fmt"
	"math/rand"
	"time"
)

type Matrix struct {
	Rows int
	Cols int
	Values [][]int
}

func Allocate(rows, cols int) Matrix {
	var dest Matrix

	dest.Rows = rows
	dest.Cols = cols
	dest.Values = make([][]int, dest.Rows)
	for i := range dest.Values {
		dest.Values[i] = make([]int, dest.Cols)
	}

	return dest
}

func Input(rows, cols int) Matrix {
	dest := Allocate(rows, cols)

	for i := 0; i < dest.Rows; i++ {
		for j := 0; j < dest.Cols; j++ {
			fmt.Scanf("%d", &dest.Values[i][j])
		}
	}

	return dest
}

func Randomize(src Matrix) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < src.Rows; i++ {
		for j := 0; j < src.Cols; j++ {
			src.Values[i][j] = rand.Intn(10) - rand.Intn(10)
		}
	}
}

func __logSqCasesOpen(i, rows int) {
	switch (i) {
	case 0:
		fmt.Print("┌")
	case rows - 1:
		fmt.Print("└")
	default:
		fmt.Print("│")
	}
}

func __logSqCasesColse(i, rows int) {
	switch (i) {
	case 0:
		fmt.Print("┐")
	case rows - 1:
		fmt.Print("┘")
	default:
		fmt.Print("│")
	}
}

func Log(src Matrix) {
	for i := 0; i < src.Rows; i++ {
		__logSqCasesOpen(i, src.Rows)
		for j := 0; j < src.Cols; j++ {
			fmt.Printf("%5d ", src.Values[i][j])
		}
		__logSqCasesColse(i, src.Rows)
		fmt.Println("")
	}
}

func LogLatex(src Matrix) {
	fmt.Println("\\begin{bmatrix} ")
	for i := 0; i < src.Rows; i++ {
		for j := 0; j < src.Cols; j++ {
			fmt.Printf("%d ", src.Values[i][j])
			if (j == src.Cols - 1) {
				fmt.Print("\\\\")
			} else {
				fmt.Print("& ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("\\end{bmatrix}")
}
