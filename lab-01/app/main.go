package main

import dist "./distance"
import "fmt"

func main()  {
	w1 := []rune("cat")
	w2 := []rune("cta")

	fmt.Println(dist.LevenshteinPartialMemo(w1, w2))

	fmt.Println(dist.LevenshteinMatrixIterative(w1, w2))

	fmt.Println(dist.LevenshteinRecursive(w1, w2))

	fmt.Println(dist.DamerauLevenshteinIterative(w1, w2))

	fmt.Println(dist.DamerauLevenshteinRecursive(w1, w2))
}