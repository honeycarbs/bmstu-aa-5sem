package interf

import (
	"fmt"
	"strings"

	"../aco"
	"../brute"
	"github.com/inancgumus/screen"
)

func RunUserMode(conf_filename, graph_filename string, daycount int) int {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Print(strings.Repeat("*", 10))
	fmt.Printf(" CONFIG FILE: %v ", conf_filename)
	fmt.Println(strings.Repeat("*", 10))
	
	fmt.Println("")

	fmt.Print(strings.Repeat("*", 10))
	fmt.Printf(" GRAPH FILE: %v ", graph_filename)
	fmt.Println(strings.Repeat("*", 10))


	buf := (aco.ParceGraphFile(graph_filename))
	if buf == nil {
		return 1
	}

	fmt.Printf("\nAnt colony optimization results: ")
	graph := *buf

	r_buf := aco.ACOWrapper(graph, daycount, conf_filename)
	if r_buf != nil {
		roots := *r_buf
		for _, ant := range roots {
			fmt.Printf("%v ", ant)
		}
	}

	fmt.Printf("\n\nBrute force approach results: ")
	roots := brute.BruteForce(graph)

	for _, path := range roots {
		fmt.Printf("%v ", path)
	}
	fmt.Println("")

	return 0
}

func RunAutoconfMode(conf_filename, graph_filename string, daycount int) int {
	buf := (aco.ParceGraphFile(graph_filename))
	if buf == nil {
		return 1
	}

	graph := *buf

	r_buf := aco.ACOWrapper(graph, daycount, conf_filename)
	if r_buf != nil {
		roots := *r_buf
		for _, ant := range roots {
			fmt.Printf("%v ", ant)
		}
	}

	return 0
}