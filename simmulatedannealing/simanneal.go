package simmulatedannealing

import (
	"math"
	"math/rand"
	"time"
	"tsp-heuristic-benchmark/readdata"
)

type annealObj struct {
	coords       readdata.Nodes
	alpha        float64
	t            float64
	stopIter     int
	initIter     int
	bestSolution readdata.Nodes
	bestFitness  float64
	currSol      readdata.Nodes
	currFit      float64
	n            int
}

func pAccept(data readdata.Nodes, candidateFitness float64, temp int) (res float64) {
	currFit := readdata.Fitness(readdata.InitialSolution(data))
	res = math.Exp(-math.Abs(candidateFitness-currFit) / float64(temp))
	return
}

func accept(candidate readdata.Nodes, s annealObj) annealObj {
	rand.Seed(time.Now().UnixMilli())
	candidateFitness := readdata.Fitness(candidate)

	p := &s

	if candidateFitness < s.currFit {

		s.currSol = readdata.Changelist(candidate)
		p.currFit = candidateFitness

		if candidateFitness < s.bestFitness {
			s.bestSolution = readdata.Changelist(candidate)
			p.bestFitness = candidateFitness
			return s
		} else {
			return s
		}
	} else {
		if rand.Float64() < pAccept(s.coords, candidateFitness, int(s.t)) {
			s.currSol = readdata.Changelist(candidate)
			p.currFit = candidateFitness
			return s
		}
	}

	return s

}

func SimulationAnnealing(filename string, temp int, iter int) float64 {
	nodes := readdata.ReadData(filename)
	n := len(nodes)

	cs := readdata.InitialSolution(nodes)
	cf := readdata.Fitness(readdata.InitialSolution(nodes))

	var s annealObj

	s.coords = nodes
	stopTemp := 0.0
	s.t = float64(temp)

	s.stopIter = iter
	s.initIter = 1

	s.alpha = 0.995

	s.currSol = cs
	s.currFit = cf

	s.n = n

	s.bestSolution = cs
	s.bestFitness = cf

	for s.t > stopTemp && s.stopIter > s.initIter {
		rand.Seed(time.Now().UnixMilli())
		candidate := readdata.Changelist(readdata.InitialSolution(s.coords))
		s = accept(candidate, s)
		s.t = s.alpha * s.t
		s.initIter += 1
	}

	return s.bestFitness
}
