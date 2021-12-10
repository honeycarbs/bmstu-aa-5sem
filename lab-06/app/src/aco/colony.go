package aco

func makeColony(gr_matrix [][]int, iters int, conf_filename string) *Colony {

	colony := new(Colony)
	conf := ParceConfigFile(conf_filename)

	if (conf == nil) {
		return nil
	} 
	colony.conf = *conf

	colony.gr_matrix = gr_matrix
	colony.iters = iters
	colony.phero_matrix = make([][]float64, len(colony.gr_matrix))

	for i := 0; i < len(colony.gr_matrix); i++ {
		colony.phero_matrix[i] = make([]float64, len(colony.gr_matrix[i]))

		for j := range colony.phero_matrix[i] {
			colony.phero_matrix[i][j] = colony.conf.PHERO_INIT
		}
	}
	return colony
}

func (col *Colony) placeAnt(start int) *Ant {
	ant := new(Ant)
	ant.colony   = col
	ant.tour     = make([][]int, len(col.gr_matrix))
	ant.memory   = make([][]bool, len(col.gr_matrix))
	ant.position = start
	ant.start    = start

	for i := range col.gr_matrix {
		ant.tour[i] = make([]int, len(col.gr_matrix[i]))
		copy(ant.tour[i], col.gr_matrix[i])
	}

	for i := range ant.memory {
		ant.memory[i] = make([]bool, len(col.gr_matrix))
	}

	return ant
}

func ACOWrapper(gr_matrix [][]int, iters int, conf_filename string) *[]int {
	shortest := make([]int, len(gr_matrix))

	colony := makeColony(gr_matrix, iters, conf_filename)
	if (colony == nil) {
		return nil
	}

	for i := 0; i < colony.iters; i++ {
		for j := 0; j < len(colony.gr_matrix); j++ {
			ant := colony.placeAnt(j)

			ant.startTour()
			current := ant.tourLength()

			if current < shortest[j] || shortest[j] == 0 {
				shortest[j] = current
			}
		}
	}
	return &shortest
}

