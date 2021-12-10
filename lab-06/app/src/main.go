package main

import (
	// "fmt"

	"fmt"
	"os"
	"strconv"

	"./interf"
	// "./aco"
	// "./brute"
)


func main() {
	if len(os.Args) != 5 {
		fmt.Println("Invalid arg count. Should be:")
		fmt.Println("* config file")
		fmt.Println("* graph file")
		fmt.Println("* days count")
		fmt.Println("* running mode(-USER or -AUTOCONF)")

		os.Exit(1)
	}

	switch mode := os.Args[4]; mode {
	case "-USER":
		daycount, _ := strconv.Atoi(os.Args[3])
		ret_err := interf.RunUserMode(os.Args[1], os.Args[2], daycount)
		os.Exit(ret_err)
	case "-AUTOCONF":
		daycount, _ := strconv.Atoi(os.Args[3])
		ret_err := interf.RunAutoconfMode(os.Args[1], os.Args[2], daycount)
		os.Exit(ret_err)
	default:
		fmt.Println("Unknown configuration.")
		os.Exit(1)
	}
}