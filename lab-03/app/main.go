package main

import (
	"fmt"
	"os"
	"strconv"

	"./generator"
	"./perftest"
	"./sort"
	"github.com/inancgumus/screen"
)

func spacing(num int) {
	for i := 0; i < num; i++ {
		fmt.Print(" ")
	} 
}

func hat()  {
	center := len([]rune("┌─┐┬─┐┬─┐┌─┐┬ ┬  ┌─┐┌─┐┬─┐┌┬┐┬┌┐┌┌─┐"))
	w, _ := screen.Size()

	space := (w - center) / 2

	spacing(space)
	fmt.Println("┌─┐┬─┐┬─┐┌─┐┬ ┬  ┌─┐┌─┐┬─┐┌┬┐┬┌┐┌┌─┐")
	spacing(space)
	fmt.Println("├─┤├┬┘├┬┘├─┤└┬┘  └─┐│ │├┬┘ │ │││││ ┬")
	spacing(space)
	fmt.Println("┴ ┴┴└─┴└─┴ ┴ ┴   └─┘└─┘┴└─ ┴ ┴┘└┘└─┘")
	spacing(space)
	fmt.Println("   Let's sort an array togeter!")
	fmt.Println()
	
}

func dash() {
	w, _ := screen.Size()
	for i := 0; i < w; i++ {
		fmt.Print("-")
	}  
	fmt.Println()
}

func slicePrint(slice []int) {
	for _, value := range slice {  
		fmt.Printf("%d ", value)  
	   }  
	fmt.Println()
}


func main() {
	if (len(os.Args) != 1) {
		if (os.Args[1] == "-PROF") {
			perftest.ProfileWrapper()
			os.Exit(0)
		}
	}
	var (
		resp string
		psort bool
	)

	screen.Clear()
	screen.MoveTopLeft()
	hat()

	fmt.Print("How many numbers should I generate? ")
	fmt.Fscan(os.Stdin, &resp)

	size, err := strconv.Atoi(resp)
	for err != nil {
		fmt.Print("Looks like it's not a number. Try again: ")
		fmt.Fscan(os.Stdin, &resp)
		size, err = strconv.Atoi(resp)
	}
	

	fmt.Print("Should I generate a sorted array? (If no - press 0, if yes - 1) ")
	fmt.Fscan(os.Stdin, &resp)
	psort, err = strconv.ParseBool(resp)

	for err != nil {
		fmt.Print("I don't understand. If no - press 0, if yes - 1 : ")
		fmt.Fscan(os.Stdin, &resp)
		psort, err = strconv.ParseBool(resp)
	}


	items := generator.Slice(size, psort)

	fmt.Println(items)

	screen.Clear()
	screen.MoveTopLeft()
	hat()


	fmt.Println("Your array:")
	slicePrint(items)
	dash()
	var ins_items = make([]int, size)
	var sel_items = make([]int, size)
	var heap_items = make([]int, size)

	fmt.Println("Insertion sort result: ")
	copy(ins_items, items)
	sort.InsertionSort(ins_items)
	slicePrint(ins_items)

	fmt.Println("\nSelection sort result: ")
	copy(sel_items, items)
	sort.InsertionSort(sel_items)
	slicePrint(sel_items)

	fmt.Println("\nHeap sort result: ")
	copy(heap_items, items)
	sort.InsertionSort(heap_items)
	slicePrint(heap_items)
}