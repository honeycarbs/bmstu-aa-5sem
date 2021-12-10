package aco


type Ant struct {
	colony  *Colony
	tour     [][]int
	memory   [][]bool
	position int
	start    int
}

type Colony struct {
	gr_matrix    [][]int
	phero_matrix [][]float64
	iters        int
	conf         Config
}

type Config struct {
	TRACE_WEIGHT float64 // alpha
	TOUR_VISIB   float64 // beta
	EVAP_RATE    float64 // p
	Q        	 float64 // q
	PHERO_INIT   float64 // start pheromone conc
	PHERO_EPS    float64
}