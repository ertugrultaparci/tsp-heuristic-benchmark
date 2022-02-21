package simmulatedannealing

import (
	"math"
	"math/rand"
	"time"
	"tsp-heuristic-benchmark/readdata"
)

type annealobj struct {
	coords        readdata.Nodes
	alpha         float64
	t             float64
	stoppint_iter int
	initial_iter  int
	best_solution readdata.Nodes
	best_fitness  float64
	curr_sol      readdata.Nodes
	curr_fit      float64
	n             int
}

func p_accept(data readdata.Nodes, candidate_fitness float64, temp int) float64 {
	curr_fit := readdata.Fitness(readdata.Initial_solution(data))
	return math.Exp(-math.Abs(candidate_fitness-curr_fit) / float64(temp))
}

func accept(candidate readdata.Nodes, s annealobj) annealobj {
	rand.Seed(time.Now().UnixMilli())
	candidate_fitness := readdata.Fitness(candidate)

	if candidate_fitness < s.curr_fit {

		s.curr_sol = readdata.Changelist(candidate)
		s = setcurfit(s, candidate_fitness)

		if candidate_fitness < s.best_fitness {
			s.best_solution = readdata.Changelist(candidate)
			s = setbestfit(s, candidate_fitness)
			return s
		} else {
			return s
		}
	} else {
		if rand.Float64() < p_accept(s.coords, candidate_fitness, int(s.t)) {
			s.curr_sol = readdata.Changelist(candidate)
			s = setcurfit(s, candidate_fitness)
			return s
		}
	}

	return s

}

func Simmulationannealing(filename string, temp int, iter int) float64 {
	nodes := readdata.Read_data(filename)
	n := len(nodes)

	cs := readdata.Initial_solution(nodes)
	cf := readdata.Fitness(readdata.Initial_solution(nodes))

	var s annealobj

	s.coords = nodes
	stop_temp := 0.0
	s.t = float64(temp)

	s.stoppint_iter = iter
	s.initial_iter = 1

	s.alpha = 0.995

	s.curr_sol = cs
	s.curr_fit = cf

	s.n = n

	s.best_solution = cs
	s.best_fitness = cf

	for s.t > stop_temp && s.stoppint_iter > s.initial_iter {
		rand.Seed(time.Now().UnixMilli())
		candidate := readdata.Changelist(readdata.Initial_solution(s.coords))
		s = accept(candidate, s)
		s.t = s.alpha * s.t
		s.initial_iter += 1
	}

	return s.best_fitness
}

func setcurfit(s annealobj, cf float64) annealobj {
	var n annealobj

	n.coords = s.coords
	n.alpha = s.alpha
	n.t = s.t
	n.stoppint_iter = s.stoppint_iter
	n.initial_iter = s.initial_iter
	n.best_solution = s.best_solution
	n.best_fitness = s.best_fitness
	n.curr_sol = s.curr_sol
	n.curr_fit = cf
	n.n = s.n

	return n

}

func setbestfit(s annealobj, bf float64) annealobj {
	var n annealobj

	n.coords = s.coords
	n.alpha = s.alpha
	n.t = s.t
	n.stoppint_iter = s.stoppint_iter
	n.initial_iter = s.initial_iter
	n.best_solution = s.best_solution
	n.best_fitness = bf
	n.curr_sol = s.curr_sol
	n.curr_fit = s.curr_fit
	n.n = s.n

	return n

}
