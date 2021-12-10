package aco

import (
	"math"
	"math/rand"
	"time"
)


func (ant *Ant) getVision() []float64 {
	var (
		vis = make([]float64, 0)
		denom = 0.0
	)

	for i, l := range ant.tour[ant.position] {
		if l != 0 {
			term := math.Pow((1.0 / float64(l)), ant.colony.conf.TRACE_WEIGHT) * +
				math.Pow(ant.colony.phero_matrix[ant.position][i], ant.colony.conf.TOUR_VISIB)
			denom += term
			vis = append(vis, term)
		} else {
			vis = append(vis, 0)
		}
	}

	for i := 0; i < len(vis); i++ {
		vis[i] /= denom
	}

	return vis
}

func fortuneWheel(vis []float64) int {
	var (
		cumul_sum = make([]float64, len(vis))
		coin = rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
		chosen = -1
	)

	for i := 0; i < len(vis); i++ {
		for j := i; j < len(vis); j++ {
			cumul_sum[i] += vis[j]
		}
	}
	for i := 0; i < len(cumul_sum); i++ {
		if i == len(cumul_sum) - 1 {
			if 0.0 <= coin && coin <= cumul_sum[i] {
				chosen = i
			}
		} else {
			if cumul_sum[i + 1] < coin && coin <= cumul_sum[i] {
					chosen = i
			}
		}
	}

	return chosen
}

func (ant *Ant) tourAnt(root int) {
	for i := range ant.tour[ant.position] {
		ant.tour[i][ant.position] = 0
		ant.tour[ant.position][i] = 0
	}
	ant.memory[ant.position][root] = true	
	ant.position = root
}

func (ant *Ant) updatePhTrace() {
	delta := 0.0

	for i := 0; i < len(ant.colony.phero_matrix); i++ {
		for j, phero := range ant.colony.phero_matrix[i] {
			if ant.colony.gr_matrix[i][j] != 0 {
				if ant.memory[i][j] {
					delta = ant.colony.conf.Q / float64(ant.colony.gr_matrix[i][j])
				} else {
					delta = 0
				}
				ant.colony.phero_matrix[i][j] = (1 - ant.colony.conf.EVAP_RATE) * (float64(phero) + delta)
				ant.colony.phero_matrix[j][i] = (1 - ant.colony.conf.EVAP_RATE) * (float64(phero) + delta)
			}
			if ant.colony.phero_matrix[i][j] <= 0 {
				ant.colony.phero_matrix[i][j] = ant.colony.conf.PHERO_EPS
				ant.colony.phero_matrix[j][i] = ant.colony.conf.PHERO_EPS
			}
		}
	}
}

func (ant *Ant) tourLength() int {
	var (
		length = 0
		last   = ant.position
	)

	for i, r := range ant.memory {
		for root, passed := range r {
			if passed {
				length += ant.colony.gr_matrix[i][root]
			}
		}
	}
	length += ant.colony.gr_matrix[ant.start][last]

	return length
}

func (ant *Ant) startTour() {
	next := -1


	for flag := true; flag; flag = (next != -1) {
		vis := ant.getVision()

		next = fortuneWheel(vis)

		if (next != -1) {
			ant.tourAnt(next)
			ant.updatePhTrace()
		}
	} 
}