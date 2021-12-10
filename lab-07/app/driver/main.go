package main

import (
	"fmt"
	"os"

	dc "./dictry"
	"github.com/inancgumus/screen"
)

func main() {
	args := os.Args

	if (len(args) == 1) {
		var filename, choise, skey string
		fmt.Print("Input .csv file with dictionary: ")
		fmt.Fscan(os.Stdin, &filename)

		d, err := dc.FromCSV(filename)
		if (err != nil) {
			fmt.Printf("Cannot open file %v: no such file or corrupt.\n", filename)
			os.Exit(1)
		}

		screen.Clear()
		screen.MoveTopLeft()
		
		fmt.Printf("Opened dictionary: %v\n", filename)
		fmt.Println("Options:\n0. Brute force approach\n1. Binary search\n2. Binary search in segmented dictioanry")
		fmt.Fscan(os.Stdin, &choise)

		fmt.Print("Input key to search: ")
		fmt.Fscan(os.Stdin, &skey)

		switch {
		case choise == "0":
			val, compars, err := d.BruteForce(skey)
			if err == nil {
				fmt.Printf("Key %v found : %v (%v compares)\n", skey, val, compars)
			} else {
				fmt.Println(err)
			}
			// fmt.Println(val, compars)
			os.Exit(0)
	
		case choise == "1":
			val, compars, _ := d.BinarySearch(skey)
			if val != 0 {
				fmt.Printf("Key %v found : %v (%v compares)\n", skey, val, compars)
			} else {
				fmt.Println("NOT FOUND")
			}
			os.Exit(0)
		
		case choise == "2":
			val, compars, _ := d.SectionedBinary(skey)
			if val != 0 {
				fmt.Printf("Key %v found : %v (%v compares)\n", skey, val, compars)
			} else {
				fmt.Println("NOT FOUND")
			}
			os.Exit(0)
		default:
			fmt.Println("Unknown option.")
			os.Exit(0)
	}



	} else if (len(args) == 4) {
		args = args[1:]

		d, err := dc.FromCSV(args[0])
		if (err != nil) {
			os.Exit(1)
		}

		switch {
		case args[1] == "0":
			_, compars, _ := d.BruteForce(args[2])
			fmt.Println(compars)
			os.Exit(0)
	
		case args[1] == "1":
			_, compars, _ := d.BinarySearch(args[2])
			fmt.Println(compars)
			os.Exit(0)
		
		case args[1] == "2":
			_, compars, _ := d.SectionedBinary(args[2])
			fmt.Println(compars)
			os.Exit(0)
	}
	}
}