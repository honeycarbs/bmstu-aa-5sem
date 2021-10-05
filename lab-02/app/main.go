package main

import (
	"fmt"
	"os"
	"strconv"

	"./mtx"
	mult "./multiply"
	"./perftest"
	"github.com/inancgumus/screen"
)

func spacing(num int) {
	for i := 0; i < num; i++ {
		fmt.Print(" ")
	} 
}


func hat()  {
	center := len([]rune("\\_\\_/_/--|_|_\\_\\_\\_\\_|_|_/_/--|_\\_\\_|_| \\ "))
	w, _ := screen.Size()

	space := (w - center) / 2

	spacing(space)
	fmt.Println("         _     _________ _ _              ")
	spacing(space)
	fmt.Println("        | |\\/|/ /\\| | |_| \\ \\_/           ")
	spacing(space)
	fmt.Println(" __   __|_|  /_/--|_|_| |_/_/_\\_____ ___  ")
	spacing(space)
	fmt.Println("/ /` / /\\| | / /`| | | |  / /\\| / / | |_) ")
	spacing(space)
	fmt.Println("\\_\\_/_/--|_|_\\_\\_\\_\\_|_|_/_/--|_\\_\\_|_| \\ ")


	center = len([]rune("It can only multiply. And that's it."))
	spacing(space)
	fmt.Println("It can only multiply. And that's it.")
	fmt.Println()
	
}

func main() {
	var latex bool = false
	if len(os.Args) >= 2 {
		if os.Args[1] == "-latex" {
			latex = true
		}
		if os.Args[1] == "-prof" {
			perftest.ProfileWrapper()
			os.Exit(0)
		}
	}
	screen.Clear()
	screen.MoveTopLeft()
	hat()
	fmt.Println("Your matrix is giong to be randomozed. I am not mentally ready to code an input.")

	var ans string
	fmt.Print("Input rows number of multiplier: ")
	fmt.Fscan(os.Stdin, &ans)

	rows1, err := strconv.Atoi(ans)
	for err != nil {
		fmt.Print("Looks like it's not a number. Try again: ")
		fmt.Fscan(os.Stdin, &ans)
		rows1, err = strconv.Atoi(ans)
	}

	fmt.Print("Input cols number of multiplier: ")
	fmt.Fscan(os.Stdin, &ans)

	cols1, err := strconv.Atoi(ans)
	for err != nil {
		fmt.Print("Looks like it's not a number. Try again: ")
		fmt.Fscan(os.Stdin, &ans)
		cols1, err = strconv.Atoi(ans)
	}

	fmt.Print("Input rows number of multicand: ")
	fmt.Fscan(os.Stdin, &ans)
	rows2, err := strconv.Atoi(ans)
	for err != nil || rows2 != cols1 {
		if (err != nil) {
			fmt.Print("Looks like it's not a number. Try again: ")
		}
		if (rows2 != cols1) {
			fmt.Print("I can't multiply it. rows number of multiplier should be equal of cols of multiplicand.\nTry again: ")
		}
	
		fmt.Fscan(os.Stdin, &ans)
		rows2, err = strconv.Atoi(ans)
	}
	

	fmt.Print("Input cols number of multiplicand: ")
	fmt.Fscan(os.Stdin, &ans)

	cols2, err := strconv.Atoi(ans)
	for err != nil {
		fmt.Print("Looks like it's not a number. Try again: ")
		fmt.Fscan(os.Stdin, &ans)
		cols2, err = strconv.Atoi(ans)
	}

	var mtx_left, mtx_right mtx.Matrix

	mtx_left = mtx.Allocate(rows1, cols1)
	mtx_right = mtx.Allocate(rows2, cols2)

	mtx.Randomize(mtx_right)
	mtx.Randomize(mtx_left)

	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("Your matrixes: ")
	// mtx.Log(mtx_left)
	if (latex) {
		mtx.LogLatex(mtx_left)
	} else {
		mtx.Log(mtx_left)
	}
	fmt.Println("")
	if (latex) {
		mtx.LogLatex(mtx_right)
	} else {
		mtx.Log(mtx_right)
	}

	var alg, win, win_opt mtx.Matrix
	alg = mult.Algebraic(mtx_left, mtx_right)
	win = mult.Winograd(mtx_left, mtx_right)
	win_opt = mult.WinogradOptimized(mtx_left, mtx_right)

	if (latex) {
		fmt.Println("\n\nResult of algebraic approach: ")
		mtx.LogLatex(alg)

		fmt.Println("\n\nResult of Collersmith-Winograd algorithm: ")
		mtx.LogLatex(win)

		fmt.Println("\n\nResult of optimized Collersmith-Winograd algorithm: ")
		mtx.LogLatex(win_opt)
	} else {
		fmt.Println("\n\nResult of algebraic approach: ")
		mtx.Log(alg)

		fmt.Println("\n\nResult of Collersmith-Winograd algorithm: ")
		mtx.Log(win)

		fmt.Println("\n\nResult of optimized Collersmith-Winograd algorithm: ")
		mtx.Log(win_opt)
	}
}