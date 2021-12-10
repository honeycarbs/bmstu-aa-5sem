package brute

import "fmt"


func permutations(path []int) [][]int {
	var res [][]int
	__permutationHelper(path, &res, 0)

	return res
}

func __permutationHelper(path []int, res *[][]int, k int) {
	for i := k; i < len(path); i++ {
		path[i], path[k] = path[k], path[i]
		__permutationHelper(path, res, k + 1)
		path[k], path[i] = path[i], path[k]
	}
	if (k == len(path) - 1) {
		r := make([]int, len(path))
		copy(r, path)
		*res = append(*res, r)
		return
	}
}

func BruteForce(gr_matrix [][]int) []int {
	var (
		roots = make([]int, 0)
		graph_len = len(gr_matrix)
		min_path = -1
	)

	for source := 0; source < graph_len; source++ {
		cur_roots := make([]int, 0)
		for i := 0; i < graph_len; i++ {
			if i != source {
				cur_roots = append(cur_roots, i)
			}
		}
		
		next_permutation := permutations(cur_roots)
		for _, perm := range next_permutation {
			fmt.Printf("Permutations: %v %v %v\n",source, perm, source);
			curr_cost := 0
			k := source

			for _, j := range perm {
				curr_cost += gr_matrix[k][j]
				k = j
			}
			curr_cost += gr_matrix[k][source]
			if curr_cost < min_path || min_path == -1 {
				min_path = curr_cost
			}
		}

		roots = append(roots, min_path)
		min_path = -1
	}
	return roots
}
