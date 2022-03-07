package main

import (
	"fmt"
	"tsp-heuristic-benchmark/simmulatedannealing"
)

func main() {

	res := simmulatedannealing.SimulationAnnealing("data/ch150.tsp", 100, 100)
	fmt.Println("Simulation Annealing Result is:", res)
}
